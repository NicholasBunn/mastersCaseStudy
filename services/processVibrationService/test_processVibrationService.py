# Standard library imports
import os
import sys
import time
import logging
import unittest

# Third party imports
import yaml

# Local application imports
import processVibrationService

class PowerTrainIntegrationTest(unittest.TestCase):
	''' This class is used to execute all integration tests on the Ocean Weather Service. Put any tests used to verify the gRPC/server implementation in this class.
	'''

	config = processVibrationService.loadConfigFile("configuration.yaml")
	serverClass = processVibrationService.ProcessVibrationServiceServicer
	hostName = os.getenv(key = "PROCESSVIBRATIONHOST", default="localhost")
	port = f'{hostName}:{config["port"]["myself"]}'

	def setUp(self):
		''' setUp is used to create a server instance to test each service call
		'''

		self.server = processVibrationService.grpc.server(processVibrationService.futures.ThreadPoolExecutor(max_workers=10))
		processVibrationService.process_vibration_service_api_v1_pb2_grpc.add_ProcessVibrationServiceServicer_to_server(self.serverClass(), self.server)
		self.server.add_insecure_port(self.port)
		self.server.start()
	
	def tearDown(self):
		''' tearDown is used to pull down the server instance that was used for testing
		'''

		self.server.stop(None)

	def test_CalculateRMSSeries(self):
		''' This function tests the CalculateRMSSeries-specific functionality. It ensures that the service call takes the correct inputs, processes them as is required for the model, and makes the correct estimates based on the provided inputs.
		'''

		print("Testing Process Vibration Service: Integration Test: Calculate RMS Series (Service Call Test)")

		with processVibrationService.grpc.insecure_channel(self.port) as channel:
			stub = processVibrationService.process_vibration_service_api_v1_pb2_grpc.ProcessVibrationServiceStub(channel)
			response = stub.CalculateRMSSeries(processVibrationService.process_vibration_service_api_v1_pb2.ProcessRequest(
				unix_time = [1626325118, 1608812145, 1626332318],
				vibration_x = [0.011090744, 0.03320345, 0.01342132],
				vibration_y = [0.001046832, 0.00335624, 0.00233526],
				vibration_z = [0.035672354, 0.04567234, 0.00345332],
			))

		# Test that the time is returned, unchanged
		self.assertEqual(response.unix_time, [1626325118, 1608812145, 1626332318])

		# Test the responses
		self.assertEqual(response.rms_vibration_x, [0.007842340506613255, 0.02347838319838047, 0.009490306489169598])
		self.assertEqual(response.rms_vibration_y, [0.0007402219926007092, 0.0023732201661914587, 0.0016512781148776412])
		self.assertEqual(response.rms_vibration_z, [0.025224164128303528, 0.03229521960020065, 0.0024418658576905727])

	def test_CalculateRMSBAtch(self):
		''' This function tests the CalculateRMSBatch-specific functionality. It ensures that the service call takes the correct inputs, processes them as is required for the model, and makes the correct estimates based on the provided inputs.
		'''

		print("Testing Process Vibration Service: Integration Test: Calculate RMS Batch (Service Call Test)")

if __name__ == '__main__':
	unittest.main()
