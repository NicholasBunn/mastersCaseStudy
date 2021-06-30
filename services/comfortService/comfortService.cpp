#include <string>

#include <grpcpp/grpcpp.h>
#include <grpc/grpc.h>
#include <grpcpp/server.h>
#include <grpcpp/server_builder.h>
#include <grpcpp/server_context.h>
#include "proto/v1/generated/comfort_service_api_v1.pb.h"
#include "proto/v1/generated/comfort_service_api_v1.grpc.pb.h"

using grpc::Server;
using grpc::ServerBuilder;
using grpc::ServerContext;
using grpc::Status;

using comfortServiceAPI::v1::ComfortService;
using comfortServiceAPI::v1::ComfortRequest;
using comfortServiceAPI::v1::ComfortResponse;
using comfortServiceAPI::v1::VDVRequest;
using comfortServiceAPI::v1::VDVResponse;

class ComfortServiceImplementation final : public ComfortService::Service {
    /* 'Comfort Service' offers three service calls that provide information about human comfort onboard, in response to vessel vibrations.
    */

    Status ComfortRating(
        ServerContext* context, 
        const ComfortRequest* request, 
        ComfortResponse* response
    ) override {
        /* The 'Comfort Rating' call provides foresight for ?? decision-making by providing a comfort rating for a proposed route, based on estimated vibrations on board.
        */
        // int a = request->a();
        // int b = request->b();

        // response->set_result(a * b);
        std::cout << "BLABLA" << std::endl;
        return Status::OK;
    }

    Status VDVEstimate(
        ServerContext* context, 
        const VDVRequest* request, 
        VDVResponse* response
    ) override {
        /* The 'VDV Estimate' call provides ?? for ?? decision-making by providing the estimated vibration dose value for a requested route.
        */

        std::cout << "BLABLA" << std::endl;
        return Status::CANCELLED;
    }

    Status VDVTracking(
        ServerContext* context, 
        const VDVRequest* request, 
        VDVResponse* response
    ) override {
        /* The 'VDV Tracking' call provides ?? for ?? decision-making by providing a real-time vibration dose value based on accelerometer feeds on board the vessel.
        */
       
        std::cout << "BLABLA" << std::endl;
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
    ComfortServiceImplementation service;

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