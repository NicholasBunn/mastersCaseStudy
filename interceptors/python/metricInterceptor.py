# Standard library imports
import logging
import time
import os

# Third party imports
import grpc
from grpc_interceptor import ServerInterceptor
import prometheus_client as prometheus

# Logger setup
try:
	logger = logging.getLogger(__name__)
except:
	print("Unable to initialise log file, good luck")

def pushToPrometheus(c, g, h, executionTime, serviceName, address, method, job, registry):
	'''This function sets the labels for a Prometheus entry and pushes metrics to the push-gateway.
	'''
	c.labels(Role="Server", grpc_type = 'unary', grpc_service = serviceName, grpc_method = method).inc() # Increment the call counter
	g.labels(Role="Server", grpc_type = 'unary', grpc_service = serviceName, grpc_method = method).set_to_current_time() # Set last call time to the current time
	h.labels(Role="Server", grpc_type = 'unary', grpc_service = serviceName, grpc_method = method).observe(executionTime) # Set the response latency to the execution time of the service
	
	# Post the metrics to the gateway
	prometheus.push_to_gateway(address, job=job, registry=registry)
	logger.info("Succesfully pushed metrics")

def sendMetrics(func):
	'''This decorator wraps the intercept method, setting the metrics before invoking the pushToPrometheus function.
	'''
	from functools import wraps

	@wraps(func)
	def wrapper(*args, **kw):
		logger.debug(" Starting Interceptor decorator")

		# Extract the service name and method from the incoming request
		if isinstance(args[3], grpc._server._Context):
			servicerContext = args[3]
			# This gives us <service>/<method name>
			serviceMethod = servicerContext._rpc_event.call_details.method
			serviceName, methodName = str(serviceMethod).rsplit('/')[1::]
		else:
			logger.warning('Cannot derive the service name and method')
			raise Exception("Could not derive service name or method.")
		
		try:
			# Set the start time of the call
			startTime = time.time()
			# Invoke the service call
			result = func(*args, **kw, )
			resultStatus = "Success"
			logger.debug("Function call: {}".format(resultStatus))
		except Exception:
			resultStatus = "Error"
			logger.warning("Function call: {}".format(resultStatus))
			raise
		finally:
			# Calculate the time since the start of the call
			responseTime = time.time() - startTime
			# Push metrics to the Prometheus server
			pushToPrometheus(args[0].c, args[0].g, args[0].h, responseTime, serviceName.rsplit(".")[2], args[0].address, methodName, args[0].microserviceName, args[0].registry)
		return result
	return wrapper

class MetricInterceptor(ServerInterceptor):
	pushGatewayaHost = os.getenv("PUSHGATEWAYHOST", "localhost") # Receives the hostname from the environmental variables for Docker, or defaults to localhost for local testing
	address = "http://" + pushGatewayaHost + ":9091" # Todo: pass/pull this from the message metadata

	def __init__(self, microserviceName):
		logger.debug("Initialising metric interceptor")

		self.microserviceName = microserviceName

		self.registry = prometheus.CollectorRegistry()
		self.c = prometheus.Counter("server_request_counter", "Number of times this API has been called", registry=self.registry, labelnames= ['Role', 'grpc_type', 'grpc_service', 'grpc_method'])
		self.g = prometheus.Gauge('server_last_call_time', 'Last time this API was called', registry=self.registry, labelnames= ['Role', 'grpc_type', 'grpc_service', 'grpc_method'])
		self.h = prometheus.Histogram('server_request_latency', 'Ammount of time for request to be processed', registry=self.registry, labelnames= ['Role', 'grpc_type', 'grpc_service', 'grpc_method'])

	@sendMetrics
	def intercept(self, method, request, context, methodName):
		'''This is the function that runs when the call is received. In this interceptor the logic is performed by a decorater, wrapping this function with the required functionality.
		'''
		
		logger.info("Starting server-side metric interceptor method")

		return method(request, context)   