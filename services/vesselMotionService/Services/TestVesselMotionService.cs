using System;
using Xunit;
using vesselMotionService;
using System.Threading.Tasks;

public class VesselMotionServiceUnitTest
{
    /* This class is used to execute all unit tests on the Vessel Motion Service. Put any tests used to verify functions in this class.
    */

    [Fact]
    public void Test_CalculateOpenWaterResponse()
    {
        /* This tests that the 'CalculateOpenWater' function is functioning as expected.
        */

        System.Console.WriteLine("Testing Vessel Motion Service: Unit Test: Calculate Open Water Response (Function Test)");

        // Create an instance of the VesselMotionService class
        VesselMotionService testInstance = new VesselMotionService(null);

        // Call the CalcuateOpenWaterResponse function and check the result
        (float xResponse, float yResponse, float zResponse) = testInstance.CalculateOpenWaterResponse(2431.24975585937F, -35.53165F, 13.5F, 337F, 212F, 1.5F);

        Assert.Equal(11.090744F, xResponse);
        Assert.Equal(7.6717124F, yResponse);
        Assert.Equal(7.0642624F, zResponse);
    }

    [Theory]
    [InlineData(-2431.24975585937F, -35.53165F, 13.5F, 337F, 212F, 1.5F)]
    [InlineData(2431.24975585937F, -35.53165F, -13.5F, 337F, 212F, 1.5F)]
    [InlineData(2431.24975585937F, -35.53165F, 13.5F, -337F, 212F, 1.5F)]
    [InlineData(2431.24975585937F, -35.53165F, 13.5F, 337F, -212F, 1.5F)]
    [InlineData(2431.24975585937F, -35.53165F, 13.5F, 337F, 212F, -1.5F)]
    [InlineData(2431.24975585937F, -35.53165F, -13.5F, 337F, -212F, -1.5F)]
    [InlineData(-2431.24975585937F, -35.53165F, -13.5F, -337F, -212F, -1.5F)]
    public void Test_CalculatOpenWaterResponse_NegativeInput(float portPropMotorPower, float latitude, float relativeWindSpeed, float relativeWindDirection, float heading, float waveHeight)
    {
        /* This tests that error handling of the 'CalculateOpenWater' function, ensuring that it correctlty throws an error when a negative value is passed in for any input other than latitude.
        */

        System.Console.WriteLine("Testing Vessel Motion Service: Unit Test: Calculate Open Water Response - negative input (Function Test)");

        // Create an instance of the VesselMotionService class
        VesselMotionService testInstance = new VesselMotionService(null);

        try
        {
            var result = testInstance.CalculateOpenWaterResponse(portPropMotorPower, latitude, relativeWindSpeed, relativeWindDirection, heading, waveHeight);
        }
        catch (Exception err)
        {
            Assert.Equal(typeof(ArgumentException), err.GetType());
        }

    }

}

public class VesselMotionServiceIntegrationTest
{
    /* This class is used to execute all integration tests on the Vessel Motion Service. Put any tests used to verify the gRPC/server implementation in this class.
    */

    [Theory]
    [InlineData(2431.24975585937F, -35.53165F, 13.5F, 337F, 212F, 1.5F, vesselMotionService.ModelTypeEnum.Openwater)]
    [InlineData(2431.24975585937F, -35.53165F, 13.5F, 337F, 212F, 1.5F, vesselMotionService.ModelTypeEnum.UnknownModel)]
    [InlineData(2431.24975585937F, -35.53165F, 13.5F, 337F, 212F, 1.5F, null)]
    public async Task Test_MotionEstimate(float portPropMotorPower, float latitude, float relativeWindSpeed, float relativeWindDirection, float heading, float waveHeight, vesselMotionService.ModelTypeEnum modelType)
    {
        /* This function tests the MotionEstimate-specific functionality. It ensures that the service call selects the correct model and makes the correct function calls to serve the request.
        */

        System.Console.WriteLine("Testing Vessel Motion Service: Integration Test: Motion Estimate (Service Call Test)");

        // Create an instance of the VesselMotionService class
        VesselMotionService testInstance = new VesselMotionService(null);

        // Create motion estimate request message
        var testRequestMessage = new MotionEstimateRequest();
        testRequestMessage.PortPropMotorPower.Add(portPropMotorPower);
        testRequestMessage.WindSpeedRelative.Add(relativeWindSpeed);
        testRequestMessage.Latitude.Add(latitude);
        testRequestMessage.Heading.Add(heading);
        testRequestMessage.WaveHeight.Add(waveHeight);
        testRequestMessage.WindDirectionRelative.Add(relativeWindDirection);
        testRequestMessage.ModelType = modelType;

        var response = await testInstance.MotionEstimate(testRequestMessage, null);
        Assert.Equal(11.090744F, response.AccelerationEstimateX[0]);
        Assert.Equal(7.6717124F, response.AccelerationEstimateY[0]);
        Assert.Equal(7.064262390136719F, response.AccelerationEstimateZ[0]);
    }
}