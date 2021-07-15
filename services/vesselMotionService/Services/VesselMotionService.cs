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
        /*'VesselMotionService' offers three service calls that provide information about the vessel's motion in response to her sailing conditions (mainly regarding the high frequency component of the vessel's acceleration/vibration).
        */
        
        private readonly ILogger<VesselMotionService> _logger;
        public VesselMotionService(ILogger<VesselMotionService> logger)
        {
            _logger = logger;
        }

        struct openWaterCoefficients {
            public float intercept;
            public float alpha;
            public float beta;
            public float gamma;
            public float delta;
            public float zeta;
            public float eta;
        };

        public override Task<MotionEstimateResponse> MotionEstimate(MotionEstimateRequest request, ServerCallContext context)
        {
            /* The 'MotionEstimate' call provides foresight for tactical decision-making (As is described by https://www.researchgate.net/publication/332173693_Designing_Ship_Digital_Services) by providing high-frequency acceleration estimates for a requested sailing conditions at a requested location on the ship (Currently, only the bridge has been implemented). It selects the appropriate model (openwater vs. ice) and iterates through the provided inputs (points) to produce acceleration estimates for each.
            */

            // _logger.LogInformation("Received Motion Estimate Service Call");

            // Create a response object            
            var response = new MotionEstimateResponse();  

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
                            (double xEstimate, double yEstimate, double zEstimate) = CalculateOpenWaterResponse(request.PortPropMotorPower[i], request.Latitude[i], request.WindSpeedRelative[i], request.WindDirectionRelative[i], request.Heading[i], request.WaveHeight[i]);
                            response.AccelerationEstimateX.Add(xEstimate);
                            response.AccelerationEstimateY.Add(yEstimate);
                            response.AccelerationEstimateZ.Add(zEstimate);
                        }

                        // _logger.LogInformation("Succesfully calculated response for open water model");

                        break;

                    case ModelTypeEnum.Ice:
                        throw new RpcException(new Status(StatusCode.Unimplemented, "Ice estimation logic has not been implemented yet"));

                    case ModelTypeEnum.UnknownModel:
                        // Iterate through the provided inputs and produce estimates for each set
                        for(int i = 0; i < request.PortPropMotorPower.Count; i++)
                        {
                            // For the current set of input variables, add the estimate to the response
                            (double xEstimate, double yEstimate, double zEstimate) = CalculateOpenWaterResponse(request.PortPropMotorPower[i], request.Latitude[i], request.WindSpeedRelative[i], request.WindDirectionRelative[i], request.Heading[i], request.WaveHeight[i]);
                            response.AccelerationEstimateX.Add(xEstimate);
                            response.AccelerationEstimateY.Add(yEstimate);
                            response.AccelerationEstimateZ.Add(zEstimate);
                        }

                        // _logger.LogInformation("Succesfully calculated response for unknown model (defaulted to open water)");
                        break;
                    
                    default:   
                        // Iterate through the provided inputs and produce estimates for each set                         
                        for(int i = 0; i < request.PortPropMotorPower.Count; i++)
                        {
                            // For the current set of input variables, add the estimate to the response
                            (double xEstimate, double yEstimate, double zEstimate) = CalculateOpenWaterResponse(request.PortPropMotorPower[i], request.Latitude[i], request.WindSpeedRelative[i], request.WindDirectionRelative[i], request.Heading[i], request.WaveHeight[i]);
                            response.AccelerationEstimateX.Add(xEstimate);
                            response.AccelerationEstimateY.Add(yEstimate);
                            response.AccelerationEstimateZ.Add(zEstimate);
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

        public override Task<MotionTrackingResponse> MotionTracking(MotionTrackingRequest request, ServerCallContext context)
        {
            /* The 'MotionTracking' call provides insight for tactical decision-making (As is described by https://www.researchgate.net/publication/332173693_Designing_Ship_Digital_Services) by providing real-time, high-frequency acceleration readings for a requested location on the ship.
            */

            return base.MotionTracking(request, context);
        }

        public override Task<MotionEvaluationResponse> MotionEstimateEvaluation(MotionEstimateRequest request, ServerCallContext context)
        {
            /* The 'MotionEstimationEvaluation' call provides hindsight for strategic decision-making (As is described by https://www.researchgate.net/publication/332173693_Designing_Ship_Digital_Services) by evaluating the accuracy of the models predictions.
            */

            return base.MotionEstimateEvaluation(request, context);
        }

        internal (double, double, double) CalculateOpenWaterResponse(float portPropMotorPower, float latitude, float relativeWindSpeed, float relativeWindDirection, float heading, float waveHeight)
        {
            /* This function calculates the vibration response (in m/s^2) of the vessel for open water sailing conditions. It takes the port motor power, the vesel's latitude, the relativeWindSpeed, the relativeWindDirection, the heading, and the wave height as inputs. It returns an acceleration estimate for the bridge in the y-axis. 
            In this implementation, only the y-axis acceleration in the bridge is required (this is human-weighted), purely as a means to demonstrate the design's ability to aggregate and coordinate information effectively. The study carried out by Keith Soal (https://scholar.sun.ac.za/handle/10019.1/96058) offers models/coefficients that account for acceleration in multiple axes at multiple locations on the vessel; these, however, were not implemented to save time. This function would need to be refactored to include these (by taking location and axis as inputs and by selecting the coefficients based on these).
            */

            if((portPropMotorPower < 0F) || (relativeWindSpeed < 0) || (relativeWindDirection < 0) || (heading < 0) || (waveHeight < 0))
            {
                throw new ArgumentException("The only input that can have a negative value is latitude");
            }
            
            // Define coefficients for a x-axis estimate on the bridge in open water. These provide the human-weighted vibrations
            openWaterCoefficients xCoefficients;
            xCoefficients.intercept = 2.7298F;
            xCoefficients.alpha = 0.0013F;
            xCoefficients.beta = -0.0006F;
            xCoefficients.gamma = 0.3430F;
            xCoefficients.delta = -0.0007F;
            xCoefficients.zeta = 0.0037F;
            xCoefficients.eta = 0F;

            // Define coefficients for a y-axis estimate on the bridge in open water. These provide the human-weighted vibrations
            openWaterCoefficients yCoefficients;
            yCoefficients.intercept = 2.5711F;
            yCoefficients.alpha = 0.0016F;
            yCoefficients.beta = -0.0004F;
            yCoefficients.gamma = 0F;
            yCoefficients.delta = -0.0010F;
            yCoefficients.zeta = 0.0011F;
            yCoefficients.eta = 0.8668F;
            
            // Define coefficients for a z-axis estimate on the bridge in open water. These provide the human-weighted vibrations
            openWaterCoefficients zCoefficients;
            zCoefficients.intercept = 1.7605F;
            zCoefficients.alpha = 0.0010F;
            zCoefficients.beta = -0.0004F;
            zCoefficients.gamma = 0.2050F;
            zCoefficients.delta = -0.0008F;
            zCoefficients.zeta = 0.0017F;
            zCoefficients.eta = 0F;

            float xResponse = xCoefficients.intercept + (xCoefficients.alpha*portPropMotorPower) + (xCoefficients.beta*latitude) + (xCoefficients.gamma*relativeWindSpeed) + (xCoefficients.delta*relativeWindDirection) + (xCoefficients.zeta*heading) + (xCoefficients.eta*waveHeight);

            float yResponse = yCoefficients.intercept + (yCoefficients.alpha*portPropMotorPower) + (yCoefficients.beta*latitude) + (yCoefficients.gamma*relativeWindSpeed) + (yCoefficients.delta*relativeWindDirection) + (yCoefficients.zeta*heading) + (yCoefficients.eta*waveHeight);

            float zResponse = zCoefficients.intercept + (zCoefficients.alpha*portPropMotorPower) + (zCoefficients.beta*latitude) + (zCoefficients.gamma*relativeWindSpeed) + (zCoefficients.delta*relativeWindDirection) + (zCoefficients.zeta*heading) + (zCoefficients.eta*waveHeight);

            return(xResponse/1000, yResponse/1000, zResponse/1000);
        }

        internal float CalculateIceResponse(float s10Bow, float s15SternShoulder, float portPropMotorPower, float longitude, float relativeWindSpeed, float relativeWindDirection, float gpsSOG, float floeSize)
        {
            /* This function calculates the vibration response of the vessel for ice sailing conditions. It takes a sensor estimate for the S10Bow, and S15SternShoulder (as the S.A. Agulhas is instrumented), the port motor power, the vessels longitude, the relative wind speed, the relative wind direction, the GPS speed over ground, and the floe size as inputs. It returns an acceleration estimate for the bridge in the y-axis.
            This model requires estimates for accelerometer values on the vessel, and until an estimation model for these sensors is developed this call cannnot be used. In this implementation, only the y-axis acceleration in the bridge is required, purely as a means to demonstrate the design's ability to aggregate and coordinate information effectively. The study carried out by Soal (2014) offers models/coefficients that account for acceleration in multiple axis at multiple locations on the vessel; these, however, were not implemented to save time. This function would need to be refactored to include these (by taking location and axis as inputs and by selecting the coefficients based on these)
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
