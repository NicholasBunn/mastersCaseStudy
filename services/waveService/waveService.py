#Package imports
    # Standard library imports
    # Third party imports
    # Local application imports


import sys
import os
import  yaml
import logging
from concurrent import futures

import grpc
from apis.waveServiceAPI.v1.python import waveServiceAPI_v1_pb2
import waveServiceAPI_v1_pb2_grpc
from interceptors.python import metricInterceptor
from interceptors.python import authenticationInterceptor

def loadConfigFile(filepath):
    with open(os.path.join(sys.path[0], filepath), "r") as f:
        config = yaml.safe_load(f)
        serverConfig = config["server"]
    return serverConfig

def loadTLSCredentials():
    	# This function loads in the generated TLS credentials from file, creates
	# a server credentials object with the key and certificate, and  returns 
	# that object for use in the server connection
	
	serverKeyFile = "certification/server-key.pem"
	serverCertFile = "certification/server-cert.pem"
	caCertFile = "certification/ca-cert.pem"

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

class WaveServiceServicer(waveServiceAPI_v1_pb2_grpc.WaveServiceServicer):
    """'Wave Service' offers two service calls that provide information about wave conditions for use in route planning.
    """

    def WaveEstimate(self, request, context):
        """The 'Wave Estimate' call provides foresight for tactical decision-making by providing future wave conditions along a requested route
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def WaveHistory(self, request, context):
        """The 'Wave History' call provides hindsight for stategic decision-making by providing historical wave conditions that the ship would have encountered along a requested route
        """
        
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def serve():
    ''' This function creates a server with specified interceptors, registers the service calls offered by that server, and exposes
    the server over a specified port. The connection to this port is secured with server-side TLS encryption. 
    '''

    # Create metric interceptor
    callMetrics = metricInterceptor.MetricInterceptor()
    # Create authentication interceptor
    callAuthentication = authenticationInterceptor.AuthenticationInterceptor(config["authentication"]["jwt"][["secretKey"], config["authentication"]["jwt"]]["tokenDuration"], {config["authentication"]["accessLevel"]["name"]["waveEstimate"]: config["authentication"]["accessLevel"]["role"]["waveEstimate"], config["authentication"]["accessLevel"]["name"]["waveHistory"]: config["authentication"]["accessLevel"]["role"]["waveHistory"]})
    # Create a list with the interceptors to be chained
    activeInterceptors = [callMetrics, callAuthentication] 

	# Create a server to serve calls in its own thread
    server = grpc.server(
        futures.ThreadPoolExecutor(max_workers=10),
        interceptors = activeInterceptors
    )

    # Register a wave service on the server
    waveServiceAPI_v1_pb2_grpc.add_WaveServiceServicer_to_server(WaveServiceServicer(), server)

    # Create an insecure connection on port
    fetchDataHost = os.getenv(key = "FETCHDATAHOST", default = "localhost") # Receives the hostname from the environmental variables (for Docker network), or defaults to localhost for local testing
    server.add_insecure_port(f'{fetchDataHost}:{config["port"]["myself"]}')

    # Create a secure (TLS encrypted) connection on port
    # creds = loadTLSCredentials()
    # fetchDataHost = os.getenv(key = "FETCHDATAHOST", default = "localhost") # Receives the hostname from the environmental variables (for Docker network), or defaults to localhost for local testing
    # server.add_secure_port(f'{fetchDataHost}:{config["port"]["myself"]}', creds)

    # Start server and listen for calls on the specified port
    server.start()
    logger.info(f'Server started on port {config["port"]["myself"]}')
	
    # Defer termination for a 'persistent' service
    server.wait_for_termination()

if __name__ == '__main__':
	# ________LOAD CONFIG FILE________
	config = loadConfigFile("configuration.yaml")

	# ________LOGGER SETUP________
	serviceName = __file__.rsplit("/")[-2].rsplit(".")[0]
	logger = logging.getLogger(serviceName)
	logger.setLevel(logging.DEBUG)

	# Set the fields to be included in the logs
	formatter = logging.Formatter('%(asctime)s:%(name)s:%(levelname)s:%(module)s:%(funcName)s:%(message)s')

	# Create/set the file in which the log will be stored
	fileHandler = logging.FileHandler("program logs/" + serviceName + ".log")
	fileHandler.setFormatter(formatter)

	logger.addHandler(fileHandler)

	# ________SERVE REQUEST________
	serve() # Finish initialisation by serving the request