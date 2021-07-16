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
import proto.v1.generated.process_vibration_service_api_v1_pb2 as process_vibration_service_api_v1_pb2
import proto.v1.generated.process_vibration_service_api_v1_pb2_grpc as process_vibration_service_api_v1_pb2_grpc

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

	def CalculateRMS(self, request, context):
		"""The 'Calculate RMS' call calculates the root mean square (RMS) vibration for time-series vibration signals."""

		logging.info("Received Calculate RMS service call.")

		# Create the response message
		responseMessage = process_vibration_service_api_v1_pb2.ProcessResponse()

		# Iterate through request fields, calculating RMS values and appending to the response
		for count, time in enumerate(request.unix_time, 0):
    		# Set the time field for the response
			responseMessage.unix_time.append(time)

			# Calculate RMS
			responseMessage.rms_vibration_x.append(request.vibration_x[count]/math.sqrt(2))
			responseMessage.rms_vibration_y.append(request.vibration_y[count]/math.sqrt(2))
			responseMessage.rms_vibration_z.append(request.vibration_z[count]/math.sqrt(2))

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
	process_vibration_service_api_v1_pb2_grpc.add_ProcessVibrationServiceServicer_to_server(ProcessVibrationServiceServicer(), server)
	logging.debug("Successfully registered process vibration service to server.")

	# Create an insecure connection on port
	processVibrationHost = os.getenv(key = "PROCESSVIBRATIONHOST", default = "localhost") # Receives the hostname from the environmental variables (for Docker network), or defaults to localhost for local testing
	try:
		server.add_insecure_port(f'{processVibrationHost}:{config["port"]["myself"]}')
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