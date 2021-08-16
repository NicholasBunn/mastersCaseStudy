# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import ocean_weather_service_api_v1_pb2 as ocean__weather__service__api__v1__pb2


class OceanWeatherServiceStub(object):
    """'Ocean Weather Service' offers two service calls that provide information about ocean weather conditions for use in route planning.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.OceanWeatherPrediction = channel.unary_unary(
                '/oceanWeatherServiceAPI.v1.OceanWeatherService/OceanWeatherPrediction',
                request_serializer=ocean__weather__service__api__v1__pb2.OceanWeatherPredictionRequest.SerializeToString,
                response_deserializer=ocean__weather__service__api__v1__pb2.OceanWeatherInformationResponse.FromString,
                )
        self.OceanWeatherHistory = channel.unary_unary(
                '/oceanWeatherServiceAPI.v1.OceanWeatherService/OceanWeatherHistory',
                request_serializer=ocean__weather__service__api__v1__pb2.OceanWeatherHistoryRequest.SerializeToString,
                response_deserializer=ocean__weather__service__api__v1__pb2.OceanWeatherInformationResponse.FromString,
                )


class OceanWeatherServiceServicer(object):
    """'Ocean Weather Service' offers two service calls that provide information about ocean weather conditions for use in route planning.
    """

    def OceanWeatherPrediction(self, request, context):
        """The 'Ocean Weather Prediction' call provides foresight for tactical decision-making by providing future ocean weather conditions along a requested route
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def OceanWeatherHistory(self, request, context):
        """The 'OceanWeather History' call provides hindsight for stategic decision-making by providing historical ocean weather conditions that the ship would have encountered along a requested route
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_OceanWeatherServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'OceanWeatherPrediction': grpc.unary_unary_rpc_method_handler(
                    servicer.OceanWeatherPrediction,
                    request_deserializer=ocean__weather__service__api__v1__pb2.OceanWeatherPredictionRequest.FromString,
                    response_serializer=ocean__weather__service__api__v1__pb2.OceanWeatherInformationResponse.SerializeToString,
            ),
            'OceanWeatherHistory': grpc.unary_unary_rpc_method_handler(
                    servicer.OceanWeatherHistory,
                    request_deserializer=ocean__weather__service__api__v1__pb2.OceanWeatherHistoryRequest.FromString,
                    response_serializer=ocean__weather__service__api__v1__pb2.OceanWeatherInformationResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'oceanWeatherServiceAPI.v1.OceanWeatherService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class OceanWeatherService(object):
    """'Ocean Weather Service' offers two service calls that provide information about ocean weather conditions for use in route planning.
    """

    @staticmethod
    def OceanWeatherPrediction(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/oceanWeatherServiceAPI.v1.OceanWeatherService/OceanWeatherPrediction',
            ocean__weather__service__api__v1__pb2.OceanWeatherPredictionRequest.SerializeToString,
            ocean__weather__service__api__v1__pb2.OceanWeatherInformationResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def OceanWeatherHistory(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/oceanWeatherServiceAPI.v1.OceanWeatherService/OceanWeatherHistory',
            ocean__weather__service__api__v1__pb2.OceanWeatherHistoryRequest.SerializeToString,
            ocean__weather__service__api__v1__pb2.OceanWeatherInformationResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
