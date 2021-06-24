using System;
using Xunit;
using vesselMotionService;
using System.Threading.Tasks;

public class VesselMotionServiceUnitTest
{
    [Fact]
    public void Test_CalculateOpenWaterResponse()
    {
        /* This tests that the 'CalculateOpenWater' function is functioning as expected.
        */

        System.Console.WriteLine("Testing Vessel Motion Service: Calculate Open Water Response (Function Test)");

        // Create an instance of the VesselMotionService class
        VesselMotionService testInstance = new VesselMotionService(null);

        // Call the CalcuateOpenWaterResponse function and check the result
        Assert.Equal(testInstance.CalculateOpenWaterResponse(2431.24975585937F, -35.53165F, 13.5F, 337F, 212F, 1.5F), 7.0642624F);
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
        /* This tests that the 'CalculateOpenWater' function correctlty throws an error when a negative value is passed in for 
        any input other than latitude
        */

        System.Console.WriteLine("Testing Vessel Motion Service: Calculate Open Water Response - negative input (Function Test)");

        // Create an instance of the VesselMotionService class
        VesselMotionService testInstance = new VesselMotionService(null);

        try
        {
            var result = testInstance.CalculateOpenWaterResponse(portPropMotorPower, latitude, relativeWindSpeed, relativeWindDirection, heading, waveHeight);
        }
        catch (Exception err)
        {
            Assert.Equal(err.GetType(), typeof(ArgumentException));
        }

    }

}

public class VesselMotionServiceIntegrationTest
{

    [Fact]
    public async Task Test_MotionEstimate()
    {

        System.Console.WriteLine("Testing Vessel Motion Service: Motion Estimate (Service Call Test)");

        // Create an instance of the VesselMotionService class
        VesselMotionService testInstance = new VesselMotionService(null);

        // Create motion estimate request message
        var testRequestMessage = new MotionEstimateRequest();
        testRequestMessage.PortPropMotorPower.Add(2431.24975585937F);
        testRequestMessage.WindSpeedRelative.Add(13.5F);
        testRequestMessage.Latitude.Add(-35.53165F);
        testRequestMessage.Heading.Add(212F);
        testRequestMessage.WaveHeight.Add(7.0642624F);
        testRequestMessage.WindDirectionRelative.Add(337F);
        testRequestMessage.ModelType = vesselMotionService.ModelTypeEnum.Openwater;

        var response = await testInstance.MotionEstimate(testRequestMessage, null);

        Assert.Equal(response.Acceleration[0], 7.064262390136719F);
    }
}