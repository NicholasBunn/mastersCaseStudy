# Standard library imports
import os
import unittest

# Local application imports
import oceanWeatherService

# ToDo: Choose relevant points and time
# ToDo: Use mock tests (from unittest.mock import patch) to run these tests so that they are not reliant on Stormglass's servers or an internet connection.
# ToDo: Add assertEqual for wave_length in test_oceanWeatherHistory once the logic has been added

class OceanWeatherServiceUnitTest(unittest.TestCase):
	''' This class is used to execute all unit tests on the Ocean Weather Service. Put any tests used to verify functions in this class.
	'''

	config = oceanWeatherService.loadConfigFile("configuration.yaml") # Load in the config file. This is needed for the Stormglass API key

	testLat = 58.7984
	testLong = 17.8081
	startTimeUnix = 1608580803.0

	def test_queryStormglassAPI(self):
		''' This tests for a succesful API call. This test should pass as long as the historical ocean weather data and the API's response (JSON structure) remain unchanged
		'''		

		print("Testing Ocean Weather Service: Unit Test: Query Stormglass API (Function Test)")

		jsonOceanData = oceanWeatherService.queryStormglassAPI(self.testLat, self.testLong, self.startTimeUnix, self.config["authentication"]["stormglass"]["apiKey"])

		# This is just here while we're working with a limited number of API queries
		if(jsonOceanData == "HTTP Error 402: You've reached the daily limit of your trial API, Stormglass wants you to pay to increase the daily request limit."):
			self.fail("Can't run queryStormglassAPI test as daily query limit has been reached. :(")

		# self.assertEqual(jsonOceanData["time"], )
		self.assertEqual(jsonOceanData["windDirection"]["icon"], 222.41)
		self.assertEqual(jsonOceanData["windSpeed"]["icon"], 7.68)
		self.assertEqual(jsonOceanData["swellDirection"]["icon"], 168.41)
		self.assertEqual(jsonOceanData["swellHeight"]["icon"], 0.65)
		self.assertEqual(jsonOceanData["swellPeriod"]["icon"], 5.89)

	def test_invalidAPIKey(self):
		''' This tests the error handling of an invalid API key
		'''

		print("Testing Ocean Weather Service: Unit Test: Invalid API Key (Error Handling Test)")

		with self.assertRaises(Exception):
			oceanWeatherService.queryStormglassAPI(self.testLat, self.testLong, self.startTimeUnix, "Wrong API key")

	def test_invalidAPIRequest(self):
		''' This tests the error handling of an invalid API request for the case where the provided latitude, longitude, or unixTime variables are of the incorrect type
		'''

		print("Testing Ocean Weather Service: Unit Test: Invalid API Request (Error Handling Test)")
		
		with self.assertRaises(ValueError):
			oceanWeatherService.queryStormglassAPI("Lets give it a string", "Here's another string, just for fun", self.startTimeUnix, self.config["authentication"]["stormglass"]["apiKey"])
			oceanWeatherService.queryStormglassAPI(self.testLat, self.testLong, "Incorrect data type",  self.config["authentication"]["stormglass"]["apiKey"])

class OceanWeatherServiceIntegrationTest(unittest.TestCase):
	''' This class is used to execute all integration tests on the Ocean Weather Service. Put any tests used to verify the gRPC/server implementation in this class.
	'''
	
	config = oceanWeatherService.loadConfigFile("configuration.yaml")
	serverClass = oceanWeatherService.OceanWeatherServiceServicer
	hostName = os.getenv(key = "FETCHDATAHOST", default = "localhost")
	port = f'{hostName}:{config["port"]["myself"]}'
	
	testLat = 58.7984
	testLong = 17.8081
	startTimeUnix = 1608580803.0

	def setUp(self):
		''' setUp is used to create a server instance to test each service call
		'''

		self.server = oceanWeatherService.grpc.server(oceanWeatherService.futures.ThreadPoolExecutor(max_workers=10))
		oceanWeatherService.config = oceanWeatherService.loadConfigFile("configuration.yaml") 		# Set up the config file on the server side so that the Stormglass API key is accessible
		oceanWeatherService.ocean_weather_service_api_v1_pb2_grpc.add_OceanWeatherServiceServicer_to_server(self.serverClass(), self.server)
		self.server.add_insecure_port(self.port)
		self.server.start()
	
	def tearDown(self):
		''' tearDown is used to pull down the server instance that was used for testing
		'''

		self.server.stop(None)

	def test_OceanWeatherPrediction(self):
		''' This function tests the OceanWeatherPrediction-specific functionality. It's main purpose is to ensure that the request iterates through the requested points as intended.
		'''

		print("Testing Ocean Weather Service: Integration Test: Ocean Weather Prediction (Service Call Test)")

		with oceanWeatherService.grpc.insecure_channel(self.port) as channel:
			stub = oceanWeatherService.ocean_weather_service_api_v1_pb2_grpc.OceanWeatherServiceStub(channel)
			response = stub.OceanWeatherPrediction(oceanWeatherService.ocean_weather_service_api_v1_pb2.OceanWeatherPredictionRequest(
				latitude={self.testLat},
				longitude={self.testLong},
				unix_time = {self.startTimeUnix},
				))

		self.assertEqual(response.wind_direction, [222.41000366210938])
		self.assertEqual(response.wind_speed, [7.679999828338623])
		self.assertEqual(response.beaufort_number, [4])
		self.assertEqual(response.swell_direction, [168.41000366210938])
		self.assertEqual(response.swell_height, [0.6499999761581421])
		self.assertEqual(response.swell_frequency, [0.16977928578853607])
		self.assertEqual(response.swell_period, [5.889999866485596])
		# self.assertEqual(response.wave_length, [])

	def test_OceanWeatherHistory(self):
		''' This function tests the OceanWeatherHistory-specific functionality. It's main purpose is to ensure that the correct archival service is selected.
		'''

		print("Testing Ocean Weather Service: Integration Test: Ocean Weather History (Service Call Test)")

		with oceanWeatherService.grpc.insecure_channel(self.port) as channel:
			stub = oceanWeatherService.ocean_weather_service_api_v1_pb2_grpc.OceanWeatherServiceStub(channel)
			response = stub.OceanWeatherHistory(oceanWeatherService.ocean_weather_service_api_v1_pb2.OceanWeatherHistoryRequest(
				latitude={self.testLat},
				longitude={self.testLong},
				unix_time = {self.startTimeUnix},
				archive_service  = oceanWeatherService.ocean_weather_service_api_v1_pb2.ArchiveService.STORMGLASS,
				))

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
