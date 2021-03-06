package main

import (
	// Native packages
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"time"

	// Third-party packages
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	// "google.golang.org/grpc/metadata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/go-yaml/yaml"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
	"github.com/NicholasBunn/mastersCaseStudy/generalComponents/authenticationStuff"
	"github.com/NicholasBunn/mastersCaseStudy/interceptors/go"

	// Proto packages
	serverPB "github.com/NicholasBunn/mastersCaseStudy/protoFiles/go/authenticationService/v1"
)

var (
	// Addresses
	addrMyself string

	// JWT stuff
	secretKey     string
	tokenDuration time.Duration

	accessibleRoles map[string][]string // This is a map of service calls with their required permission levels

	// Database stuff
	dbDriverName string
	dbUsername string
	dbPassword string
	dbPort string
	dbName string

	// Logging stuff
	DebugLogger   *log.Logger
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger

	// Metric interceptors
	serverMetricInterceptor *interceptors.ServerMetricStruct
)

func init() {
	/* The init functin is used to load in configuration variables, and set up the logger and metric interceptors whenever the service is started
	*/

	// ________CONFIGURATION________
	// Load YAML configurations into config struct
	config, _ := DecodeConfig("configuration.yaml")

	// Load port addresses from config
	addrMyself = os.Getenv("AUTHENTICATIONHOST") + ":" + config.Server.Port.Myself

	// Load JWT parameters from config
	secretKey = config.Server.Authentication.Jwt.SecretKey
	tokenDuration = time.Duration(config.Server.Authentication.Jwt.TokenDuration) * (time.Minute)

	accessibleRoles = map[string][]string{
		config.Server.Authentication.AccessLevel.Name.CreateNewUser: config.Server.Authentication.AccessLevel.Role.CreateNewUser,
	}

	// Load database details from config
	dbDriverName = config.Database.DriverName
	dbUsername = config.Database.User.Username
	dbPassword = config.Database.User.Password
	dbPort = os.Getenv("DATABASEHOST") + ":" + config.Database.Details.Port
	dbName = config.Database.Details.DBName

	// If the file doesn't exist, create it, otherwise append to the file
	pathSlice := strings.Split(os.Args[0], "/") // This just extracts the services name (filename)
	file, err := os.OpenFile("program logs/"+pathSlice[len(pathSlice)-1]+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		// If opening the log file throws an error, continue to create the loggers but print to terminal instead
		log.Println("Unable to initialise log file, good luck :)")
	} else {
		log.SetOutput(file)
	}

	DebugLogger = log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

	// Metric interceptor
	serverMetricInterceptor = interceptors.NewServerMetrics("AuthenticationService") // Custom metric (Prometheus) interceptor
}

func main() {
	/* The main function sets up a server to listen on the specified port,
	encrypts the server connection with TLS, and registers the services on
	offer */

	InfoLogger.Println("Stated authentication service")

	// Load in TLS credentials
	creds, err := loadTLSCredentials()
	if err != nil {
		ErrorLogger.Printf("Error loading TLS credentials: ", err)
	} else {
		DebugLogger.Println("Succesfully loaded TLS certificates")
	}

	// Create a listener on the specified tcp port
	listener, err := net.Listen("tcp", addrMyself)
	if err != nil {
		ErrorLogger.Fatalf("Failed to listen on port %v: \n%v", addrMyself, err)
	}
	InfoLogger.Println("Listening on port: ", addrMyself)

	// Create the interceptors required for this connection
	authInterceptor := interceptors.ServerAuthStruct{          // Custom auth (JWT) interceptor
		JwtManager:	authentication.NewJWTManager(secretKey, tokenDuration),
		AuthenticatedMethods: accessibleRoles,
	}

	rateLimitInterceptor := interceptors.ServerRateLimitStruct {
		CallLimit: 5,
	}

	// Create an interceptor chain with the above interceptors
	interceptorChain := grpc_middleware.ChainUnaryServer(
		serverMetricInterceptor.ServerMetricInterceptor,
		authInterceptor.ServerAuthInterceptor,
		rateLimitInterceptor.ServerRateLimitInterceptor,
	)
	
	fmt.Println(creds)
	// Create a gRPC server object
	authenticationServer := grpc.NewServer(
		// grpc.Creds(creds), // Add the TLS credentials to this server
		grpc.UnaryInterceptor(interceptorChain), // Add the interceptor chain to this server
	)

	// Attach the authentication service offering to the server
	serverPB.RegisterAuthenticationServiceServer(authenticationServer, &server{})
	DebugLogger.Println("Succesfully registered Authentication Service to the server")

	// Start the server
	if err := authenticationServer.Serve(listener); err != nil {
		ErrorLogger.Fatalf("Failed to expose service: \n%v", err)
	}
}

// ________REQUIRED STRUCTURES________

type Config struct {
	Server struct {
		Port struct {
			Myself string `yaml:"myself"`
		} `yaml:"port"`
		Authentication struct {
			Jwt struct {
				SecretKey     string `yaml:"secretKey"`
				TokenDuration int    `yaml:"tokenDuration"`
			} `yaml:"jwt"`
			AccessLevel struct {
				Name struct {
					CreateNewUser string `yaml:"createNewUser`
				} `yaml:"name"`
				Role struct {
					CreateNewUser []string `yaml:"createNewUser`
				} `yaml:"role"`
			} `yaml:"accessLevel"`
		} `yaml:"authentication"`
	} `yaml:"server"`
	
	Database struct {
		DriverName string `yaml:"driverName"`
		User struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		}	`yaml:"user"`
		Details struct {
			Port string `yaml:"port"`
			DBName string `yaml:"dbName"`
		}	`yaml:"details"`
	} `yaml:"database"`
}

type server struct {
	// Use this to implement the authentication service

	serverPB.UnimplementedAuthenticationServiceServer
}

// ________IMPLEMENT THE OFFERED SERVICES________

func (s *server) LoginAuth(ctx context.Context, request *serverPB.LoginAuthRequest) (*serverPB.LoginAuthResponse, error) {
	/* This service logs the user in by checking the provided details against a user
	database. If the user exists, a JWT is generated and returned to them. */

	InfoLogger.Println("Received LoginAuth service call")
	
	// Find the user with the provided username, return a NotFound error if they don't exist
	user, err := find(request.Username)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "cannot find user: %v", err)
	}

	// Check that the provided password matches the password stored in the DB
	if !user.CheckPassword(request.Password) {
		ErrorLogger.Println("Failed to verify user: Passwords don't match.")
		return nil, status.Errorf(codes.Unauthenticated, "the password you provided is incorrect")
	}

	// Create a jwtManager object for the user
	jwtManager := authentication.JWTManager{
		SecretKey:     secretKey,
		TokenDuration: tokenDuration,
	}

	// Generate and return a JWT for the user
	token, err := jwtManager.GenerateToken(user)
	if err != nil {
		ErrorLogger.Println("Failed to generate access token: \n", err)
		return nil, status.Errorf(codes.Internal, "could not generate access token")
	}

	InfoLogger.Println("Successfully generated access token.")

	// Create and populate the response message for the request being served
	response := &serverPB.LoginAuthResponse{
		Permissions: user.Role,
		AccessToken: token,
	}

	return response, nil
}

// ________SUPPORTING FUNCTIONS________

func save(user *authentication.User) error {
	// Still need to implement
	return nil
}

func find(username string) (*authentication.User, error) {
	/* This function takes a username as an input and fetches the user details relating to that user from a mySQL database.
	*/

	// Create connection to the user database
	accessDetails := fmt.Sprintf("%s:%s@(%s)/%s", dbUsername, dbPassword, dbPort, dbName)
    db, err := sql.Open(dbDriverName, accessDetails)
    if (err != nil) {
		ErrorLogger.Println("Failed to connect to user database: \n", err)
        return nil, status.Errorf(codes.Unavailable, "Failed to connect to user database.")
    }

	defer db.Close()

    // Create select query to fetch hashed password and user permission (paramaterised to avoid those pesky SQL injections)    
	myQuery := "SELECT * FROM users WHERE username = ?;"

	// Execute query
	userInfo, err := db.Query(myQuery, username)

	if err != nil {
		// Close the database connection
		db.Close()

		ErrorLogger.Println("Failed to query user database: \n", err)
        return nil, status.Errorf(codes.Internal, "Failed to query user database.")
    } else if (!userInfo.Next()) {
		// If userInfo.Next() returns False, then no results were returned by the DB (and thus, the requested username doesn't exist)

		// Close the database connection
		db.Close()

		ErrorLogger.Println("Failed to log user in: Requested user does not exist.")
		return nil, status.Errorf(codes.NotFound, "requested user does not exist.") 
	} else {
		// Create user object
		var userObject authentication.User
		
		// Scan the result into our user object
		err = userInfo.Scan(&userObject.Username, &userObject.HashedPassword, &userObject.Role)
		if err != nil {
			ErrorLogger.Println("Failed to map database response: \n", err)
			return nil, status.Errorf(codes.Internal, "Failed to map database response.") 
		}

		// Close the database connection
		db.Close()

		// Return the user object
		return &userObject, nil
	}
}

func DecodeConfig(configPath string) (*Config, error) {
	
	// Create a new config structure
	config := &Config{}

	// Open the config file
	file, err := os.Open(configPath)
	if err != nil {
		fmt.Println("Could not open config file")
		return nil, err
	}
	defer file.Close()

	// Initialise a new YAML decoder
	decoder := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := decoder.Decode(&config); err != nil {
		fmt.Println("Could not decode config file: \n", err)
		return nil, err
	}

	return config, nil
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	/* This function loads TLS credentials for both the client and server,
	enabling mutual TLS authentication between the client and server. It takes no inputs and returns a gRPC TransportCredentials object. */

	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("../../certification/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	// Load the server CA's certificates
	certificatePool := x509.NewCertPool()
	if !certificatePool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add the server CA's certificate")
	}

	// Load the client's certificate and private key
	clientCertificate, err := tls.LoadX509KeyPair("../../certification/client-cert.pem", "../../certification/client-key.pem")
	if err != nil {
		return nil, err
	}

	// Create and return the credentials object
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCertificate},
		RootCAs:      certificatePool,
	}

	return credentials.NewTLS(config), nil
}