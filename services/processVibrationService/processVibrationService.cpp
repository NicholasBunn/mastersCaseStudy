#include <string>
#include <math.h>

#include <grpcpp/grpcpp.h>
#include <grpc/grpc.h>
#include <grpcpp/server.h>
#include <grpcpp/server_builder.h>
#include <grpcpp/server_context.h>
#include "proto/v1/generated/process_vibration_service_api_v1.pb.h"
#include "proto/v1/generated/process_vibration_service_api_v1.grpc.pb.h"

using grpc::Server;
using grpc::ServerBuilder;
using grpc::ServerContext;
using grpc::Status;

using processVibrationServiceAPI::v1::ProcessVibrationService;
using processVibrationServiceAPI::v1::ProcessRequest;
using processVibrationServiceAPI::v1::ProcessResponse;

class ProcessVibrationServiceImplementation final : public ProcessVibrationService::Service
{
    Status ProcessEstimates(
        ServerContext* context,
        const ProcessRequest* request,
        ProcessResponse* response
    ) {

        std::cout << "Received a Process Estimates service call" << std::endl;
        
        float k = 1.4;
        float timeDifference;

        // Set the initial responses before entering the iteration
        response->set_unix_time(0, request->unix_time(0));
        timeDifference = request->unix_time(1) - request->unix_time(0);

        // Calculate inital RMS
        response->set_rms_vibration_x(0, (request->vibration_x(0))/sqrt(2));
        response->set_rms_vibration_y(0, (request->vibration_y(0))/sqrt(2));
        response->set_rms_vibration_z(0, (request->vibration_z(0))/sqrt(2));

        // Iterate through the request fields
        for(int i = 1; i <= request->unix_time_size(); i++)
        {
            response->set_unix_time(i, request->unix_time(i));
            timeDifference = request->unix_time(i) - request->unix_time(i-1);

            // ________Calculate RMS (in m/s^2)________
            response->set_rms_vibration_x(i, (request->vibration_x(i))/sqrt(2));
            response->set_rms_vibration_y(i, (request->vibration_y(i))/sqrt(2));
            response->set_rms_vibration_z(i, (request->vibration_z(i))/sqrt(2));
        }

        return Status::OK;
    }

    Status ProcessTracking(
        ServerContext* context,
        const ProcessRequest* request,
        ProcessResponse* response
    ) {
        return Status::CANCELLED;
    }
};

void Serve() {
    /* This function creates a server with specified interceptors, registers the service calls offered by that server, and exposes
	the server over a specified port.
	*/

    // Set the address for the service to host itself on
    std::string address("0.0.0.0:5000"); //ToDo get this from config

    // Add the hosted services 
    ProcessVibrationServiceImplementation service;

    // Create a builder instance for the server
    ServerBuilder builder;

    // Add an insecure port to the server builder
    builder.AddListeningPort(address, grpc::InsecureServerCredentials());
    // Register the hosted services to the server (builder)
    builder.RegisterService(&service);

    // Build and start the server, and listen for calls on the specified port
    std::unique_ptr<Server> server(builder.BuildAndStart());
    std::cout << "Server listening on port: " << address << std::endl;

    // Defer termination for a 'persistent' service
    server->Wait();
}

int main(int argc, char** argv) {
    Serve();

    return 0;
}