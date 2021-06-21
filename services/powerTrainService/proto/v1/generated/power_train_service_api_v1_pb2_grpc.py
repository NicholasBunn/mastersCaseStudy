# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import power_train_service_api_v1_pb2 as power__train__service__api__v1__pb2


class PowerTrainServiceStub(object):
    """'Power Train Service; offers four service calls that provide information about the power train of the vessel (namely power requirements and their assosciated costs)
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.PowerEstimate = channel.unary_unary(
                '/powerTrainService.v1.PowerTrainService/PowerEstimate',
                request_serializer=power__train__service__api__v1__pb2.EstimateRequest.SerializeToString,
                response_deserializer=power__train__service__api__v1__pb2.PowerEstimateResponse.FromString,
                )
        self.CostEstimate = channel.unary_unary(
                '/powerTrainService.v1.PowerTrainService/CostEstimate',
                request_serializer=power__train__service__api__v1__pb2.EstimateRequest.SerializeToString,
                response_deserializer=power__train__service__api__v1__pb2.CostEstimateResponse.FromString,
                )
        self.PowerTracking = channel.unary_unary(
                '/powerTrainService.v1.PowerTrainService/PowerTracking',
                request_serializer=power__train__service__api__v1__pb2.TrackingRequest.SerializeToString,
                response_deserializer=power__train__service__api__v1__pb2.TrackingResponse.FromString,
                )
        self.PowerEstimateEvaluation = channel.unary_unary(
                '/powerTrainService.v1.PowerTrainService/PowerEstimateEvaluation',
                request_serializer=power__train__service__api__v1__pb2.EstimateRequest.SerializeToString,
                response_deserializer=power__train__service__api__v1__pb2.PowerEvaluationResponse.FromString,
                )


class PowerTrainServiceServicer(object):
    """'Power Train Service; offers four service calls that provide information about the power train of the vessel (namely power requirements and their assosciated costs)
    """

    def PowerEstimate(self, request, context):
        """The 'Power Estimate' call provides foresight for tactical decision-making by providing power estimates for a requested route and sailing conditions
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CostEstimate(self, request, context):
        """The 'Cost Estimate' call provides foresight for ?? decision-making by providing cost estimates for a requested route and sailign conditions
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def PowerTracking(self, request, context):
        """The 'Power Tracking' call provides insight for ?? decision-making by providing real-time power use by the vessel
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def PowerEstimateEvaluation(self, request, context):
        """The 'Power Estimate Evaluation' call provdes ?? for ?? decision-making by evaluating the accuracy of the models predictions
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_PowerTrainServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'PowerEstimate': grpc.unary_unary_rpc_method_handler(
                    servicer.PowerEstimate,
                    request_deserializer=power__train__service__api__v1__pb2.EstimateRequest.FromString,
                    response_serializer=power__train__service__api__v1__pb2.PowerEstimateResponse.SerializeToString,
            ),
            'CostEstimate': grpc.unary_unary_rpc_method_handler(
                    servicer.CostEstimate,
                    request_deserializer=power__train__service__api__v1__pb2.EstimateRequest.FromString,
                    response_serializer=power__train__service__api__v1__pb2.CostEstimateResponse.SerializeToString,
            ),
            'PowerTracking': grpc.unary_unary_rpc_method_handler(
                    servicer.PowerTracking,
                    request_deserializer=power__train__service__api__v1__pb2.TrackingRequest.FromString,
                    response_serializer=power__train__service__api__v1__pb2.TrackingResponse.SerializeToString,
            ),
            'PowerEstimateEvaluation': grpc.unary_unary_rpc_method_handler(
                    servicer.PowerEstimateEvaluation,
                    request_deserializer=power__train__service__api__v1__pb2.EstimateRequest.FromString,
                    response_serializer=power__train__service__api__v1__pb2.PowerEvaluationResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'powerTrainService.v1.PowerTrainService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class PowerTrainService(object):
    """'Power Train Service; offers four service calls that provide information about the power train of the vessel (namely power requirements and their assosciated costs)
    """

    @staticmethod
    def PowerEstimate(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/powerTrainService.v1.PowerTrainService/PowerEstimate',
            power__train__service__api__v1__pb2.EstimateRequest.SerializeToString,
            power__train__service__api__v1__pb2.PowerEstimateResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CostEstimate(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/powerTrainService.v1.PowerTrainService/CostEstimate',
            power__train__service__api__v1__pb2.EstimateRequest.SerializeToString,
            power__train__service__api__v1__pb2.CostEstimateResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def PowerTracking(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/powerTrainService.v1.PowerTrainService/PowerTracking',
            power__train__service__api__v1__pb2.TrackingRequest.SerializeToString,
            power__train__service__api__v1__pb2.TrackingResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def PowerEstimateEvaluation(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/powerTrainService.v1.PowerTrainService/PowerEstimateEvaluation',
            power__train__service__api__v1__pb2.EstimateRequest.SerializeToString,
            power__train__service__api__v1__pb2.PowerEvaluationResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
