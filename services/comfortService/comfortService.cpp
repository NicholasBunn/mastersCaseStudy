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
    
    void AssessComfort (double equivalentAcceleration, ComfortResponse* response)
    {
        if(equivalentAcceleration < 0.315)
        {
            response->set_rating(comfortServiceAPI::v1::NOT_UNCOMFORTABLE);
        } else if (equivalentAcceleration < 0.6)
        {
            response->set_rating(comfortServiceAPI::v1::SLIGHTLY_UNCOMFORTABLE);
        } else if (equivalentAcceleration < 0.9)
        {
            response->set_rating(comfortServiceAPI::v1::FAIRLY_UNCOMFORTABLE);
        } else if (equivalentAcceleration < 1.4)
        {
            response->set_rating(comfortServiceAPI::v1::UNCOMFORTABLE);
        } else if (equivalentAcceleration < 2)
        {
            response->set_rating(comfortServiceAPI::v1::VERY_UNCOMFORTABLE);
        } else
        {
            response->set_rating(comfortServiceAPI::v1::EXTREMELY_UNCOMFORTABLE);
        }
    }

    double calculateEquivalentVibration(double timeStamps[], double weightedVibration[])
    {
        /* This function takes a time series of weighted vibration data as well as their assosciated timestamps. Using these, it calculates the equivalent vibration magnitude according to SANS 2631-1 (Appendix C, equ C.1).
        */

        float timeGap = 0; // This variable holds the time difference between vibration samples
        double numerator = 0; // This variables holds the accumulated numerator use in the calculation of the equivalent vibration magnitude
        double denominator = 0; // This variables holds the accumulated denominator use in the calculation of the equivalent vibration magnitude
        
        // Iterate through the request samples
        for(int i = 1; i < sizeof(weightedVibration); i++)
        {
            // Find the time difference between estimates. Assume the acceleration signal is constant over this time period
            timeGap = timeStamps[i] - timeStamps[i-1];

            // Increment numerator of the equivalent vibration equation
            numerator += ( std::pow(weightedVibration[i],2)*timeGap );

            // Increment denominator of the equivalent vibration equation
            denominator += timeGap;
        }

        // Calculate and return equivalent vibration
        return sqrt(numerator/denominator);
    }

    Status ComfortRating(
        ServerContext* context, 
        const ComfortRequest* request, 
        ComfortResponse* response
    ) override {
        /* The 'Comfort Rating' call provides foresight for ?? decision-making by providing a comfort rating for a proposed route, based on estimated vibrations on board. 
        */

        double timeStampArray[request->unix_time_size()];
        double vibrationXArray[request->human_weighted_vibration_x_size()];

        // Unpack and add all timestamps to the response message
        for(int i = 0; i < request->unix_time_size(); i++)
        {
            timeStampArray[i] = request->unix_time(i); // Unpack the request message into a local array
            response->set_unix_time(i, timeStampArray[i]); // Set response message time
        }

        // Unpack all human weighted vibrations
        for(int i = 0; i < request->human_weighted_vibration_x(); i++_)
        {
            vibrationXArray[i] = request->human_weighted_vibration_x(i);
        }

        // ________Calculate equivalent vibration________

        double equivalentVibration = calculateEquivalentVibration(timeStampArray, vibrationXArray);

        // 2. Calculate VDV
        double vdvAccelerationZ = 1.4 * equivalentVibration * std::pow((timeStampArray[sizeof(timeStampArray) - 1] - timeStampArray[0]), 0.25);

        // 3. Assess comfort
        AssessComfort(equivalentVibration, response);

        return Status::OK;
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