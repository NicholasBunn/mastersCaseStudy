# Standard library imports
import logging

# Third party imports
import grpc
from grpc_interceptor import ServerInterceptor

# Logger setup
try:
	logger = logging.getLogger(__name__)
except:
	print("Unable to initialise log file, good luck :)")

class RateLimitInterceptor(ServerInterceptor):

	def __init__(self, callLimit):
		logger.debug("Intialising Rate Limit Interceptor")

		self.callLimit = callLimit
		self.callCounter = 0
		
	def intercept(self, method, request, context, methodName):
		self.callCounter += 1
		
		if (self.callCounter > self.callLimit):
			return grpc.StatusCode.RESOURCE_EXHAUSTED
		else:
			response =  method(request, context)
			self.callCounter -= 1

			return response