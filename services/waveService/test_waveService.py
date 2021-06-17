# Standard library imports
import os
import sys
import time
import unittest
from datetime import datetime

# Third party imports
import yaml

# Local application imports
import waveService

# ToDo: Use mock tests (from unittest.mock import patch) to run these tests so that they are not reliant on Stormglass's servers or an internet connection.
def loadConfigFile(filepath):
	with open(os.path.join(sys.path[0], filepath), "r") as f:
		config = yaml.safe_load(f)
	serverConfig = config["server"]
	return serverConfig

class TestWaveService(unittest.TestCase):
	config = loadConfigFile("configuration.yaml")

	def setUp(self):
		''' This function sets up variables for the tests to use. It runs at the beginning of each test function
		'''

		print("Wave Service Test Setup")
		# TODO: Choose relevant points and time
		self.testLat = 58.7984
		self.testLong = 17.8081
		self.startTimeUnix = time.mktime(datetime.strptime("22:00:03 21 December 2020", "%H:%M:%S %d %B %Y").timetuple())

	def tearDown(self):
		''' This function tears down variables that the tests use. It runs at the end of each test function
		'''

		print("Wave Service Test Teardown")
		pass

	def test_queryWaveAPI(self):
		''' This tests for a succesful API call. This test should pass as long as the historical wave data and the API's response (JSON structure) remain unchanged.
		'''		

		print("Wave Service Query Wave API Test")

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

		print("Wave Service Invalid API Key Test")

		with self.assertRaises(Exception):
			waveService.queryWaveAPI(self.testLat, self.testLong, self.startTimeUnix, "Wrong API key")

	def test_invalidAPIRequest(self):
		''' This tests the error handling of an invalid API request for the case where the provided latitude, longitude, or unixTime variables are of the incorrect type.
		'''

		print("Wave Service Invalid API Request Test")
		
		with self.assertRaises(ValueError):
			waveService.queryWaveAPI("Lets give it a string", "Here's another string for fun", self.startTimeUnix, self.config["authentication"]["stormglass"]["apiKey"])
			waveService.queryWaveAPI(self.testLat, self.testLong, "Incorrect data type",  self.config["authentication"]["stormglass"]["apiKey"])

if __name__ == '__main__':
	unittest.main()
