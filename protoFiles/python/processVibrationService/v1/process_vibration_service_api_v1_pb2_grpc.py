# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import process_vibration_service_api_v1_pb2 as process__vibration__service__api__v1__pb2


class ProcessVibrationServiceStub(object):
    """'Process Vibration Service' offers two service calls that process time-series vibration signals according to accepted practices.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CalculateRMSSeries = channel.unary_unary(
                '/processVibrationServiceAPI.v1.ProcessVibrationService/CalculateRMSSeries',
                request_serializer=process__vibration__service__api__v1__pb2.ProcessRequest.SerializeToString,
                response_deserializer=process__vibration__service__api__v1__pb2.ProcessResponseSeries.FromString,
                )
        self.CalculateRMSBatch = channel.unary_unary(
                '/processVibrationServiceAPI.v1.ProcessVibrationService/CalculateRMSBatch',
                request_serializer=process__vibration__service__api__v1__pb2.ProcessRequest.SerializeToString,
                response_deserializer=process__vibration__service__api__v1__pb2.ProcessResponseBatch.FromString,
                )


class ProcessVibrationServiceServicer(object):
    """'Process Vibration Service' offers two service calls that process time-series vibration signals according to accepted practices.
    """

    def CalculateRMSSeries(self, request, context):
        """The 'Calculate RMS Series' call calculates the root mean square (RMS) vibration for individual time-series vibration signals.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CalculateRMSBatch(self, request, context):
        """The 'Calculate RMS Batch' call calculates the root mean square (RMS) vibration for an "ensemble" value of a vibration signal time-series.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ProcessVibrationServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CalculateRMSSeries': grpc.unary_unary_rpc_method_handler(
                    servicer.CalculateRMSSeries,
                    request_deserializer=process__vibration__service__api__v1__pb2.ProcessRequest.FromString,
                    response_serializer=process__vibration__service__api__v1__pb2.ProcessResponseSeries.SerializeToString,
            ),
            'CalculateRMSBatch': grpc.unary_unary_rpc_method_handler(
                    servicer.CalculateRMSBatch,
                    request_deserializer=process__vibration__service__api__v1__pb2.ProcessRequest.FromString,
                    response_serializer=process__vibration__service__api__v1__pb2.ProcessResponseBatch.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'processVibrationServiceAPI.v1.ProcessVibrationService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ProcessVibrationService(object):
    """'Process Vibration Service' offers two service calls that process time-series vibration signals according to accepted practices.
    """

    @staticmethod
    def CalculateRMSSeries(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/processVibrationServiceAPI.v1.ProcessVibrationService/CalculateRMSSeries',
            process__vibration__service__api__v1__pb2.ProcessRequest.SerializeToString,
            process__vibration__service__api__v1__pb2.ProcessResponseSeries.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CalculateRMSBatch(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/processVibrationServiceAPI.v1.ProcessVibrationService/CalculateRMSBatch',
            process__vibration__service__api__v1__pb2.ProcessRequest.SerializeToString,
            process__vibration__service__api__v1__pb2.ProcessResponseBatch.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)