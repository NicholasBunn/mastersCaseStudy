package main

import (
	// Native packages
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	// Third-party packages
	"google.golang.org/grpc"
	"github.com/go-yaml/yaml"

	// Proto packages
	serverPB "github.com/NicholasBunn/mastersCaseStudy/protoFiles/go/webGateway/v1"
	authenticationServicePB "github.com/NicholasBunn/mastersCaseStudy/protoFiles/go/authenticationService/v1"
	routeAnalysisAggregatorPB "github.com/NicholasBunn/mastersCaseStudy/protoFiles/go/routeAnalysisAggregator/v1"
	powerTrainAggregatorPB "github.com/NicholasBunn/mastersCaseStudy/protoFiles/go/powerTrainAggregator/v1"
	vesselMotionAggregatorPB "github.com/NicholasBunn/mastersCaseStudy/protoFiles/go/vesselMotionAggregator/v1"
)

var (
	// Addresses
	addrMyself string
	addrAS string
	addrRAA string
	addrPTA string
	addrVMA string

	timeoutDuration     int           // The time, in seconds, that the client should wait when dialing (connecting to) the server before throwing an error
	callTimeoutDuration time.Duration // The time, in seconds, that the client should wait when making a call to the server before throwing an error

	// JWT stuff
	secretKey     string
	tokenDuration time.Duration

	// Logging stuff
	DebugLogger   *log.Logger
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	/* The init functin is used to load in configuration variables, and set up the logger and metric interceptors whenever the service is started
	*/

	// ________CONFIGURATION________
	// Load YAML configurations into config struct
	config, _ := DecodeConfig("configuration.yaml")

	// Load port addresses from config
	addrMyself = os.Getenv("WEBHOST") + ":" + config.Server.Port.Myself
	addrAS = os.Getenv("AUTHENTICATIONHOST") + ":" + config.Client.Port.AuthenticationService
	addrRAA = os.Getenv("ROUTEANALYSISHOST") + ":" + config.Client.Port.RouteAnalysisAggregator
	addrPTA = os.Getenv("POWERTRAINHOST") + ":" + config.Client.Port.PowerTrainAggregator
	addrVMA = os.Getenv("VESSELMOTIONHOST") + ":" + config.Client.Port.VesselMotionAggregator

	// Load timeouts from config
	timeoutDuration = config.Client.Timeout.Connection
	callTimeoutDuration = time.Duration(config.Client.Timeout.Call) * time.Second
	
	// Load JWT parameters from config
	secretKey = config.Server.Authentication.Jwt.SecretKey
	tokenDuration = time.Duration(config.Server.Authentication.Jwt.TokenDuration) * (time.Minute)

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
	/* The main function sets up a server to listen on the specified port,
	encrypts the server connection with TLS, and registers the services on
	offer */

	InfoLogger.Println("Stated web gateway")

	// Load in TLS credentials
	// creds, err := loadTLSCredentials()
	// if err != nil {
	// 	ErrorLogger.Printf("Error loading TLS credentials")
	// } else {
	// 	DebugLogger.Println("Succesfully loaded TLS certificates")
	// }

	// Create a listener on the specified tcp port
	listener, err := net.Listen("tcp", addrMyself)
	if err != nil {
		ErrorLogger.Fatalf("Failed to listen on port %v: \n%v", addrMyself, err)
	}
	InfoLogger.Println("Listening on port: ", addrMyself)

	// Create a gRPC server object
	webServer := grpc.NewServer(
	// grpc.Creds(creds), // Add the TLS credentials to this server
	// grpc.UnaryInterceptor(interceptors.AuthenticationInterceptor), // Add the interceptor chain to this server
	)

	// Attach the authentication service offering to the server
	serverPB.RegisterLoginServiceServer(webServer, &loginServer{})
	DebugLogger.Println("Succesfully registered Login Service to the server")

	serverPB.RegisterRouteAnalysisAggregatorServer(webServer, &routeAnalysisServer{})
	DebugLogger.Println("Succesfully registered Route Analysis Service to the server")

	serverPB.RegisterRoutePowerAggregatorServer(webServer, &routePowerServer{})
	DebugLogger.Println("Succesfully registered Route Power Service to the server")

	serverPB.RegisterRouteMotionAggregatorServer(webServer, &routeMotionServer{})
	DebugLogger.Println("Succesfully registered Route Motion Service to the server")

	// Start the server
	if err := webServer.Serve(listener); err != nil {
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
		} `yaml:"authentication"`
	} `yaml:"server"`
	
	Client struct {
		Port struct {
			AuthenticationService      string `yaml:"authenticationService"`
			RouteAnalysisAggregator string `yaml:"routeAnalysisAggregator"`
			PowerTrainAggregator string `yaml:"powerTrainAggregator"`
			VesselMotionAggregator string `yaml:"vesselMotionAggregator"`
		} `yaml:"port"`
		Timeout struct {
			Connection int `yaml:"connection"`
			Call       int `yaml:"call"`
		} `yaml:"timeout"`
	// 		AuthenticatedMethods struct {
	// 			Name struct {
	// 				FetchDataService   string `yaml:"fetchDataService"`
	// 				PrepareDataService string `yaml:"prepareDataService"`
	// 				EstimateService    string `yaml:"estimateService"`
	// 			} `yaml:"name"`
	// 			RequiresAuthentication struct {
	// 				FetchDataService   bool `yaml:"fetchDataService"`
	// 				PrepareDataService bool `yaml:"prepareDataService"`
	// 				EstimateService    bool `yaml:"estimateService"`
	// 			} `yaml:"requiresAuthentication"`
	// 		} `yaml:"authenticatedMethods"`
	} `yaml:"client"`
}

type loginServer struct {
	// Use this to implement the authentication service

	serverPB.UnimplementedLoginServiceServer
}

type routeAnalysisServer struct {
	// Use this to implement the route analysis aggregator

	serverPB.UnimplementedRouteAnalysisAggregatorServer
}

type routePowerServer struct {
	// Use this to implement the power train aggregator

	serverPB.UnimplementedRoutePowerAggregatorServer
}

type routeMotionServer struct {
	// Use this to implement the motion aggregator

	serverPB.UnimplementedRouteMotionAggregatorServer
}

// ________IMPLEMENT THE OFFERED SERVICES________

func (s *loginServer) Login(ctx context.Context, request *serverPB.LoginRequest) (*serverPB.LoginResponse, error) {

	InfoLogger.Println("Received Login service call.")
		
	// Create an insecure connection to the login service server
	connAS, err := createInsecureServerConnection(
		addrAS,
		timeoutDuration,
	)
	if err != nil {
		return nil, err
	}

	InfoLogger.Println("Creating Authentication Service client.")
	clientAS := authenticationServicePB.NewAuthenticationServiceClient(connAS)
	DebugLogger.Println("Succesfully created client connection to Authentication Service.")

	// Create a login auth request message
	requestMessageAS := authenticationServicePB.LoginAuthRequest{
		Username: request.Username,
		Password: request.Password,
	}

	DebugLogger.Println("Succesfully created a Login Auth Request.")

	InfoLogger.Println("Making Login Auth Call.")
	asContext, cancel := context.WithTimeout(context.Background(), callTimeoutDuration)
	defer cancel()

	// Invoke the Authentication Service
	responseMessageAS, err := clientAS.LoginAuth(asContext, &requestMessageAS)
	if err != nil {
		ErrorLogger.Println("Failed to make Login Auth service call: \n", err)
		return nil, err
	} else {
		DebugLogger.Println("Successfully made service call to Authentication Service.")
		connAS.Close()
	}

	response := serverPB.LoginResponse{
		Permissions: responseMessageAS.Permissions,
		AccessToken: responseMessageAS.AccessToken,
	}

	return &response, nil
}

func (s *routeAnalysisServer) RouteAnalysis(ctx context.Context, request *serverPB.RouteAnalysisRequest) (*serverPB.RouteAnalysisResponse, error) {
	
	InfoLogger.Println("Received Route Analysis Service Call.")

	// Create an insecure connection to the route analysis aggregator server
	connRAA, err := createInsecureServerConnection(
		addrRAA,
		timeoutDuration,
	)
	if err != nil {
		return nil, err
	}	


	InfoLogger.Println("Creating Route Analysis Aggregator client.")
	clientRAA := routeAnalysisAggregatorPB.NewAnalysisServiceClient(connRAA)
	DebugLogger.Println("Succesfully created client connection to Route Analysis Aggregator.")

	// Create an analysis request message
	requestMessageRAA := routeAnalysisAggregatorPB.AnalysisRequest{
		UnixTime: request.UnixTime,
		Latitude: request.Latitude,
		Longitude: request.Longitude,
		Heading: request.Heading,
		PropPitch: request.PropPitch,
		MotorSpeed: request.MotorSpeed,
		SOG: request.SOG,
	}
	
	DebugLogger.Println("Succesfully created an Analysis Request.")

	InfoLogger.Println("Making Analyse Route Call.")
	raaContext, cancel := context.WithTimeout(context.Background(), callTimeoutDuration)
	defer cancel()

	// Invoke the Route Analysis Aggregator
	responseMessageRAA, err := clientRAA.AnalyseRoute(raaContext, &requestMessageRAA)
	if err != nil {
		ErrorLogger.Println("Failed to make Analyse Route service call: \n", err)

		return nil, err
	} else {
		DebugLogger.Println("Successfully made service call to Route Analysis Aggregator.")
		connRAA.Close()
	}

	fmt.Println(responseMessageRAA)

	response := serverPB.RouteAnalysisResponse{
		UnixTime: responseMessageRAA.UnixTime,
		AveragePower: responseMessageRAA.AveragePower,
		TotalCost: responseMessageRAA.TotalCost,
		AverageRmsX: responseMessageRAA.AverageRmsX,
		AverageRmsY: responseMessageRAA.AverageRmsY,
		AverageRmsZ: responseMessageRAA.AverageRmsZ,
		// ComfortLevel: responseMessageRAA.ComfortLevel,
	}

	return &response, nil
}

func (s *routePowerServer) RoutePower(ctx context.Context, request *serverPB.RoutePowerRequest) (*serverPB.RoutePowerResponse, error) {

	InfoLogger.Println("Received Route Power Service Call.")

	// Create an insecure connection to the route power aggregator server
	connPTA, err := createInsecureServerConnection(
		addrPTA,
		timeoutDuration,
	)
	if err != nil {
		return nil, err
	}	


	InfoLogger.Println("Creating Route Power Aggregator client.")
	clientRPA := powerTrainAggregatorPB.NewPTEstimateServiceClient(connPTA)
	DebugLogger.Println("Succesfully created client connection to Power Train Aggregator.")

	// Create a PT estimate request message
	requestMessagePTA := powerTrainAggregatorPB.PTEstimateRequest{
		UnixTime: request.UnixTime,
		Latitude: request.Latitude,
		Longitude: request.Longitude,
		Heading: request.Heading,
		PropPitch: request.PropPitch,
		MotorSpeed: request.MotorSpeed,
		SOG: request.SOG,
	}
	
	DebugLogger.Println("Succesfully created a PT Estimate Request.")

	InfoLogger.Println("Making Estimate Power Train Call.")
	ptaContext, cancel := context.WithTimeout(context.Background(), callTimeoutDuration)
	defer cancel()

	// Invoke the Power Train Aggregator
	responseMessagePTA, err := clientRPA.EstimatePowerTrain(ptaContext, &requestMessagePTA)
	if err != nil {
		ErrorLogger.Println("Failed to make Estimate Power Train service call: \n", err)

		return nil, err
	} else {
		DebugLogger.Println("Successfully made service call to Power Train Aggregator.")
		connPTA.Close()
	}

	fmt.Println(responseMessagePTA)

	response := serverPB.RoutePowerResponse{
		UnixTime: responseMessagePTA.UnixTime,
		PowerEstimate: responseMessagePTA.PowerEstimate,
		CostEstimate: responseMessagePTA.CostEstimate,
		TotalCost: responseMessagePTA.TotalCost,
	}

	return &response, nil
}

func (s *routeMotionServer) RouteMotion(ctx context.Context, request *serverPB.RouteMotionRequest) (*serverPB.RouteMotionResponse, error) {

	InfoLogger.Println("Received Route Motion Service Call.")

	// Create an insecure connection to the route motion aggregator server
	connVMA, err := createInsecureServerConnection(
		addrVMA,
		timeoutDuration,
	)
	if err != nil {
		return nil, err
	}	


	InfoLogger.Println("Creating Route Motion Aggregator client.")
	clientVMA := vesselMotionAggregatorPB.NewVMEstimateServiceClient(connVMA)
	DebugLogger.Println("Succesfully created client connection to Vessel Motion Aggregator.")

	// Create a VM estimate request message
	requestMessageVMA := vesselMotionAggregatorPB.VMEstimateRequest{
		UnixTime: request.UnixTime,
		Latitude: request.Latitude,
		Longitude: request.Longitude,
		Heading: request.Heading,
		PropPitch: request.PropPitch,
		MotorSpeed: request.MotorSpeed,
		SOG: request.SOG,
	}
	
	DebugLogger.Println("Succesfully created a VM Estimate Request.")

	InfoLogger.Println("Making Estimate Vessel Motion Call.")
	vmaContext, cancel := context.WithTimeout(context.Background(), callTimeoutDuration)
	defer cancel()

	// Invoke the Vessel Motion Aggregator
	responseMessageVMA, err := clientVMA.EstimateVesselMotion(vmaContext, &requestMessageVMA)
	if err != nil {
		ErrorLogger.Println("Failed to make Estimate Vessel Motion service call: \n", err)

		return nil, err
	} else {
		DebugLogger.Println("Successfully made service call to Vessel Motion Aggregator.")
		connVMA.Close()
	}

	fmt.Println(responseMessageVMA)

	response := serverPB.RouteMotionResponse{
		UnixTime: responseMessageVMA.UnixTime,
		AccelerationEstimateX: responseMessageVMA.AccelerationEstimateX,
		AccelerationEstimateY: responseMessageVMA.AccelerationEstimateY,
		AccelerationEstimateZ: responseMessageVMA.AccelerationEstimateZ,
	}

	return &response, nil
}

// ________SUPPORTING FUNCTIONS________

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

func createInsecureServerConnection(port string, timeout int) (*grpc.ClientConn, error) {
	/* This (unexported) function takes a port address and timeout as inputs. It creates a connection to the server	at the port adress
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
	)

	// Hamndle errors, if any
	if err != nil {
		ErrorLogger.Println("Failed to create connection to the server on port: " + port)
		return nil, err
	}

	InfoLogger.Println("Succesfully created connection to the server on port: " + port)
	return conn, nil
}
