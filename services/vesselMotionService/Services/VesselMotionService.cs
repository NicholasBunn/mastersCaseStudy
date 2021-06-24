using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Grpc.Core;
using Microsoft.Extensions.Logging;

namespace vesselMotionService
{
    
    public class VesselMotionService : vesselMotionService.vesselMotionServiceBase
    {
        /* 'Vessel Motion Service; offers three service calls that provide information 
        about the vessel's motion in response to her sailing conditions (mainly regarding 
        the high frequency component of the vessel's acceleration/vibration)
        */
        
        private readonly ILogger<VesselMotionService> _logger;
        public VesselMotionService(ILogger<VesselMotionService> logger)
        {
            _logger = logger;
        }

        public override Task<MotionResponse> MotionEstimate(MotionEstimateRequest request, ServerCallContext context)
        {
            /* The 'Motion Estimate' call provides foresight for tactical decision-making by providing high-frequency 
            acceleration estimates for a requested sailing conditions at a requested location on the ship 
            */

            // _logger.LogInformation("Received Motion Estimate Service Call");

            // Create a response object            
            var response = new MotionResponse();  

            try
            {
                // Select the correct model to call (openwater vs ice). Defaults to open water
                switch(request.ModelType)
                {
                    case ModelTypeEnum.Openwater: 
                        // Iterate through the provided inputs and produce estimates for each set
                        for(int i = 0; i < request.PortPropMotorPower.Count; i++)
                        {   
                            // For the current set of input variables, add the estimate to the response
                            response.Acceleration.Add(CalculateOpenWaterResponse(request.PortPropMotorPower[i], request.Latitude[i], request.WindSpeedRelative[i], request.WindDirectionRelative[i], request.Heading[i], request.WaveHeight[i]));
                        }

                        // _logger.LogInformation("Succesfully calculated response for open water model");

                        break;

                    case ModelTypeEnum.Ice:
                        throw new RpcException(new Status(StatusCode.Unimplemented, "Ice estimation logic has not been implemented yet"));

                    case ModelTypeEnum.UnknownModel:
                        // Iterate through the provided inputs and produce estimates for each set
                        for(int i = 0; i < request.PortPropMotorPower.Count; i++)
                        {
                            response.Acceleration.Add(CalculateOpenWaterResponse(request.PortPropMotorPower[i], request.Latitude[i], request.WindSpeedRelative[i], request.WindDirectionRelative[i], request.Heading[i], request.WaveHeight[i]));
                        }

                        // _logger.LogInformation("Succesfully calculated response for unknown model (defaulted to open water)");
                        break;
                    
                    default:   
                        // Iterate through the provided inputs and produce estimates for each set                         
                        for(int i = 0; i < request.PortPropMotorPower.Count; i++)
                        {
                            response.Acceleration.Add(CalculateOpenWaterResponse(request.PortPropMotorPower[i], request.Latitude[i], request.WindSpeedRelative[i], request.WindDirectionRelative[i], request.Heading[i], request.WaveHeight[i]));
                        }

                        // _logger.LogInformation("Succesfully calculated response for no provided model (defaulted to open water)");
                        break;
                }
            }
            catch (RpcException err) when (err.StatusCode == StatusCode.Internal)
            {
                _logger.LogDebug($"Internal Error, need to actually interpret this error properly: \n {err}");
                throw err;
            }
            catch (RpcException err)
            {
                _logger.LogDebug($"Unaccounted for error, please contact a developer and let them know about this \n {err}");
            }
            
            return Task.FromResult(response);
        }

        public override Task<MotionResponse> MotionTracking(MotionTrackingRequest request, ServerCallContext context)
        {
            /* The 'Motion Tracking' call provides insight for tactical decision-making by providing real-time, 
            high-frequency acceleration readings for a requested location on the ship
            */

            return base.MotionTracking(request, context);
        }

        public override Task<MotionEvaluationResponse> MotionEstimateEvaluation(MotionEstimateRequest request, ServerCallContext context)
        {
            /* The 'Motion Estimation Evaluation' call provides hindsight for strategic decision-making by evaluating the accuracy of the models predictions
            */

            return base.MotionEstimateEvaluation(request, context);
        }

        internal float CalculateOpenWaterResponse(float portPropMotorPower, float latitude, float relativeWindSpeed, float relativeWindDirection, float heading, float waveHeight)
        {
            /* This function calculates the vibration response of the vessel for open water sailing conditions. In this implementation, only the y-axis acceleration in the bridge is required, purely as a means to demonstrate the design's ability to aggregate and coordinate information effectively. The study carried out by Soal (2014) offers models/coefficients that account for acceleration in multiple axis at multiple locations on the vessel; these, however, were not implemented to save time. This function would need to be refactored to include these (by taking location and axis as inputs and by selecting the coefficients based on these).
            */

            if((portPropMotorPower < 0F) || (relativeWindSpeed < 0) || (relativeWindDirection < 0) || (heading < 0) || (waveHeight < 0))
            {
                throw new ArgumentException("The only input that can have a negative value is latitude");
            }

            // Define coefficients for a z-axis estimate on the bridge in open water
            float intercept = 1.7605F;
            float alpha = 0.0010F;
            float beta = -0.0004F;
            float gamma = 0.2050F;
            float delta = -0.0008F;
            float zeta = 0.0017F;
            float eta = 0F;
            
            return (intercept + (alpha*portPropMotorPower) + (beta*latitude) + (gamma*relativeWindSpeed) + (delta*relativeWindDirection) + (zeta*heading) + (eta*waveHeight));
        }

        internal float CalculateIceResponse(float s10Bow, float s15SternShoulder, float portPropMotorPower, float longitude, float relativeWindSpeed, float relativeWindDirection, float gpsSOG, float floeSize)
        {
            /* This function calculates the vibration response of the vessel for ice sailing conditions. This model requires estimates for accelerometer values on the vessel, and until an estimation model for these sensors is developed this call cannnot be used. In this implementation, only the y-axis acceleration in the bridge is required, purely as a means to demonstrate the design's ability to aggregate and coordinate information effectively. The study carried out by Soal (2014) offers models/coefficients that account for acceleration in multiple axis at multiple locations on the vessel; these, however, were not implemented to save time. This function would need to be refactored to include these (by taking location and axis as inputs and by selecting the coefficients based on these)
            */

            // Define coefficients for a z-axis estimate on the bridge in open water
            float intercept = 0.4014F;
            float alpha = 0.0015F;
            float beta = 0.0001F;
            float gamma = 0F;
            float delta = 0F;
            float zeta = -0.0151F;
            float eta = 0.0001F;
            float theta = 0F;
            float phi = 0F;

            return (intercept + (alpha*s10Bow) + (beta*s15SternShoulder) + (gamma*portPropMotorPower) + (delta*longitude) + (zeta*relativeWindSpeed) + (eta*relativeWindDirection) + (theta*gpsSOG) + (phi*floeSize));
        }
    }
}
