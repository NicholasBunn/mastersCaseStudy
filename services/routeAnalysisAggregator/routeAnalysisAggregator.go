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
	vesselMotionServicePB "github.com/NicholasBunn/mastersCaseStudy/protoFiles/go/vesselMotionService/v1"
	processVibrationServicePB "github.com/NicholasBunn/mastersCaseStudy/protoFiles/go/processVibrationService/v1"
	comfortServicePB "github.com/NicholasBunn/mastersCaseStudy/protoFiles/go/comfortService/v1"
	serverPB "github.com/NicholasBunn/mastersCaseStudy/protoFiles/go/routeAnalysisAggregator/v1"
)

var (
	// Addresses (To be passed in a config file)
	addrMyself string
	addrOWS	string
	addrPTS string
	addrVMS string
	addrPVS string
	addrCS string

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

	addrMyself = os.Getenv("ROUTEANALYSISHOST") + ":" + config.Server.Port.Myself
	addrOWS = os.Getenv("OWSHOST") + ":" + config.Client.Port.OceanWeatherService
	addrPTS = os.Getenv("PTSHOST") + ":" + config.Client.Port.PowerTrainService
	addrVMS = os.Getenv("VMSHOST") + ":" + config.Client.Port.VesselMotionService
	addrCS = os.Getenv("CSHOST") + ":" + config.Client.Port.ComfortService
	addrPVS = os.Getenv("PVSHOST") + ":" + config.Client.Port.ProcessVibrationService

	// Load timeouts from config
	timeoutDuration = config.Client.Timeout.Connection
	callTimeoutDuration = time.Duration(config.Client.Timeout.Call) * time.Second

	// Load JWT parameters from config
	secretKey = config.Server.Authentication.Jwt.SecretKey
	tokenDuration = time.Duration(config.Server.Authentication.Jwt.TokenDuration) * (time.Minute)

	accessibleRoles = map[string][]string{
		config.Server.Authentication.AccessLevel.Name.AnalyseRoute: config.Server.Authentication.AccessLevel.Role.AnalyseRoute,
	}

	authMethods = map[string]bool{
		config.Client.AuthenticatedMethods.Name.OceanWeatherPrediction:   config.Client.AuthenticatedMethods.RequiresAuthentication.OceanWeatherPrediction,
		config.Client.AuthenticatedMethods.Name.OceanWeatherHistory:   config.Client.AuthenticatedMethods.RequiresAuthentication.OceanWeatherHistory,
		config.Client.AuthenticatedMethods.Name.PowerEstimate:   config.Client.AuthenticatedMethods.RequiresAuthentication.PowerEstimate,
		config.Client.AuthenticatedMethods.Name.CostEstimate:   config.Client.AuthenticatedMethods.RequiresAuthentication.CostEstimate,
		config.Client.AuthenticatedMethods.Name.MotionEstimate:   config.Client.AuthenticatedMethods.RequiresAuthentication.MotionEstimate,
		config.Client.AuthenticatedMethods.Name.CalculateRMSBatch:   config.Client.AuthenticatedMethods.RequiresAuthentication.CalculateRMSBatch,
		config.Client.AuthenticatedMethods.Name.CalculateRMSSeries:   config.Client.AuthenticatedMethods.RequiresAuthentication.CalculateRMSSeries,
		config.Client.AuthenticatedMethods.Name.ComfortRating:   config.Client.AuthenticatedMethods.RequiresAuthentication.ComfortRating,
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
	clientMetricInterceptor = interceptors.NewClientMetrics("RouteAnalysisAggregator") // Custom metric (Prometheus) interceptor
	serverMetricInterceptor = interceptors.NewServerMetrics("RouteAnalysisAggregator") // Custom metric (Prometheus) interceptor
}

func main() {
	/* The main function sets up a server to listen on the specified port,
	encrypts the server connection with TLS, and registers the services on
	offer */

	InfoLogger.Println("Started route analysis aggregator.")

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
	// Create an interceptor chain with the above interceptors
	interceptorChain := grpc_middleware.ChainUnaryServer(
		serverMetricInterceptor.ServerMetricInterceptor,
		authInterceptor.ServerAuthInterceptor,
	)

	// Create a gRPC server object
	analysisServer := grpc.NewServer(
		grpc.Creds(creds), // Add the TLS credentials to this server
		grpc.UnaryInterceptor(interceptorChain), // Add the interceptor chain to this server
	)

	// Attach the analysis service offering to the server
	serverPB.RegisterAnalysisServiceServer(analysisServer, &server{})
	DebugLogger.Println("Successfully registered Analysis Service to the server.")

	// Start the server
	if err := analysisServer.Serve(listener); err != nil {
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
					AnalyseRoute string `yaml:"analyseRoute"`
				} `yaml:"name"`
				Role struct {
					AnalyseRoute []string `yaml:"analyseRoute"`
				} `yaml:"role"`
			} `yaml:"accessLevel"`
		} `yaml:"authentication"`
	} `yaml:"server"`

	Client struct {
		Port struct {
			OceanWeatherService      string `yaml:"oceanWeatherService"`
			PowerTrainService string `yaml:"powerTrainService"`
			VesselMotionService string `yaml:"vesselMotionService"`
			ProcessVibrationService string `yaml:"processVibrationService"`
			ComfortService string `yaml:"comfortService"`
		} `yaml:"port"`
		Timeout struct {
			Connection int `yaml:"connection"`
			Call       int `yaml:"call"`
		} `yaml:"timeout"`
		AuthenticatedMethods struct {
			Name struct {
				OceanWeatherPrediction string `yaml:"oceanWeatherPrediction"`
				OceanWeatherHistory string `yaml:"oceanWeatherHistory"`
				PowerEstimate string `yaml:"powerEstimate"`
				CostEstimate  string `yaml:"costEstimate"`
				MotionEstimate string `yaml:"motionEsimate"`
				CalculateRMSBatch string `yaml:"calculateRMSBatch"`
				CalculateRMSSeries string `yaml:"calculateRMSSeries"`
				ComfortRating string `yaml:"comfortRating"`
			} `yaml:"name"`
			RequiresAuthentication struct {
				OceanWeatherPrediction bool `yaml:"oceanWeatherPrediction"`
				OceanWeatherHistory bool `yaml:"oceanWeatherHistory"`
				PowerEstimate bool `yaml:"powerEstimate"`
				CostEstimate  bool `yaml:"costEstimate"`
				MotionEstimate bool `yaml:"motionEsimate"`
				CalculateRMSBatch bool `yaml:"calculateRMSBatch"`
				CalculateRMSSeries bool `yaml:"calculateRMSSeries"`
				ComfortRating bool `yaml:"comfortRating"`
			} `yaml:"requiresAuthentication"`
		} `yaml:"authenticatedMethods"`
	} `yaml:"client"`
}

type server struct {
	// Use this to implement the power estimation service package

	serverPB.UnimplementedAnalysisServiceServer
}

// ________IMPLEMENT THE OFFERED SERVICES________

func (s *server) AnalyseRoute(ctx context.Context, request *serverPB.AnalysisRequest) (*serverPB.AnalysisResponse, error) {
	/* This call provides foresight for tactical decision-making by providing a summary of the provided route.
	*/

	InfoLogger.Println("Received Analyse Route service call.")

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
	
	// Create an insecure connection to the ocean weather service server
	connOWS, err := createInsecureServerConnection(
		addrOWS, // Set the address of the server
		timeoutDuration, // Set the duration that the client will wait before timing out
		interceptorChain, // Add the interceptor chain to this server
	)
	if err != nil {
		return nil, fmt.Errorf("Failure in Route Analysis Aggregator: \n%v", err)
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
		return nil, fmt.Errorf("Failure in Route Analysis Aggregator: \n%v", err)
	} else {
		DebugLogger.Println("Successfully made service call to Ocean Weather Service.")
		connOWS.Close()
	}

	// ________Query Power Train Service________
	
	// Create an insecure connection to the power train service server
	connPTS, err := createInsecureServerConnection(
		addrPTS,
		timeoutDuration,
		interceptorChain, // Add the interceptor chain to this server
	)
	if err != nil {
		return nil, fmt.Errorf("Failure in Route Analysis Aggregator: \n%v", err)
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

	DebugLogger.Println(requestMessagePTS)

	requestMessagePTS.WindDirectionRelative, err = aggregator.CalculateRelativeWindDirection(responseMessageOWS.WindDirection, request.Heading)
	if err != nil {
		ErrorLogger.Println("Failed to calculate relative wind direction: \n", err)
		return nil, fmt.Errorf("Failure in Route Analysis Aggregator: \n%v", err)
	}

	DebugLogger.Println("Succesfully created a Power Train Estimate Request.")

	InfoLogger.Println("Making Cost Estimate Call.")
	ptsContext, cancel := context.WithTimeout(context.Background(), callTimeoutDuration)
	defer cancel()

	// Invoke the Power Train Service
	responseMessagePTS, err := clientPTS.CostEstimate(ptsContext, &requestMessagePTS)
	if err != nil {
		ErrorLogger.Println("Failed to make Cost Estimate service call: \n", err)
		return nil, fmt.Errorf("Failure in Route Analysis Aggregator: \n%v", err)
	} else {
		DebugLogger.Println("Successfully made service call to Power Train Service.")
		connPTS.Close()
	}

	// ________Query Vessel Motion Service________
	
	// Create an insecure connection to the power train service server
	connVMS, err := createInsecureServerConnection(
		addrVMS,
		timeoutDuration,
		interceptorChain, // Add the interceptor chain to this server
	)
	if err != nil {
		return nil, fmt.Errorf("Failure in Route Analysis Aggregator: \n%v", err)
	}

	InfoLogger.Println("Creating Vessel Motion Service client.")
	clientVMS := vesselMotionServicePB.NewVesselMotionServiceClient(connVMS)
	DebugLogger.Println("Succesfully created client connection to Vessel Motion Service.")

	// Create a motion estimate request message
	requestMessageVMS := vesselMotionServicePB.MotionEstimateRequest{
		PortPropMotorPower: responseMessagePTS.PowerEstimate,
		Latitude: request.Latitude,
		Heading: request.Heading,
		WaveHeight: responseMessageOWS.SwellHeight,
		WindDirectionRelative: requestMessagePTS.WindDirectionRelative,
		ModelType: vesselMotionServicePB.ModelTypeEnum_OPENWATER,
		QueryLocation: vesselMotionServicePB.LocationOnShipEnum_UNKNOWN_LOCATION,
	}

	requestMessageVMS.WindSpeedRelative, err = aggregator.CalculateRelativeWindSpeed(responseMessageOWS.WindSpeed, requestMessagePTS.WindDirectionRelative, request.SOG)
	if err != nil {
		ErrorLogger.Println("Failed to calculate relative wind direction: \n", err)
		return nil, fmt.Errorf("Failure in Route Analysis Aggregator: \n%v", err)
	}

	DebugLogger.Println("Succesfully created a Motion Estimate Request.")

	InfoLogger.Println("Making Motion Estimate Call.")
	vmsContext,cancel := context.WithTimeout(context.Background(), callTimeoutDuration)
	defer cancel()

	// Invoke the Vessel Motion Service
	responseMessageVMS, err := clientVMS.MotionEstimate(vmsContext, &requestMessageVMS)
	if err != nil {
		ErrorLogger.Println("Failed to make Motion Estimate service call: \n", err)
		return nil, fmt.Errorf("Failure in Route Analysis Aggregator: \n%v", err)
	} else {
		DebugLogger.Println("Successfully made service call to Vessel Motion Service.")
		connVMS.Close()
	}

	// ________Query Process Vibration Service________

	// Create an insecure connection to the process vibration service server
	connPVS, err := createInsecureServerConnection(
		addrPVS,
		timeoutDuration,
		interceptorChain, // Add the interceptor chain to this server
	)
	if err != nil {
		return nil, fmt.Errorf("Failure in Route Analysis Aggregator: \n%v", err)
	}

	InfoLogger.Println("Creating Process Vibration Service client.")
	clientPVS := processVibrationServicePB.NewProcessVibrationServiceClient(connPVS)
	DebugLogger.Println("Successfully created client connection to Process Vibration Service.")

	// Create a process request message
	requestMessagePVS := processVibrationServicePB.ProcessRequest {
		UnixTime: request.UnixTime,
		VibrationX: responseMessageVMS.AccelerationEstimateX,
		VibrationY: responseMessageVMS.AccelerationEstimateY,
		VibrationZ: responseMessageVMS.AccelerationEstimateZ,
	}

	DebugLogger.Println("Succesfully created a Process Request.")

	InfoLogger.Println("Making Calculate RMS Batch Call.")
	pvsContext, cancel := context.WithTimeout(context.Background(), callTimeoutDuration)
	defer cancel()

	// Invoke the Process Vibratiion Service
	responseMessagePVS, err := clientPVS.CalculateRMSBatch(pvsContext, &requestMessagePVS)
	if (err != nil) {
		ErrorLogger.Println("Failed to make Calculate RMS Batch service call: \n", err)
		return nil, fmt.Errorf("Failure in Route Analysis Aggregator: \n%v", err)
	} else {
		DebugLogger.Println("Successfully made service call to Process Vibration Service.")
		connPVS.Close()
	} 

	// ________Query Comfort Service________
	
	// Create an insecure connection to the comfort service server
	connCS, err := createInsecureServerConnection(
		addrCS,
		timeoutDuration,
		interceptorChain, // Add the interceptor chain to this server
	)
	if err != nil {
		return nil, fmt.Errorf("Failure in Route Analysis Aggregator: \n%v", err)
	}

	InfoLogger.Println("Creating Comfort Service client.")
	clientCS := comfortServicePB.NewComfortServiceClient(connCS)
	DebugLogger.Println("Succesfully created client connection to Comfort Service.")

	// Create a comfort request message
	requestMessageCS := comfortServicePB.ComfortRequest{
		UnixTime: request.UnixTime,
		HumanWeightedVibrationX: responseMessageVMS.AccelerationEstimateX,
		HumanWeightedVibrationY: responseMessageVMS.AccelerationEstimateY,
		HumanWeightedVibrationZ: responseMessageVMS.AccelerationEstimateZ,
	}

	DebugLogger.Println("Succesfully created a Comfort Request.")

	InfoLogger.Println("Making Comfort Rating Call.")
	csContext, cancel := context.WithTimeout(context.Background(), callTimeoutDuration)
	defer cancel()

	// Invoke the Comfort Service
	responseMessageCS, err := clientCS.ComfortRating(csContext, &requestMessageCS)
	if err != nil {
		ErrorLogger.Println("Failed to make Comfort Rating service call: \n", err)
		return nil, fmt.Errorf("Failure in Route Analysis Aggregator: \n%v", err)
	} else {
		DebugLogger.Println("Successfully made service call to Comfort Service.")
		connCS.Close()
	}

	fmt.Println(responseMessageCS)

	// Create the response message
	responseMessage := serverPB.AnalysisResponse{
		UnixTime: request.UnixTime,
		AveragePower: responseMessagePTS.PowerEstimateAverage,
		TotalCost: responseMessagePTS.TotalCost,
		AverageRmsX: responseMessagePVS.RmsVibrationX,
		AverageRmsY: responseMessagePVS.RmsVibrationY,
		AverageRmsZ: responseMessagePVS.RmsVibrationZ,
		// ComfortLevel: 
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