# Adapted from https://github.com/shaanz/hellonerd/blob/master/client.py

# Standard library imports
import concurrent.futures
import sys
import os

# Third party imports
import grpc

# Local application imports
sys.path.append( os.path.dirname( os.path.dirname( os.path.dirname( os.path.abspath(__file__) ) ) ) )
import protoFiles.python.powerTrainService.v1.power_train_service_api_v1_pb2 as server_pb2
import protoFiles.python.powerTrainService.v1.power_train_service_api_v1_pb2_grpc as server_pb2_grpc

class dosAttacker():

    def __init__(self):
        self.successfulResponses = 0
        self.failedResponses = 0

    
    def sendRequest(self, requestMessage):
        # print("Sending request")
        with grpc.insecure_channel("localhost:50051") as channel:
            # print("Made channel")
            stub = server_pb2_grpc.PowerTrainServiceStub(channel)
            # print("LET'S GO!")
            try:
                response = stub.PowerEstimate(requestMessage)
                self.successfulResponses += 1
            except grpc.RpcError as rpc_error:
                if rpc_error.code() == grpc.StatusCode.RESOURCE_EXHAUSTED:
                    self.failedResponses += 1
                else:
                    print(f"Received an unexpected RPC error: code={rpc_error.code()}")

    def sendMultiple(self, requestMessage, numRequests):
        # Create a list of request messages for the map
        requestMessageList = []
        for _ in range(numRequests):
            requestMessageList.append(requestMessage)

        # Run the calls concurrently
        with concurrent.futures.ThreadPoolExecutor(max_workers=20) as executor:
            executor.map(self.sendRequest, requestMessageList)

if __name__ == '__main__':
    # Create request message
    requestMessage = server_pb2.PowerTrainEstimateRequest(
        unix_time = [1608811845, 1608812145, 1609157745],
				port_prop_motor_speed = [83.5450057983399, 104.089996337891, 120.443740844727],
				stbd_prop_motor_speed = [84.4112548828125, 105.743743896484, 120.522499084473],
				propeller_pitch_port = [-40.5200004577637, 51.3299980163574, 95.3400039672852],
				propeller_pitch_stbd = [-46.2599983215332, 50.5299987792969, 92.3999938964844],
				sog = [0.545311111064, 2.973488888632, 13.756244443256],
				wind_direction_relative = [15.0, 337, 332],
				wind_speed = [2.5, 4.2, 13.6],
				beaufort_number = [0, 0, 3],
				wave_direction = [0.0, 0.0, 255],
				wave_length = [0.0, 0.0, 69.3333333333333],
				model_type = server_pb2.OPENWATER,
    )
    myAttacker = dosAttacker()
    myAttacker.sendMultiple(requestMessage, 100)

    print(f"There were {myAttacker.successfulResponses} succesful calls with {myAttacker.failedResponses} calls being rejected by the rate limit interceptor.")