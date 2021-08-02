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
	"github.com/go-yaml/yaml"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	// Proto packages
	oceanWeatherServicePB "github.com/NicholasBunn/mastersCaseStudy/services/routeAnalysisAggregator/proto/v1/generated/oceanWeatherService"
	powerTrainServicePB "github.com/NicholasBunn/mastersCaseStudy/services/routeAnalysisAggregator/proto/v1/generated/powerTrainService"
	vesselMotionServicePB "github.com/NicholasBunn/mastersCaseStudy/services/routeAnalysisAggregator/proto/v1/generated/vesselMotionService"
	comfortServicePB "github.com/NicholasBunn/mastersCaseStudy/services/routeAnalysisAggregator/proto/v1/generated/comfortService"
	serverPB "github.com/NicholasBunn/mastersCaseStudy/services/routeAnalysisAggregator/proto/v1/generated/routeAnalysisAggregator"
)

var (
	// Addresses (To be passed in a config file)
	addrMyself string
	addrOWS	string
	addrPTS string
	addrVMS string

	timeoutDuration     int           // The time, in seconds, that the client should wait when dialing (connecting to) the server before throwing an error
	callTimeoutDuration time.Duration // The time, in seconds, that the client should wait when making a call to the server before throwing an error

	// JWT stuff, load this in from config
	secretkey     string
	tokenduration time.Duration

	accessibleRoles map[string][]string // This is a map of service calls with their required permission levels

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

	addrMyself = os.Getenv("ROUTEANALYSISHOST") + ":" + config.Server.Port.Myself
	addrOWS = os.Getenv("OWSHOST") + ":" + config.Client.Port.OceanWeatherService
	addrPTS = os.Getenv("PTSHOST") + ":" + config.Client.Port.PowerTrainService
	addrVMS = os.Getenv("VMSHOST") + ":" + config.Client.Port.VesselMotionService
	addrCS = os.Getenv("CSHOST") + ":" + config.Client.Port.ComfortService

	// Load timeouts from config
	timeoutDuration = config.Client.Timeout.Connection
	callTimeoutDuration = time.Duration(config.Client.Timeout.Call) * time.Second

	// Load JWT parameters from config
	secretkey = config.Server.Authentication.Jwt.SecretKey
	tokenduration = time.Duration(config.Server.Authentication.Jwt.TokenDuration) * (time.Minute)

	accessibleRoles = map[string][]string{
		config.Server.Authentication.AccessLevel.Name.AnalyseRoute: config.Server.Authentication.AccessLevel.Role.AnalyseRoute,
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
	/* The main function sets up a server to listen on the specified port,
	encrypts the server connection with TLS, and registers the services on
	offer */

	arbContext, cancel := context.WithTimeout(context.Background(), callTimeoutDuration)
	defer cancel()

	analysisRequest := serverPB.AnalysisRequest{
		UnixTime: 1608580803,
		Latitude: 58.7984,
		Longitude: 17.8081,
		Heading: 15,
		PropPitch: 0.26854406323815316,
		MotorSpeed: 0.597549569477592,
		SOG: 0.030389908256880732,
	}

	response, err := AnalyseRoute(arbContext, &analysisRequest)

	fmt.Println(response, err)

	InfoLogger.Println("Started route analysis aggregator.")

	// Create a listener on the specified tcp port
	listener, err := net.Listen("tcp", addrMyself)
	if err != nil {
		ErrorLogger.Fatalf("Failed to listen on port %v: \n%v", addrMyself, err)
	}
	InfoLogger.Println("Listening on port: ", addrMyself)

	// Create a gRPC server object
	analysisServer := grpc.NewServer()

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
	// 			PrepareService    string `yaml:"prepare"`
	// 			EstimationService string `yaml:"estimation"`
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

type server struct {
	// Use this to implement the power estimation service package

	serverPB.UnimplementedAnalysisServiceServer
}

// ________IMPLEMENT THE OFFERED SERVICES________

func (s *server) AnalyseRoute(ctx context.Context, request *serverPB.AnalysisRequest) (*serverPB.AnalysisResponse, error) {
	/* This call provides foresight for tactical decision-making by providing a summary of the provided route.
	 */

	InfoLogger.Println("Received Analyse Route service call.")

	// ________Query Ocean Weather Service________

	// Create an insecure connection to the ocean weather service server
	connOWS, err := createInsecureServerConnection(
		addrOWS, // Set the address of the server
		timeoutDuration, // Set the duration that the clinet will wait before timing out
	)
	if err != nil {
		return nil, err
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
		ErrorLogger.Println("Failed to make Ocean Weather Prediction service call: ")
		return nil, err
	} else {
		DebugLogger.Println("Successfully made service call to Ocean Weather Service.")
		connOWS.Close()
	}

	// ________Query Power Train Service________
	
	// Create an insecure connection to the power train service server
	connPTS, err := createInsecureServerConnection(
		addrPTS,
		timeoutDuration,
	)
	if err != nil {
		return nil, err
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

	requestMessagePTS.WindDirectionRelative, err = calculateRelativeWindDirection(responseMessageOWS.WindDirection, request.Heading)
	if err != nil {
		ErrorLogger.Println("Failed to calculate relative wind direction: ")
		return nil, err
	}

	DebugLogger.Println("Succesfully created a Power Train Estimate Request.")

	InfoLogger.Println("Making Power Estimate Call.")
	ptsContext, cancel := context.WithTimeout(context.Background(), callTimeoutDuration)
	defer cancel()

	// Invoke the Power Train Service
	responseMessagePTS, err := clientPTS.PowerEstimate(ptsContext, &requestMessagePTS)
	if err != nil {
		ErrorLogger.Println("Failed to make Power Estimate service call: ")
		return nil, err
	} else {
		DebugLogger.Println("Successfully made service call to Power Train Service.")
		connPTS.Close()
	}

	// ________Query Vessel Motion Service________
	
	// Create an insecure connection to the power train service server
	connPTS, err := createInsecureServerConnection(
		addrVMS,
		timeoutDuration,
	)
	if err != nil {
		return nil, err
	}

	InfoLogger.Println("Creating Vessel Motion Service client.")
	clientVMS := vesselMotionServicePB.NewVesselMotionServiceClient(connVMS)
	DebugLogger.Println("Succesfully created client connection to Vessel Motion Service.")

	// Create a motion estimate request message
	requestMessageVMS := vesselMotionServicePB.MotionEstimateRequest{
		portPropMotorPower: responseMessagePTS.PowerEstimate,
		eindSpeedRelative: (responseMessageOWS.WindSpeed - request.SOG),
		latitude: request.Latitude,
		heading, request.Heading,
		waveHeight: responseMessageOWS.SwellHeight,
		WindDirectionRelative: requestMessagePTS.relativeWindDirection,
		ModelType: vesselMotionServicePB.ModelTypeEnum_OPENWATER,
	}

	DebugLogger.Println("Succesfully created a Motion Estimate Request.")

	InfoLogger.Println("Making Motion Estimate Call.")
	vmsContext,cancel := context.WithTimeout(context.Background(), callTimeoutDuration)
	defer cancel()

	// Invoke the Vessel Motion Service
	responseMessageVMS, err := clientVMS.MotionEstimate(vmsContext, &responseMessageVMS)
	if err != nil {
		ErrorLogger.Println("Failed to make Motion Estimate service call: ")
		return nil, err
	} else {
		DebugLogger.Println("Successfully made service call to Vessel Motion Service.")
		connVMS.Close()
	}

	// ________Query Comfort Service________
	
	// Create an insecure connection to the power train service server
	connCS, err := createInsecureServerConnection(
		addrCS,
		timeoutDuration,
	)
	if err != nil {
		return nil, err
	}

	InfoLogger.Println("Creating Comfort Service client.")
	clientCS := comfortServicePB.NewComfortServiceClient(connCS)
	DebugLogger.Println("Succesfully created client connection to Comfort Service.")

	// Create a comfort request message
	requestMessageCS := comfortServicePB.ComfortRequest{
		unixTime: request.UnixTime,
		humanWeightedVibrationX: responseMessageVMS.AccelerationEstimateX,
		humanWeightedVibrationY: responseMessageVMS.AccelerationEstimateY,
		humanWeightedVibrationZ: responseMessageVMS.AccelerationEstimateZ,
	}

	DebugLogger.Println("Succesfully created a Comfort Request.")

	InfoLogger.Println("Making Comfort Rating Call.")
	csContext,cancel := context.WithTimeout(context.Background(), callTimeoutDuration)
	defer cancel()

	// Invoke the Vessel Motion Service
	responseMessageCS, err := clientCS.ComfortRating(csContext, &requestMessageCS)
	if err != nil {
		ErrorLogger.Println("Failed to make Comfort Rating service call: ")
		return nil, err
	} else {
		DebugLogger.Println("Successfully made service call to Comfort Service.")
		connVMS.Close()
	}

	return nil, status.Errorf(codes.Unimplemented, "method AnalyseRoute not implemented")
}

// ________SUPPORTING FUNCTIONS________

func calculateRelativeWindDirection(windDirection []float32, heading []float32) ([]float32, error) {
	/* This function takes the wind direction and vessel heading as inputs. Using these, it 
	calculates and returns the wind direction relative to the vessels direction.
	*/

	var relativeWindDirection []float32
	var tempRelativeWindDirection float32

	for count, windDirectionInstance := range windDirection {
		tempRelativeWindDirection = windDirectionInstance - heading[count]

		// Check whether the relative wind direction is negative and add 360 degrees until it is positive so that all directions returned are on the same coordinate system.
		for tempRelativeWindDirection < 0 {
			tempRelativeWindDirection += 360
		}
		relativeWindDirection = append(relativeWindDirection, tempRelativeWindDirection)
	}

	return relativeWindDirection, nil
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
