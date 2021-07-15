#include "processVibrationService.cpp"
#include "proto/v1/generated/process_vibration_service_api_v1_mock.grpc.pb.h"
#include <gtest/gtest.h>

// FOR MORE INFO, VISIT https://www.eriksmistad.no/getting-started-with-google-test-on-ubuntu/ or https://grpc.github.io/grpc/cpp/md_doc_unit_testing.html

TEST(ProcessEstimateTest, StandardOperation) 
{
    // ASSERT_EQ(0, functionName());
}

class MockedServer {
    public:
    explicit MockedServer(ProcessVibrationService::StubInterface *stub) : stub_(stub) {}

    void DoProcessEstimate() {
        ClientContext context;
        ProcessRequest request;
        ProcessResponse response;
        request->set_vibration_x(0, );
        request->set_vibration_y(0, );
        request->set_vibration_z(0, );
        Status s = stub_->ProcessEstimate(&context, request, &response);

        std::cout << response << std::endl;
        // EXPECT_EQ(request.message(), response.message());
        // EXPECT_TRUE(s.ok());
    }
}

int main(int argc, char **argv)
{
    testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}