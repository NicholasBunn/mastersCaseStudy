using System;
using Xunit;
using vesselMotionService;

public class TestVesselMotionService
{
    [Fact]
    public void TestCalculateOpenWaterRespone()
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
    public void TestCalculatOpenWaterResponse_NegativeInput(float portPropMotorPower, float latitude, float relativeWindSpeed, float relativeWindDirection, float heading, float waveHeight)
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