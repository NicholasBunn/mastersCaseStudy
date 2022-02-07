package main

import (
	// Native packages
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	// Third-party packages
	"github.com/go-yaml/yaml"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/NicholasBunn/mastersCaseStudy/interceptors/go"

	// Proto packages
	serverPB "github.com/NicholasBunn/mastersCaseStudy/protoFiles/go/processVibrationService/v1"
)

var (
	// Addresses (To be passed in a config file)
	addrMyself string
	addrServer string

	timeoutDuration     int           // The time, in seconds, that the client should wait when dialing (connecting to) the server before throwing an error
	callTimeoutDuration time.Duration // The time, in seconds, that the client should wait when making a call to the server before throwing an error

	// JWT stuff, load this in from config
	authToken string
	secretKey     string
	tokenDuration time.Duration

	authMethods map[string]bool // This is a map of which service calls require authentication

	// Logging stuff
	DebugLogger   *log.Logger
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	/* The init function is used to load in configuration variables, and set up the logger and metric interceptors whenever the service is started
	 */

	// Load YAML configurations into config struct
	config, _ := DecodeConfig("configuration.yaml")

	addrMyself = os.Getenv("[::]") + ":" + config.Server.Port.Myself
	addrServer = os.Getenv("[::]") + ":" + config.Client.Port.Server

	// Load timeouts from config
	timeoutDuration = config.Client.Timeout.Connection
	callTimeoutDuration = time.Duration(config.Client.Timeout.Call) * time.Second

	// Load JWT parameters from config
	authToken = config.Server.Authentication.Jwt.Token
	secretKey = config.Server.Authentication.Jwt.SecretKey
	tokenDuration = time.Duration(config.Server.Authentication.Jwt.TokenDuration) * (time.Minute)

	authMethods = map[string]bool{
		config.Client.AuthenticatedMethods.Name.ServerCall:   config.Client.AuthenticatedMethods.RequiresAuthentication.ServerCall,

	}

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
}

func main() {

	InfoLogger.Println("Setting up client")
	// Load in TLS credentials
	creds, err := loadTLSCredentials()
	if err != nil {
		ErrorLogger.Printf("Error loading TLS credentials: \n%v", err)
	} else {
		DebugLogger.Println("Succesfully loaded TLS certificates")
	}

	// Create the interceptors required for this connection
	authInterceptor := interceptors.ClientAuthStruct{          // Custom auth (JWT) interceptor
		AccessToken:          authToken, // Pass the user's JWT to the outgoing request
		AuthenticatedMethods: authMethods,
	}

	// Create an interceptor chain with the above interceptors
	interceptorChain := grpc_middleware.ChainUnaryClient(
		authInterceptor.ClientAuthInterceptor,
	)

	InfoLogger.Println("Creating connection to server")
	
	// Create an insecure connection to the ocean weather service server
	connServer, err := createSecureServerConnection(
		addrServer, 			// Set the address of the server
		creds,            	// Add the TLS credentials
		timeoutDuration, 	// Set the duration that the client will wait before timing out
		interceptorChain, 	// Add the interceptor chain to this server
	)
	if err != nil {
		ErrorLogger.Printf("Failure creating connection to server: \n%v", err)
	}

	InfoLogger.Println("Successfully created connection to server.")

	clientServer := serverPB.NewProcessVibrationServiceClient(connServer)
	DebugLogger.Println("Succesfully created client connection to server.")

	// Create an ocean weather prediction request message
	requestMessage := serverPB.ProcessRequest{
		UnixTime: []float64{1626325118, 1608812145, 1626332318},
		VibrationX: []float64{0.011090744, 0.03320345, 0.01342132},
		VibrationY: []float64{0.001046832, 0.00335624, 0.00233526},
		VibrationZ: []float64{0.035672354, 0.04567234, 0.00345332},
	}
	DebugLogger.Println("Succesfully created an Request.")

	InfoLogger.Println("Making service invocation call.")
	serverContext, cancel := context.WithTimeout(context.Background(), callTimeoutDuration)
	defer cancel()

	// Invoke the Ocean Weather Service
	start := time.Now()
	responseMessage, err := clientServer.CalculateRMSBatch(serverContext, &requestMessage)
	end := time.Since(start).Seconds()
	DebugLogger.Println(responseMessage)
	if err != nil {
		ErrorLogger.Printf("Failed to make service call: \n", err)
	} else {
		DebugLogger.Println("Successfully made service call.")
		fmt.Printf("The server took %v seconds to response.", end)
		connServer.Close()
	}
}

type Config struct {
	Server struct {
		Port struct {
			Myself string `yaml:"myself"`
		} `yaml:"port"`
		Authentication struct {
			Jwt struct {
				Token			string `yaml:"token"`
				SecretKey     string `yaml:"secretKey"`
				TokenDuration int    `yaml:"tokenDuration"`
			} `yaml:"jwt"`
			AccessLevel struct {
				Name struct {
					EstimatePowerTrain string `yaml:"estimatePowerTrain"`
				} `yaml:"name"`
				Role struct {
					EstimatePowerTrain []string `yaml:"estimatePowerTrain"`
				} `yaml:"role"`
			} `yaml:"accessLevel"`
		} `yaml:"authentication"`
	} `yaml:"server"`


	Client struct {
		Port struct {
			Server      string `yaml:"server"`
		} `yaml:"port"`
		Timeout struct {
			Connection int `yaml:"connection"`
			Call       int `yaml:"call"`
		} `yaml:"timeout"`
		AuthenticatedMethods struct {
			Name struct {
				ServerCall string `yaml:"serverCall"`
			} `yaml:"name"`
			RequiresAuthentication struct {
				ServerCall bool `yaml:"serverCall"`
			} `yaml:"requiresAuthentication"`
		} `yaml:"authenticatedMethods"`
	} `yaml:"client"`
}

// ________SUPPORTING FUNCTIONS________

func DecodeConfig(configPath string) (*Config, error) {
	
	// Create a new config structure
	config := &Config{}

	// Open the config file
	file, err := os.Open(configPath)
	if err != nil {
		fmt.Println("Could not open config file")
		return nil, fmt.Errorf("Failure in Route Analysis Aggregator: \n%v", err)
	}
	defer file.Close()

	// Initialise a new YAML decoder
	decoder := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := decoder.Decode(&config); err != nil {
		fmt.Println("Could not decode config file: \n", err)
		return nil, fmt.Errorf("Failure in Route Analysis Aggregator: \n%v", err)
	}

	return config, nil
}

func createSecureServerConnection(port string, credentials credentials.TransportCredentials, timeout int, interceptor grpc.UnaryClientInterceptor) (*grpc.ClientConn, error) {
	/* This (unexported) function takes a port address, gRPC TransportCredentials object, timeout,
	and UnaryClientInterceptor object as inputs. It creates a connection to the server
	at the port adress and returns a secure gRPC connection with the specified
	interceptor */

	// Create the context for the request
	ctx, cancel := context.WithTimeout(
		context.Background(),
		(time.Duration(timeoutDuration) * time.Second),
	)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,              // Add the created context to the connection
		port,             // Add the port that the server is listening on
		grpc.WithBlock(), // Make the dial a blocking call so that we can ensure the connection is indeed created
		grpc.WithTransportCredentials(credentials), // Add the TLS credentials
		grpc.WithUnaryInterceptor(interceptor),     // Add the provided interceptors to the connection
	)

	// Handle errors, if any
	if err != nil {
		ErrorLogger.Println("Failed to create connection to the server on port: " + port)
		return nil, err
	}

	InfoLogger.Println("Succesfully created connection to the server on port: " + port)
	return conn, nil
}

func createInsecureServerConnection(port string, timeout int, interceptor grpc.UnaryClientInterceptor) (*grpc.ClientConn, error) {
	/* This (unexported) function takes a port address, timeout, and UnaryClientInterceptor
	object as inputs. It creates a connection to the server	at the port adress
	and returns an insecure gRPC connection with the specified interceptor */

	// Create the context for the request
	ctx, cancel := context.WithTimeout(
		context.Background(),
		(time.Duration(timeoutDuration) * time.Second),
	)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,                                    // Add the created context to the connection
		port,                                   // Add the port that the server is listening on
		grpc.WithBlock(),                       // Make the dial a blocking call so that we can ensure the connection is indeed created
		grpc.WithInsecure(),                    // Specify that the connection is insecure (no credentials/authorisation required)
		grpc.WithUnaryInterceptor(interceptor), // Add the provided interceptors to the connection
	)

	// Hamndle errors, if any
	if err != nil {
		ErrorLogger.Println("Failed to create connection to the server on port: " + port)
		return nil, err
	}

	InfoLogger.Println("Succesfully created connection to the server on port: " + port)
	return conn, nil
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