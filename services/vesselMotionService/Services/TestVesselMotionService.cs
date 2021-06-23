using Xunit;
using vesselMotionService;

public class TestVesselMotionService
{
    VesselMotionService testInstance = new VesselMotionService(null);

    [Fact]
    public void TestCalculateOpenWaterRespone()
    {
        float result = testInstance.CalculateOpenWaterResponse(2431.24975585937F, -35.53165F, 13.5F, 337F, 212F, 1.5F);
        Assert.Equal(result, 7.0642624F);
    }
}