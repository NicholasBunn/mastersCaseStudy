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
	"github.com/go-yaml/yaml"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/NicholasBunn/mastersCaseStudy/generalComponents/aggregatorStuff"
	"github.com/NicholasBunn/mastersCaseStudy/generalComponents/authenticationStuff"
	"github.com/NicholasBunn/mastersCaseStudy/interceptors/go"

	// Proto packages
	oceanWeatherServicePB "github.com/NicholasBunn/mastersCaseStudy/protoFiles/go/oceanWeatherService/v1"
	powerTrainServicePB "github.com/NicholasBunn/mastersCaseStudy/protoFiles/go/powerTrainService/v1"
	serverPB "github.com/NicholasBunn/mastersCaseStudy/protoFiles/go/powerTrainAggregator/v1"
)

var (
	// Addresses (To be passed in a config file)
	addrMyself string
	addrOWS	string
	addrPTS string

	timeoutDuration     int           // The time, in seconds, that the client should wait when dialing (connecting to) the server before throwing an error
	callTimeoutDuration time.Duration // The time, in seconds, that the client should wait when making a call to the server before throwing an error

	// JWT stuff, load this in from config
	secretKey     string
	tokenDuration time.Duration

	accessibleRoles map[string][]string // This is a map of service calls with their required permission levels

	authMethods map[string]bool // This is a map of which service calls require authentication

	// Logging stuff
	DebugLogger   *log.Logger
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger

	// Metric interceptors
	clientMetricInterceptor *interceptors.ClientMetricStruct
	serverMetricInterceptor *interceptors.ServerMetricStruct
)

func init() {
	/* The init function is used to load in configuration variables, and set up the logger and metric interceptors whenever the service is started
	 */

	// Load YAML configurations into config struct
	config, _ := DecodeConfig("configuration.yaml")

	addrMyself = os.Getenv("POWERTRAINHOST") + ":" + config.Server.Port.Myself
	addrOWS = os.Getenv("OWSHOST") + ":" + config.Client.Port.OceanWeatherService
	addrPTS = os.Getenv("PTSHOST") + ":" + config.Client.Port.PowerTrainService

	// Load timeouts from config
	timeoutDuration = config.Client.Timeout.Connection
	callTimeoutDuration = time.Duration(config.Client.Timeout.Call) * time.Second

	// Load JWT parameters from config
	secretKey = config.Server.Authentication.Jwt.SecretKey
	tokenDuration = time.Duration(config.Server.Authentication.Jwt.TokenDuration) * (time.Minute)

	accessibleRoles = map[string][]string{
		config.Server.Authentication.AccessLevel.Name.EstimatePowerTrain: config.Server.Authentication.AccessLevel.Role.EstimatePowerTrain,
	}

	authMethods = map[string]bool{
		config.Client.AuthenticatedMethods.Name.OceanWeatherPrediction:   config.Client.AuthenticatedMethods.RequiresAuthentication.OceanWeatherPrediction,
		config.Client.AuthenticatedMethods.Name.OceanWeatherHistory:   config.Client.AuthenticatedMethods.RequiresAuthentication.OceanWeatherHistory,
		config.Client.AuthenticatedMethods.Name.PowerEstimate:   config.Client.AuthenticatedMethods.RequiresAuthentication.PowerEstimate,
		config.Client.AuthenticatedMethods.Name.CostEstimate:   config.Client.AuthenticatedMethods.RequiresAuthentication.CostEstimate,
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

	// Metric interceptor
	clientMetricInterceptor = interceptors.NewClientMetrics("PowerTrainAggregator") // Custom metric (Prometheus) interceptor
	serverMetricInterceptor = interceptors.NewServerMetrics("PowerTrainAggregator") // Custom metric (Prometheus) interceptor
}

func main() {
	/* The main function sets up a server to listen on the specified port,
	encrypts the server connection with TLS, and registers the services on
	offer */

	InfoLogger.Println("Started power train aggregator.")

	// Load in TLS credentials
	creds, err := loadTLSCredentials()
	if err != nil {
		ErrorLogger.Printf("Error loading TLS credentials: \n%v", err)
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
		rateLimitInterceptor.ServerRateLimitInterceptor,
		authInterceptor.ServerAuthInterceptor,
	)

	fmt.Println(creds)

	// Create a gRPC server object
	ptServer := grpc.NewServer(
		// grpc.Creds(creds), // Add the TLS credentials to this server
		grpc.UnaryInterceptor(interceptorChain), // Add the interceptor chain to this server
	)

	// Attach the analysis service offering to the server
	serverPB.RegisterPTEstimateServiceServer(ptServer, &server{})
	DebugLogger.Println("Successfully registered Estimate Service to the server.")

	// Start the server
	if err := ptServer.Serve(listener); err != nil {
		ErrorLogger.Fatalf("Failed to expose service: \n%v", err)
	}
}

// ________REQUIRED STRUCTS________

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
			OceanWeatherService      string `yaml:"oceanWeatherService"`
			PowerTrainService string `yaml:"powerTrainService"`
		} `yaml:"port"`
		Timeout struct {
			Connection int `yaml:"connection"`
			Call       int `yaml:"call"`
		} `yaml:"timeout"`
			AuthenticatedMethods struct {
				Name struct {
					OceanWeatherPrediction   string `yaml:"oceanWeatherPrediction"`
					OceanWeatherHistory   string `yaml:"oceanWeatherHistory"`
					PowerEstimate   string `yaml:"powerEstimate"`
					CostEstimate   string `yaml:"costEstimate"`
				} `yaml:"name"`
				RequiresAuthentication struct {
					OceanWeatherPrediction   bool `yaml:"oceanWeatherPrediction"`
					OceanWeatherHistory   bool `yaml:"oceanWeatherHistory"`
					PowerEstimate   bool `yaml:"powerEstimate"`
					CostEstimate   bool `yaml:"costEstimate"`
				} `yaml:"requiresAuthentication"`
			} `yaml:"authenticatedMethods"`
	} `yaml:"client"`
}

type server struct {
	// Use this to implement the power estimation service package

	serverPB.UnimplementedPTEstimateServiceServer
}

// ________IMPLEMENT THE OFFERED SERVICES________

func (s *server) EstimatePowerTrain(ctx context.Context, request *serverPB.PTEstimateRequest) (*serverPB.PTEstimateResponse, error) {
	/* This call provides a detailed estimate of the cost and power consumption of the S.A. Agulhas along a proposed route.
	*/

	InfoLogger.Println("Received Estimate Power Train service call.")

	// Load in TLS credentials
	creds, err := loadTLSCredentials()
	if err != nil {
		ErrorLogger.Printf("Error loading TLS credentials: \n%v", err)
	} else {
		DebugLogger.Println("Succesfully loaded TLS certificates")
	}

	// Extract the user's JWT from the incoming request. Can ignore the ok output as ths has already been checked.
	md, _ := metadata.FromIncomingContext(ctx)

	// Create the interceptors required for this connection
	authInterceptor := interceptors.ClientAuthStruct{          // Custom auth (JWT) interceptor
		AccessToken:          md["authorisation"][0], // Pass the user's JWT to the outgoing request
		AuthenticatedMethods: authMethods,
	}

	// Create the retry options to specify how the client should retry connection interrupts
	retryOptions := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(100 * time.Millisecond)), // Use exponential backoff to progressively wait longer between retries
		grpc_retry.WithMax(5), // Set the maximum number of retries
	}

	// Create an interceptor chain with the above interceptors
	interceptorChain := grpc_middleware.ChainUnaryClient(
		clientMetricInterceptor.ClientMetricInterceptor,
		authInterceptor.ClientAuthInterceptor,
		grpc_retry.UnaryClientInterceptor(retryOptions...),
	)

	// ________Query Ocean Weather Service________

	// Create a ssecure connection to the ocean weather service server
	connOWS, err := createSecureServerConnection(
		addrOWS, 			// Set the address of the server
		creds,            	// Add the TLS credentials
		timeoutDuration, 	// Set the duration that the client will wait before timing out
		interceptorChain, 	// Add the interceptor chain to this server
	)
	if err != nil {
		return nil, fmt.Errorf("Failure in Power Train Aggregator: \n%v", err)
	}

	InfoLogger.Println("Creating Ocean Weather Service client.")
	clientOWS := oceanWeatherServicePB.NewOceanWeatherServiceClient(connOWS)
	DebugLogger.Println("Succesfully created client connection to Ocean Weather Service.")

	// Create an ocean weather prediction request message
	requestMessageOWS := oceanWeatherServicePB.OceanWeatherPredictionRequest{
		UnixTime: request.UnixTime,
		Latitude: request.Latitude,
		Longitude: request.Longitude,
	}
	DebugLogger.Println("Succesfully created an Ocean Weather Service Request.")

	InfoLogger.Println("Making Ocean Weather Prediction call.")
	owsContext, cancel := context.WithTimeout(context.Background(), callTimeoutDuration)
	defer cancel()

	// Invoke the Ocean Weather Service
	responseMessageOWS, err := clientOWS.OceanWeatherPrediction(owsContext, &requestMessageOWS)
	if err != nil {
		ErrorLogger.Println("Failed to make Ocean Weather Prediction service call: \n", err)
		return nil, fmt.Errorf("Failure in Power Train Aggregator: \n%v", err)
	} else {
		DebugLogger.Println("Successfully made service call to Ocean Weather Service.")
		connOWS.Close()
	}

	// ________Query Power Train Service________
	
	// Create a secure connection to the power train service server
	connPTS, err := createSecureServerConnection(
		addrPTS,			// Set the address of the server
		creds,            	// Add the TLS credentials
		timeoutDuration,	// Set the duration that the client will wait before timing out
		interceptorChain, 	// Add the interceptor chain to this server
	)
	if err != nil {
		return nil, fmt.Errorf("Failure in Power Train Aggregator: \n%v", err)
	}

	InfoLogger.Println("Creating Power Train Service client.")
	clientPTS := powerTrainServicePB.NewPowerTrainServiceClient(connPTS)
	DebugLogger.Println("Succesfully created client connection to Power Train Service.")

	// Create a power train estimate request message
	requestMessagePTS := powerTrainServicePB.PowerTrainEstimateRequest{
		UnixTime: request.UnixTime,
		PortPropMotorSpeed: request.MotorSpeed,
		StbdPropMotorSpeed: request.MotorSpeed,
		PropellerPitchPort: request.PropPitch,
		PropellerPitchStbd: request.PropPitch,
		Sog: request.SOG,
		WindSpeed: responseMessageOWS.WindSpeed,
		BeaufortNumber: responseMessageOWS.BeaufortNumber,
		WaveDirection: responseMessageOWS.SwellDirection,
		WaveLength: responseMessageOWS.WaveLength,
		ModelType: powerTrainServicePB.ModelTypeEnum_OPENWATER,
	}

	requestMessagePTS.WindDirectionRelative, err = aggregator.CalculateRelativeWindDirection(responseMessageOWS.WindDirection, request.Heading)
	if err != nil {
		ErrorLogger.Println("Failed to calculate relative wind direction: \n", err)
		return nil, fmt.Errorf("Failure in Power Train Aggregator: \n%v", err)
	}

	DebugLogger.Println("Succesfully created a Power Train Estimate Request.")

	InfoLogger.Println("Making Cost Estimate Call.")
	ptsContext, cancel := context.WithTimeout(context.Background(), callTimeoutDuration)
	defer cancel()

	// Invoke the Power Train Service
	responseMessagePTS, err := clientPTS.CostEstimate(ptsContext, &requestMessagePTS)
	if err != nil {
		ErrorLogger.Println("Failed to make Cost Estimate service call: \n", err)
		return nil, fmt.Errorf("Failure in Power Train Aggregator: \n%v", err)
	} else {
		DebugLogger.Println("Successfully made service call to Power Train Service.")
		connPTS.Close()
	}

	// Create the response message
	responseMessage := serverPB.PTEstimateResponse{
		UnixTime: responseMessagePTS.UnixTime,
		PowerEstimate: responseMessagePTS.PowerEstimate,
		CostEstimate: responseMessagePTS.CostEstimate,
		TotalCost: responseMessagePTS.TotalCost,
	}

	return &responseMessage, nil
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