# Standard library imports
import os
import sys
import time
import unittest
from datetime import datetime

# Third party imports
import yaml
from google.protobuf.timestamp_pb2 import Timestamp

# Local application imports
import waveService

# ToDo: Use mock tests (from unittest.mock import patch) to run these tests so that they are not reliant on Stormglass's servers or an internet connection.
# ToDo: Add assertEqual for wave_length in test_WaveHistory once the logic has been added

def loadConfigFile(filepath):
	with open(os.path.join(sys.path[0], filepath), "r") as f:
		config = yaml.safe_load(f)
	serverConfig = config["server"]
	return serverConfig

class WaveServiceUnitTest(unittest.TestCase):
	config = loadConfigFile("configuration.yaml")
	# TODO: Choose relevant points and time
	testLat = 58.7984
	testLong = 17.8081
	startTimeUnix = time.mktime(datetime.strptime("22:00:03 21 December 2020", "%H:%M:%S %d %B %Y").timetuple())

	def setUp(self):
		pass

	def tearDown(self):
		pass

	def test_queryWaveAPI(self):
		''' This tests for a succesful API call. This test should pass as long as the historical wave data and the API's response (JSON structure) remain unchanged.
		'''		

		print("Testing Wave Service: Query Wave API (Function Test)")

		jsonWaveData = waveService.queryWaveAPI(self.testLat, self.testLong, self.startTimeUnix, self.config["authentication"]["stormglass"]["apiKey"])

		# This is just here while we're working with a limited number of API queries
		if(jsonWaveData == "HTTP Error 402: You've reached the daily limit of your trial API, Stormglass wants you to pay to increase the daily request limit."):
			self.fail("Can't run queryWaveAPI test as daily query limit has been reached.")

		self.assertEqual(jsonWaveData["windDirection"]["icon"], 222.41)
		self.assertEqual(jsonWaveData["windSpeed"]["icon"], 7.68)
		self.assertEqual(jsonWaveData["swellDirection"]["icon"], 168.41)
		self.assertEqual(jsonWaveData["swellHeight"]["icon"], 0.65)
		self.assertEqual(jsonWaveData["swellPeriod"]["icon"], 5.89)

	def test_invalidAPIKey(self):
		''' This tests the error handling of an invalid API key.
		'''

		print("Testing Wave Service: Invalid API Key (Error Handling Test)")

		with self.assertRaises(Exception):
			waveService.queryWaveAPI(self.testLat, self.testLong, self.startTimeUnix, "Wrong API key")

	def test_invalidAPIRequest(self):
		''' This tests the error handling of an invalid API request for the case where the provided latitude, longitude, or unixTime variables are of the incorrect type.
		'''

		print("Testing Wave Service: Invalid API Request (Error Handling Test)")
		
		with self.assertRaises(ValueError):
			waveService.queryWaveAPI("Lets give it a string", "Here's another string, just for fun", self.startTimeUnix, self.config["authentication"]["stormglass"]["apiKey"])
			waveService.queryWaveAPI(self.testLat, self.testLong, "Incorrect data type",  self.config["authentication"]["stormglass"]["apiKey"])

class WaveServiceIntegrationTest(unittest.TestCase):
	config = loadConfigFile("configuration.yaml")
	serverClass = waveService.WaveServiceServicer
	hostName = os.getenv(key = "FETCHDATAHOST", default = "localhost")
	port = f'{hostName}:{config["port"]["myself"]}'
	
	# TODO: Choose relevant points and time
	testLat = 58.7984
	testLong = 17.8081
	startTimeUnix = time.mktime(datetime.strptime("22:00:03 21 December 2020", "%H:%M:%S %d %B %Y").timetuple())

	def setUp(self):
		''' setUp is used to create a server instance to test each service call
		'''

		self.server = waveService.grpc.server(waveService.futures.ThreadPoolExecutor(max_workers=10))
		waveService.config = waveService.loadConfigFile("configuration.yaml") 		# Set up the config file on the server side so that the Stormglass API key is accessible
		waveService.wave_service_api_v1_pb2_grpc.add_WaveServiceServicer_to_server(self.serverClass(), self.server)
		self.server.add_insecure_port(self.port)
		self.server.start()
	
	def tearDown(self):
		''' tearDown is used to pull down the server instance that was used for testing
		'''

		self.server.stop(None)

	def test_WaveHistory(self):
		''' This function tests the WaveHistory-specific functionality. It's main purpose is to ensure that the request iterates through the requested points as intended
		'''

		print("Testing Wave Service: Wave History (Service Call Test)")

		with waveService.grpc.insecure_channel(self.port) as channel:
			stub = waveService.wave_service_api_v1_pb2_grpc.WaveServiceStub(channel)
			response = stub.WaveHistory(waveService.wave_service_api_v1_pb2.WaveInformationRequest(latitude={self.testLat}, longitude={self.testLong}, timestamp = {self.startTimeUnix}))

		self.assertEqual(response.wind_direction, [222.41000366210938])
		self.assertEqual(response.wind_speed, [7.679999828338623])
		self.assertEqual(response.beaufort_number, [4])
		self.assertEqual(response.swell_direction, [168.41000366210938])
		self.assertEqual(response.swell_height, [0.6499999761581421])
		self.assertEqual(response.swell_frequency, [0.16977928578853607])
		self.assertEqual(response.swell_period, [5.889999866485596])
		# self.assertEqual(response.wave_length, [])

if __name__ == '__main__':
	unittest.main()
