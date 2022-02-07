# Adapted from https://github.com/shaanz/hellonerd/blob/master/client.py

# Standard library imports
import concurrent.futures
import sys
import os

# Third party imports
import grpc

# Local application imports
sys.path.append( os.path.dirname( os.path.dirname( os.path.dirname( os.path.abspath(__file__) ) ) ) )
import protoFiles.python.comfortService.v1.comfort_service_api_v1_pb2 as server_pb2
import protoFiles.python.comfortService.v1.comfort_service_api_v1_pb2_grpc as server_pb2_grpc
import interceptors.python.retryInterceptor as retryInterceptor


if __name__ == '__main__':
    # Create request message
    requestMessage = server_pb2.ComfortRequest(
		unix_time = {1608806700, 1608810300, 1608813900},
		human_weighted_vibration_x = {0.0023422532, 0.001232312, 0.002415324},
		human_weighted_vibration_y = {0.0032141522, 0.002412421, 0.003421513},
		human_weighted_vibration_z = {0.0070642624, 0.003242324, 0.004021232},
	)
    
    interceptor = [retryInterceptor.RetryInterceptor(100, 5)]
    with grpc.insecure_channel("192.168.8.102:50053") as channel:
            channel = grpc.intercept_channel(channel, *interceptor)
            stub = server_pb2_grpc.ComfortServiceStub(channel)
            # print("LET'S GO!")
            try:
                response = stub.ComfortRating(requestMessage)
                print(response)
            except grpc.RpcError as rpc_error:
                print(f"Received an unexpected RPC error: code={rpc_error.code()}")

