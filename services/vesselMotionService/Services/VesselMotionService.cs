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

            return base.MotionEstimate(request, context);
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

    }
}
