# Standard library imports
import sys
import os
import logging
import requests
from concurrent import futures

# Third party imports
import grpc
from requests.sessions import TooManyRedirects
import yaml
from urllib.error import HTTPError

# Local application imports
import proto.v1.generated.ocean_weather_service_api_v1_pb2 as ocean_weather_service_api_v1_pb2
import proto.v1.generated.ocean_weather_service_api_v1_pb2_grpc as ocean_weather_service_api_v1_pb2_grpc

# ToDo: Add wave_length to response message for OceanWeatherHistory

def loadConfigFile(filepath):
	with open(os.path.join(sys.path[0], filepath), "r") as f:
		config = yaml.safe_load(f)
		serverConfig = config["server"]
	return serverConfig

def queryWaveAPI(latitude, longitude, unixTime, apiKey):
    	
	# Check if the provided unixTime is a float (the request can't be built if it isn't as the end parameter is a concatenation of this input)
	if (type(unixTime) != float):
		raise ValueError("Requested time is not of the correct type (should be float).")

	try:
		response = requests.get(
			'https://api.stormglass.io/v2/weather/point',
			params={
				'lat': latitude, 
				'lng': longitude, 
				'params': 'windSpeed',
				'params': ','.join(['windSpeed', 'windDirection', 'swellDirection', 'swellHeight', 'swellPeriod']),
				'start': unixTime,
				'end': unixTime + 3600,
			},
			headers={
				'Authorization': apiKey
			}
		)
		response.raise_for_status()
	except requests.HTTPError as exception:
		if (exception.response.status_code == 400):
			raise Exception(f"HTTP Error {exception.response.status_code}: Invalid request.")
		elif (exception.response.status_code == 401):
			raise Exception(f"HTTP Error {exception.response.status_code}: Missing authentication credentials (Most likely that API key has not been provided).")
		elif (exception.response.status_code == 402):
			raise Exception(f"HTTP Error {exception.response.status_code}: You've reached the daily limit of your trial API, Stormglass wants you to pay to increase the daily request limit.")
		elif (exception.response.status_code == 403):
			raise Exception( f"HTTP Error {exception.response.status_code}: Provided API key does not authenticate the request.")
		elif (exception.response.status_code == 422):
			raise ValueError(f"HTTP Error {exception.response.status_code}: Server can't process request (Most likely because of incorrect lat/long type.)")
		elif (exception.response.status_code == 429):
			raise Exception(f"HTTP Error {exception.response.status_code}: You have exhausted your daily API query limit.")
		elif (exception.response.status_code == 500):
			raise Exception(f"HTTP Error {exception.response.status_code}: There was an error on Stormglass's side, not much we can until they fix it.")
		elif (exception.response.status_code == 503):
			raise Exception(f"HTTP Error {exception.response.status_code}: Stormglass's servers are down for maintenance, we'll have to wait this one out.")
		else:
			raise Exception(exception)
	except requests.ConnectionError:
		raise ConnectionError("Failed to connect to Stormglass.")
	except requests.Timeout:
		return TimeoutError("Request timed out")
	except Exception as exception:
		return exception

	return response.json()["hours"][0]

class OceanWeatherServiceServicer(ocean_weather_service_api_v1_pb2_grpc.OceanWeatherServiceServicer):
	"""'Ocean Weather Service' offers two service calls that provide information about ocean weather conditions for use in route planning.
	"""

	def OceanWeatherPrediction(self, request, context):
		"""The 'Ocean Weather Prediction' call provides foresight for tactical decision-making by providing future ocean weather conditions along a requested route
		"""

		logging.info("Received Ocean Weather Prediction service call.")
		context.set_code(grpc.StatusCode.UNIMPLEMENTED)
		context.set_details('Method not implemented!')
		raise NotImplementedError('Method not implemented!')

	def OceanWeatherHistory(self, request, context):
		"""The 'Ocean Weather History' call provides hindsight for stategic decision-making by providing historical ocean weather conditions that the ship would have encountered along a requested route
		"""
		
		logging.info("Received Ocean Weather History service call.")
		
		# Create response message
		responseMessage = ocean_weather_service_api_v1_pb2.OceanWeatherInformationResponse()

		# Iterate through all the requested points, fetching the weather data for each. This approach uses one query per point of interest (not particularly efficient)
		for testLat, testLong, startTimeUnix in zip(request.latitude, request.longitude, request.timestamp):
			try:
				# Query Stormglass API
				jsonOceanData = queryWaveAPI(testLat, testLong, startTimeUnix, config["authentication"]["stormglass"]["apiKey"])

				# Populate response message  
				responseMessage.wind_direction.append(jsonOceanData["windDirection"]["icon"])
				responseMessage.wind_speed.append(jsonOceanData["windSpeed"]["icon"])
				responseMessage.swell_direction.append(jsonOceanData["swellDirection"]["icon"])
				responseMessage.swell_height.append(jsonOceanData["swellHeight"]["icon"])
				responseMessage.swell_frequency.append(1/jsonOceanData["swellPeriod"]["icon"])
				responseMessage.swell_period.append(jsonOceanData["swellPeriod"]["icon"])

				# Set the beaufort number based on the wind speed
				if(jsonOceanData["windSpeed"]["icon"] < 0.5):
					responseMessage.beaufort_number.append(0)
				elif(jsonOceanData["windSpeed"]["icon"] < 1.5):
					responseMessage.beaufort_number.append(1)
				elif(jsonOceanData["windSpeed"]["icon"] < 3.3):
					responseMessage.beaufort_number.append(2)
				elif(jsonOceanData["windSpeed"]["icon"] < 5.5):
					responseMessage.beaufort_number.append(3)
				elif(jsonOceanData["windSpeed"]["icon"] < 7.9):
					responseMessage.beaufort_number.append(4)
				elif(jsonOceanData["windSpeed"]["icon"] < 10.7):
					responseMessage.beaufort_number.append(5)
				elif(jsonOceanData["windSpeed"]["icon"] < 13.8):
					responseMessage.beaufort_number.append(6)
				elif(jsonOceanData["windSpeed"]["icon"] < 17.1):
					responseMessage.beaufort_number.append(7)
				elif(jsonOceanData["windSpeed"]["icon"] < 20.7):
					responseMessage.beaufort_number.append(8)
				elif(jsonOceanData["windSpeed"]["icon"] < 24.4):
					responseMessage.beaufort_number.append(9)
				elif(jsonOceanData["windSpeed"]["icon"] < 28.4):
					responseMessage.beaufort_number.append(10)
				elif(jsonOceanData["windSpeed"]["icon"] < 32.6):
					responseMessage.beaufort_number.append(11)
				else:
					responseMessage.beaufort_number.append(12)
			except Exception as e:
				logging.debug(f"Failed to query Stormglass API: \n{e}")
				context.set_code(grpc.StatusCode.INTERNAL)
				# context.set_details("bla bla")
				raise e

		return responseMessage

def serve():
	''' This function creates a server with specified interceptors, registers the service calls offered by that server, and exposes
	the server over a specified port. The connection to this port is secured with server-side TLS encryption. 
	'''

	# Create a server to serve calls in its own thread
	server = grpc.server(
		futures.ThreadPoolExecutor(max_workers=10),
		# interceptors = activeInterceptors
	)
	logging.debug("Successfully created server.")

	# Register an ocean weather service on the server
	ocean_weather_service_api_v1_pb2_grpc.add_OceanWeatherServiceServicer_to_server(OceanWeatherServiceServicer(), server)
	logging.debug("Successfully registered ocean weather service to server.")

	# Create an insecure connection on port
	oceanWeatherHost = os.getenv(key = "OCEANWEATHERHOST", default = "localhost") # Receives the hostname from the environmental variables (for Docker network), or defaults to localhost for local testing
	try:
		server.add_insecure_port(f'{oceanWeatherHost}:{config["port"]["myself"]}')
		logging.debug("Succesfully added (insecure) port to server.")
	except Exception as e:
		logging.debug(f"Failed to add (insecure) port to server: \n{e}")

	try:
		# Start server and listen for calls on the specified port
		server.start()
		logging.info(f'Server started on port {config["port"]["myself"]}')

		# Defer termination for a 'persistent' service
		server.wait_for_termination()
	except Exception as e:
		logging.debug(f'Failed to start server on port {config["port"]["myself:"]}: \n{e}')
		
if __name__ == '__main__':
	# ________LOAD CONFIG FILE________
	config = loadConfigFile("configuration.yaml")

	# ________LOGGER SETUP________
	serviceName = __file__.rsplit("/")[-2].rsplit(".")[0]

	logging.basicConfig(filename="services/oceanWeatherService/program logs/" + serviceName + ".log", format="%(asctime)s:%(name)s:%(levelname)s:%(module)s:%(funcName)s:%(message)s", level=logging.DEBUG)

	# ________SERVE REQUEST________
	serve() # Finish initialisation by serving the request