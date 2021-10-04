package interceptors

import (
	// Native packages

	"context"
	"log"
	"os"
	"strings"

	// gRPC packages
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	// Personal packages
	"github.com/NicholasBunn/mastersCaseStudy/generalComponents/authenticationStuff"
)

var (
	callCounter = 0 // Quick work around from refactoring with a class (the end is near)
// 	// Logging stuff
// 	DebugLogger   *log.Logger
// 	InfoLogger    *log.Logger
// 	WarningLogger *log.Logger
// 	ErrorLogger   *log.Logger
)

func init() {
	/* The init function is used to set up the logger and metric interceptors whenever the service is started
	*/

	// If the file doesn't exist, create it, otherwise append to the file
	pathSlice := strings.Split(os.Args[0], "/") // This just extracts the services name (filename)
	file, err := os.OpenFile("program logs/"+pathSlice[len(pathSlice)-1]+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		// If opening the log file throws an error, continue to create the loggers but print to terminal instead
		log.Println("Unable to initialise log file, good luck :)")
	} else {
		log.SetOutput(file)
	}

	callCounter = 0

	DebugLogger = log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}

type ServerRateLimitStruct struct {
	CallLimit           int
}

func (interceptor *ServerRateLimitStruct) ServerRateLimitInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	InfoLogger.Println("Starting server-side rate limit interceptor")

	callCounter += 1
	
	if (callCounter > interceptor.CallLimit) {
		return nil, status.Errorf(codes.ResourceExhausted, "%s is rejected by grpc_ratelimit middleware, please retry later.", info.FullMethod)
	} else {
		response = handler(ctx, req)
		callCounter -= 1

		return response
	}
}