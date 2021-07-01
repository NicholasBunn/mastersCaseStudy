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

    enum evaluationCase {Passenger, Crew, Work};

    void EvaluateHabitability (float acceleration, evaluationCase area, ComfortResponse* response)
    {
        /* This function takes an acceleration, an analysis area, and the response message it will be modifying as inputs. It evaluates the habitibility of a vessel using the acceleration input, based on what zone it falls within according to ISO 6954:2000.
        */

        int upperLimit;
        int lowerLimit;

        // Set the upper and lower limits for evaluation, based on the selected evaluation area.
        switch (area)
        {
            case Passenger:
                upperLimit = 4;
                lowerLimit = 2;
                break;
            case Crew:
                upperLimit = 6;
                lowerLimit = 3;
                break;
            case Work:
                upperLimit = 8;
                lowerLimit = 4;
                break;
            default:
                // Default to the most sensitive case (passenger cabins)
                upperLimit = 4;
                lowerLimit = 2;
                break;
        }
    
        // Evaluate habitibility
        if (acceleration < (float) lowerLimit)
        {
            // If acceleration is less than the lower limit, adverse commetns are unlikely.
            response->set_rating(comfortServiceAPI::v1::ADVERSE_COMMENTS_UNLIKELY);
        } else if (acceleration < upperLimit)
        {
            // Acceleration in the zone between the lower and upper limits is considered to be commonly accepted operating conditions.
            response->set_rating(comfortServiceAPI::v1::NORMAL_OPERATING_CONDITIONS);

        } else
        {
            // If acceleration is greater than the upper limit, adverse comments are probable.
            response->set_rating(comfortServiceAPI::v1::ADVERSE_COMMENTS_PROBABLE);
        }

    }

    Status ComfortRating(
        ServerContext* context, 
        const ComfortRequest* request, 
        ComfortResponse* response
    ) override {
        /* The 'Comfort Rating' call provides foresight for ?? decision-making by providing a comfort rating for a proposed route, based on estimated vibrations on board.
        */

        int humanWeightedVibrationX;
        int humanWeightedVibrationY;
        int humanWeightedVibrationZ;

        float iterationMax = 0;
        float max = 0;   

        // Iterate through the request fields
        for(int i = 0; i < request->unix_time_size(); i++)
        {
            // First, add the timestamp used to the response message
            response->set_unix_time(i, request->unix_time(i));

            // Extract the human-weighted vibrations from the request
            humanWeightedVibrationX = request->human_weighted_vibration_x(i);
            humanWeightedVibrationY = request->human_weighted_vibration_y(i);
            humanWeightedVibrationZ = request->human_weighted_vibration_z(i);

            // Find the maximum vibration among the current x, y, and z
            iterationMax = std::max({humanWeightedVibrationX, humanWeightedVibrationY, humanWeightedVibrationZ});

            // Check if this local maximum is greater than the current maximum and, if it is, re-evaluate the habitability
            if(iterationMax > max)
            {
                EvaluateHabitability (iterationMax, evaluationCase::Passenger, response);
            }
        }

        return Status::OK;
    }

    Status VDVEstimate(
        ServerContext* context, 
        const VDVRequest* request, 
        VDVResponse* response
    ) override {
        /* The 'VDV Estimate' call provides ?? for ?? decision-making by providing the estimated vibration dose value for a requested route.
        */

        // 1. Calculate time between samples
        
        // 2. Calculate estamated VDV

        // 3. Increment equivalent vibration (use denominator and numerator)

        std::cout << "BLABLA" << std::endl;
        return Status::CANCELLED;
    }

    // Status VDVTracking(>
    //     ServerContext* context, 
    //     const VDVRequest* request, 
    //     VDVResponse* response
    // ) override {
    //     /* The 'VDV Tracking' call provides ?? for ?? decision-making by providing a real-time vibration dose value based on accelerometer feeds on board the vessel.
    //     */
       
    //     std::cout << "BLABLA" << std::endl;
    //     return Status::CANCELLED;
    // }

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