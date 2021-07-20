package main

import (
	oceanWeatherServicePB "github.com/nicholasbunn/mastersCaseStudy/services/routeAnalysisAggregator/proto/v1/generated/oceanWeatherService"
	serverPB "github.com/nicholasbunn/mastersCaseStudy/services/routeAnalysisAggregator/proto/v1/generated/routeAnalysisAggregator"
)

func init() {
	/* The init function is used to load in configuration variables, and set up the logger and metric interceptors whenever the service is started
	 */

	// ________CONFIGURATION________
	// Load YAML configurations into config struct
	config, _ := DecodeConfig("src/powerEstimationSP/configuration.yaml")

	addrMyself = os.Getenv("ROUTEANALYSISHOST") + ":" + config.Server.Port.Myself
	addrOWS = os.Getenv("OWSHOST") + ":" + config.Client.Port.OceanWeatherService

	// Load timeouts from config
	timeoutDuration = config.Client.Timeout.Connection
	callTimeoutDuration = time.Duration(config.Client.Timeout.Call) * time.Second

	// Load JWT parameters from config
	secretkey = config.Server.Authentication.Jwt.SecretKey
	fmt.Println(secretkey)
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
	// 			PrepareService    string `yaml:"prepare"`
	// 			EstimationService string `yaml:"estimation"`
	// 		} `yaml:"port"`
			Timeout struct {
				Connection int `yaml:"connection"`
				Call       int `yaml:"call"`
	// 		} `yaml:"timeout"`
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
	// 	} `yaml:"client"`
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

	// Create an insecure connection to the ocean weather service server
	connOWS, err := createInsecureServerConnection(
		addrOWS, // Set the address of the server
		timeoutDuration, // Set the duration that the clinet will wait before timing out
	)
	if err != nil {
		return nil, err
	}
	
	return nil, status.Errorf(codes.Unimplemented, "method AnalyseRoute not implemented")
}

// ________SUPPORTING FUNCTIONS________
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
