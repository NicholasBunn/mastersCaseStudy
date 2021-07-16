# Standard library imports
import sys
import os
import logging
import math
from concurrent import futures

# Third party imports
import grpc
import yaml

# Local application imports
import proto.v1.generated.comfort_service_api_v1_pb2 as comfort_service_api_v1_pb2
import proto.v1.generated.comfort_service_api_v1_pb2_grpc as comfort_service_api_v1_pb2_grpc

def loadConfigFile(filepath):
	''' This function reads in a YAML configuration file. It takes the relative filepath as an input. It returns a dictionary (?) containing configuration variables.
	'''
	
	with open(os.path.join(sys.path[0], filepath), "r") as f:
		config = yaml.safe_load(f)
		serverConfig = config["server"]

	logging.debug("Succesfully loaded configuration file")
	return serverConfig

def calculateEquivalentVibration(timeStamps, weightedVibration):
	''' This function takes a time series of weighted vibration data as well as their assosciated timestamps. Using these, it calculates the equivalent vibration magnitude according to SANS 2631-1 (Appendix C, equ C.1).
	'''

	numerator = 0.0 # This variables holds the accumulated numerator use in the calculation of the equivalent vibration magnitude

	denominator = 0.0 # This variables holds the accumulated denominator use in the calculation of the equivalent vibration magnitude

	# Iterate through the request samples
	for i in range(1, len(weightedVibration)):
		# Find the time difference between estimates. Assume the acceleration signal is constant over this time period.

		timeGap = timeStamps[i] - timeStamps[i-1]

		# Increment numerator of the equivalent vibration equation.
		numerator += ((weightedVibration[i]**2) * timeGap)

		# Increment denominator of the equivalent vibration equation
		denominator += timeGap

	# Calculate and return the equivalent vibration
	return math.sqrt(numerator/denominator)

def assessComfort(equivalentAcceleration, response):
	''' This function takes an equivalent acceleration and a gRPC response message as inputs. Based on the equivalent acceleration, it sets the rating field of the message according to SANS 2631-1 (Appendix C, section C2.3).
	'''
	
	if(equivalentAcceleration < 0.315):
		response.rating = comfort_service_api_v1_pb2.NOT_UNCOMFORTABLE
	elif(equivalentAcceleration < 0.6):
		response.rating = comfort_service_api_v1_pb2.SLIGHTLY_UNCOMFORTABLE
	elif(equivalentAcceleration < 0.9):
		response.rating = comfort_service_api_v1_pb2.FAIRLY_UNCOMFORTABLE
	elif(equivalentAcceleration < 1.4):
		response.rating = comfort_service_api_v1_pb2.UNCOMFORTABLE
	elif(equivalentAcceleration < 2):
		response.rating = comfort_service_api_v1_pb2.VERY_UNCOMFORTABLE
	else:
		response.rating = comfort_service_api_v1_pb2.EXTREMELY_UNCOMFORTABLE

	return response
		
class ComfortServiceServicer(object):
	"""'Comfort Service' offers one service call that provides information about human comfort onboard, in response to vessel vibrations.
	"""

	def ComfortRating(self, request, context):
		"""The 'Comfort Rating' call provides foresight for tactical decision-making by providing a comfort rating for a proposed route, based on estimated vibrations on board.
		"""

		responseMessage = comfort_service_api_v1_pb2.ComfortResponse()
		responseMessage.unix_time.extend(request.unix_time)
		
		# Calculate equivalent vibration
		equivalentVibration = calculateEquivalentVibration(request.unix_time, request.human_weighted_vibration_z)

		# Assess comfort
		responseMessage = assessComfort(equivalentVibration, responseMessage)

		return responseMessage
	
def serve():
	''' This function creates a server with specified interceptors, registers the service calls offered by that server, and exposes
	the server over a specified port.
	'''

	# Create a server to serve calls in its own thread
	server = grpc.server(
		futures.ThreadPoolExecutor(max_workers=10),
		# interceptors = activeInterceptors
	)
	logging.debug("Successfully created server.")

	# Register an ocean weather service on the server
	comfort_service_api_v1_pb2_grpc.add_ComfortServiceServicer_to_server(ComfortServiceServicer(), server)
	logging.debug("Successfully registered comfort service to server.")

	# Create an insecure connection on port
	comfortHost = os.getenv(key = "COMFORTHOST", default = "localhost") # Receives the hostname from the environmental variables (for Docker network), or defaults to localhost for local testing
	try:
		server.add_insecure_port(f'{comfortHost}:{config["port"]["myself"]}')
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

	logging.basicConfig(filename="services/comfortService/program logs/" + serviceName + ".log", format="%(asctime)s:%(name)s:%(levelname)s:%(module)s:%(funcName)s:%(message)s", level=logging.DEBUG, force = True)

	# ________SERVE REQUESTS________
	serve() # Finish initialisation by serving the request