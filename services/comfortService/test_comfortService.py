# Standard library imports
import os
import sys
import time
import unittest

# Third party imports
import yaml

# Local application imports
import comfortService

class ComfortServiceUnitTest(unittest.TestCase):
	''' This class is used to execute all unit tests on the Comfort Service. Put any tests used to verify functions in this class.
	'''

	config = comfortService.loadConfigFile("configuration.yaml") # Load in the config file. This is needed for the Stormglass API key

	def test_calculateEquivalentVibrationSingleAxis(self):
		''' This tests that the calculate equivalent vibration function correctly iterates through and sums inputs for a single axis.
		'''

		print("Testing Comfort Service: Unit Test: Calculate Equivalent Vibration (Function Test)")

		response = comfortService.calculateEquivalentVibrationSingleAxis([1608806700, 1608810300, 1608813900], [0.0070642624, 0.003242324, 0.004021232])

		self.assertEqual(response, 0.003652599876717952)

	def test_calculateEquivalentVibrationMultiAxis(self):
		''' This tests that the calculate equivalent vibration function correctly iterates through and sums inputs for multiple axis inputs.
		'''

		print("Testing Comfort Service: Unit Test: Calculate Equivalent Vibration (Function Test)")

		response = comfortService.calculateEquivalentVibrationMultiAxis([1608806700, 1608810300, 1608813900], [0.0070642624, 0.003242324, 0.004021232], [0.0070642624, 0.003242324, 0.004021232], [0.0070642624, 0.003242324, 0.004021232])

		self.assertEqual(response, 0.00632648856619531)

	def test_assessComfort(self):
		''' This tests that the assess comfort function succesfully classifies vibration comfort according to SANS 2631-1.
		'''

		print("Testing Comfort Service: Unit Test: Assess Comfort (Function Test)")

		response = comfortService.comfort_service_api_v1_pb2.ComfortResponse()

		response.unix_time.extend([1608806700, 1608810300, 1608813900]) 

		self.assertEqual(comfortService.assessComfort(0.2, response).unix_time, [1608806700, 1608810300, 1608813900])
		self.assertEqual(comfortService.assessComfort(0.2, response).rating, comfortService.comfort_service_api_v1_pb2.NOT_UNCOMFORTABLE)
		self.assertEqual(comfortService.assessComfort(0.315, response).rating, comfortService.comfort_service_api_v1_pb2.SLIGHTLY_UNCOMFORTABLE)
		self.assertEqual(comfortService.assessComfort(0.6, response).rating, comfortService.comfort_service_api_v1_pb2.FAIRLY_UNCOMFORTABLE)
		self.assertEqual(comfortService.assessComfort(0.9, response).rating, comfortService.comfort_service_api_v1_pb2.UNCOMFORTABLE)
		self.assertEqual(comfortService.assessComfort(1.4, response).rating, comfortService.comfort_service_api_v1_pb2.VERY_UNCOMFORTABLE)
		self.assertEqual(comfortService.assessComfort(2, response).rating, comfortService.comfort_service_api_v1_pb2.EXTREMELY_UNCOMFORTABLE)

class ComfortServiceIntegrationTest(unittest.TestCase):
	''' This class is used to execute all integration tests on the Comfort Service. Put any tests used to verify the gRPC/server implementation in this class.
	'''
	
	config = comfortService.loadConfigFile("configuration.yaml")
	serverClass = comfortService.ComfortServiceServicer
	hostName = os.getenv(key = "COMFORTHOST", default = "localhost")
	port = f'{hostName}:{config["port"]["myself"]}'

	def setUp(self):
		''' setUp is used to create a server instance to test each service call
		'''

		self.server = comfortService.grpc.server(comfortService.futures.ThreadPoolExecutor(max_workers=10))
		comfortService.config = comfortService.loadConfigFile("configuration.yaml") 		# Set up the config file on the server side so that the Stormglass API key is accessible
		comfortService.comfort_service_api_v1_pb2_grpc.add_ComfortServiceServicer_to_server(self.serverClass(), self.server)
		self.server.add_insecure_port(self.port)
		self.server.start()
	
	def tearDown(self):
		''' tearDown is used to pull down the server instance that was used for testing
		'''

		self.server.stop(None)

	def test_ComfortRating(self):
		''' This function tests the ComfortRating-specific functionality. It's main purpose is to ensure that the correct functions are called.
		'''

		print("Testing Comfort Rating Service: Integration Test: Comfort Rating (Service Call Test)")

		with comfortService.grpc.insecure_channel(self.port) as channel:
			stub = comfortService.comfort_service_api_v1_pb2_grpc.ComfortServiceStub(channel)
			response = stub.ComfortRating(comfortService.comfort_service_api_v1_pb2.ComfortRequest(
				unix_time = {1608806700, 1608810300, 1608813900},
				human_weighted_vibration_x = {0.0023422532, 0.001232312, 0.002415324},
				human_weighted_vibration_y = {0.0032141522, 0.002412421, 0.003421513},
				human_weighted_vibration_z = {0.0070642624, 0.003242324, 0.004021232},
			))

			self.assertEqual(response.unix_time, [1608806700, 1608810300, 1608813900])
			self.assertEqual(response.rating, comfortService.comfort_service_api_v1_pb2.NOT_UNCOMFORTABLE)

if __name__ == '__main__':
	unittest.main()
