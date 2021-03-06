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
sys.path.append( os.path.dirname( os.path.dirname( os.path.dirname( os.path.abspath(__file__) ) ) ) )
import protoFiles.python.processVibrationService.v1.process_vibration_service_api_v1_pb2 as process_vibration_service_api_v1_pb2
import protoFiles.python.processVibrationService.v1.process_vibration_service_api_v1_pb2_grpc as process_vibration_service_api_v1_pb2_grpc
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

class ProcessVibrationServiceServicer(object):
	"""'Process Vibration Service' offers two service calls that process time-series vibration signals according to accepted practices."""

	def CalculateRMSSeries(self, request, context):
		"""The 'Calculate RMS Series' call calculates the root mean square (RMS) vibration for individual time-series vibration signals.
		"""

		logging.info("Received Calculate RMS Series service call.")

		# Create the response message
		responseMessage = process_vibration_service_api_v1_pb2.ProcessResponseSeries()

		# Iterate through request fields, calculating RMS values and appending to the response
		for count, time in enumerate(request.unix_time, 0):
    		# Set the time field for the response
			responseMessage.unix_time.append(time)

			# Calculate RMS
			responseMessage.rms_vibration_x.append(request.vibration_x[count]/math.sqrt(2))
			responseMessage.rms_vibration_y.append(request.vibration_y[count]/math.sqrt(2))
			responseMessage.rms_vibration_z.append(request.vibration_z[count]/math.sqrt(2))

		return responseMessage
	
	def CalculateRMSBatch(self,request, context):
		'''The 'Calculate RMS Batch' call calculates the root mean square (RMS) vibration for an "ensemble" value of a vibration signal time-series.
		'''

		logging.info("Received Calculate RMS Batch service call.")

		# Create the response message
		responseMessage = process_vibration_service_api_v1_pb2.ProcessResponseBatch()

		# Set initial ensemble values for the upcoming loop
		rmsEnsembleValueX = 0
		rmsEnsembleValueY = 0
		rmsEnsembleValueZ = 0

		# Iterate through request fields, summing the RMS vibrations into an ensemble value as it goes through
		for count, _ in enumerate(request.unix_time, 0):
			rmsEnsembleValueX += request.vibration_x[count]**2
			rmsEnsembleValueY += request.vibration_y[count]**2
			rmsEnsembleValueZ += request.vibration_z[count]**2
		
		# Set start time of ensemble value
		responseMessage.unix_time_start = request.unix_time[0]
		# Set end time of ensemble value
		responseMessage.unix_time_end = request.unix_time[-1]
		# Set ensemble RMS values
		responseMessage.rms_vibration_x = math.sqrt((1/len(request.unix_time)*rmsEnsembleValueX))
		responseMessage.rms_vibration_y = math.sqrt((1/len(request.unix_time)*rmsEnsembleValueY))
		responseMessage.rms_vibration_z = math.sqrt((1/len(request.unix_time)*rmsEnsembleValueZ))

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
		metricInterceptor.MetricInterceptor("ProcessVibrationService"), 
		authenticationInterceptor.AuthenticationInterceptor(
			config["authentication"]["jwt"]["secretKey"], 
			config["authentication"]["jwt"]["tokenDuration"], 
			{config["authentication"]["accessLevel"]["name"]["calculateRMSBatch"]: config["authentication"]["accessLevel"]["role"]["calculateRMSBatch"], config["authentication"]["accessLevel"]["name"]["calculateRMSSeries"]: config["authentication"]["accessLevel"]["role"]["calculateRMSSeries"]
			}
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
	process_vibration_service_api_v1_pb2_grpc.add_ProcessVibrationServiceServicer_to_server(ProcessVibrationServiceServicer(), server)
	logging.debug("Successfully registered process vibration service to server.")

	# Load TLS credentials
	creds = loadTLSCredentials("certification")

	# Create a secure connection on port
	processVibrationHost = os.getenv(key = "PROCESSVIBRATIONHOST", default = "[::]") # Receives the hostname from the environmental variables (for Docker network), or defaults to localhost for local testing
	try:
		server.add_secure_port(f'{processVibrationHost}:{config["port"]["myself"]}', creds) # server.add_insecure_port(f'{processVibrationHost}:{config["port"]["myself"]}')
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

	logging.basicConfig(filename="services/processVibrationService/program logs/" + serviceName + ".log", format="%(asctime)s:%(name)s:%(levelname)s:%(module)s:%(funcName)s:%(message)s", level=logging.DEBUG, force = True)

	# ________SERVE REQUESTS________
	serve() # Finish initialisation by serving the request