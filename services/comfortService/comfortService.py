# Standard library imports
import sys
import os
import logging
import math
import time
from concurrent import futures

# Third party imports
import grpc
import yaml

# Local application imports
sys.path.append( os.path.dirname( os.path.dirname( os.path.dirname( os.path.abspath(__file__) ) ) ) )
import protoFiles.python.comfortService.v1.comfort_service_api_v1_pb2 as comfort_service_api_v1_pb2
import protoFiles.python.comfortService.v1.comfort_service_api_v1_pb2_grpc as comfort_service_api_v1_pb2_grpc
import interceptors.python.metricInterceptor as metricInterceptor
import interceptors.python.authenticationInterceptor as authenticationInterceptor
import interceptors.python.rateLimitInterceptor as rateLimitInterceptor

def loadConfigFile(filepath):
	''' This function reads in a YAML configuration file. It takes the relative filepath as an input. It returns a dictionary (?) containing configuration variables.
	'''
	
	with open(os.path.join(sys.path[0], filepath), "r") as f:
		config = yaml.safe_load(f)
		serverConfig = config["server"]

	logging.debug("Succesfully loaded configuration file")
	return serverConfig

def calculateEquivalentVibrationMultiAxis(timeStamps, weightedVibrationX, weightedVibrationY, weightedVibrationZ):
	''' This function takes a time series of weighted vibration data for 3 axes as well as their assosciated timestamps. Using these, it calculates the equivalent vibration magnitude according to SANS 2631-1 (Appendix C, equ C.1).
	'''

	numerator = 0.0 # This variables holds the accumulated numerator use in the calculation of the equivalent vibration magnitude

	denominator = 0.0 # This variables holds the accumulated denominator use in the calculation of the equivalent vibration magnitude

	# Iterate through the request samples
	for i in range(1, len(timeStamps)):
		# Find the time difference between estimates. Assume the acceleration signal is constant over this time period.

		timeGap = timeStamps[i] - timeStamps[i-1]

		# Increment numerator of the equivalent vibration equation.

		numerator += (((weightedVibrationX[i]**2) + (weightedVibrationY[i]**2) + (weightedVibrationZ[i]**2)) * timeGap)

		# Increment denominator of the equivalent vibration equation
		denominator += timeGap

	# Calculate and return the equivalent vibration
	return math.sqrt(numerator/denominator)

def calculateEquivalentVibrationSingleAxis(timeStamps, weightedVibration):
	''' This function takes a time series of weighted vibration data for a single axis as well as their assosciated timestamps. Using these, it calculates the equivalent vibration magnitude according to SANS 2631-1 (Appendix C, equ C.1).
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

		# print("Starting")
		# startTime = time.time()
		# while ((time.time() - startTime) < 5):
		# 	pass

		responseMessage = comfort_service_api_v1_pb2.ComfortResponse()
		responseMessage.unix_time.extend(request.unix_time)
		
		# Calculate equivalent vibration
		if (len(request.unix_time) > 1):
			equivalentVibration = calculateEquivalentVibrationMultiAxis(request.unix_time, request.human_weighted_vibration_x, request.human_weighted_vibration_y, request.human_weighted_vibration_z)
		else:
			equivalentVibration = math.sqrt((request.human_weighted_vibration_x[0]**2)+(request.human_weighted_vibration_y[0]**2)+(request.human_weighted_vibration_z[0]*2))

		# Assess comfort
		responseMessage = assessComfort(equivalentVibration, responseMessage)

		# print("Sending")
		return responseMessage
	
def loadTLSCredentials(certDirectory):
		# This function loads in the generated TLS credentials from file, creates
	# a server credentials object with the key and certificate, and  returns 
	# that object for use in the server connection
	
	serverKeyFile = f"{certDirectory}/server-key.pem"
	serverCertFile = f"{certDirectory}/server-cert.pem"
	caCertFile = f"{certDirectory}/ca-cert.pem"

	# Load the server's certificate and private key
	private_key = open(serverKeyFile).read()
	certificate_chain = open(serverCertFile).read()

	# Load certificates of the CA who signed the client's certificate
	certificate_pool = open(caCertFile).read()

	credentials = grpc.ssl_server_credentials(
		private_key_certificate_chain_pairs = [(bytes(private_key, 'utf-8'), bytes(certificate_chain, 'utf-8'))],
		root_certificates = certificate_pool,
		require_client_auth = True
	)
	
	return credentials
	
def serve():
	''' This function creates a server with specified interceptors, registers the service calls offered by that server, and exposes
	the server over a specified port.
	'''

	# Create interceptor chain
	activeInterceptors = [
		metricInterceptor.MetricInterceptor("ComfortService"), 
		authenticationInterceptor.AuthenticationInterceptor(
			config["authentication"]["jwt"]["secretKey"], 
			config["authentication"]["jwt"]["tokenDuration"], 
			{config["authentication"]["accessLevel"]["name"]["comfortRating"]: config["authentication"]["accessLevel"]["role"]["comfortRating"]}
		),
		rateLimitInterceptor.RateLimitInterceptor(4)
	] # List containing the interceptors to be chained
	
	# Create a server to serve calls in its own thread
	server = grpc.server(
		futures.ThreadPoolExecutor(max_workers=10),
		interceptors = activeInterceptors
	)
	logging.debug("Successfully created server.")

	# Register an ocean weather service on the server
	comfort_service_api_v1_pb2_grpc.add_ComfortServiceServicer_to_server(ComfortServiceServicer(), server)
	logging.debug("Successfully registered comfort service to server.")

	# Load TLS credentials
	creds = loadTLSCredentials("certification")

	# Create a secure connection on port
	comfortHost = os.getenv(key = "COMFORTHOST", default = "[::]") # Receives the hostname from the environmental variables (for Docker network), or defaults to localhost for local testing
	try:
		# server.add_insecure_port(f'{comfortHost}:{config["port"]["myself"]}')
		server.add_secure_port(f'{comfortHost}:{config["port"]["myself"]}', creds) # server.add_insecure_port(f'{comfortHost}:{config["port"]["myself"]}')
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