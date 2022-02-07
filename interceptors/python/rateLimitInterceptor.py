# To anyone looking at this without context - do not use this for a production 
# application! This was thrown together very roughly to test a hypothesis and has 
# not been written with actual implementation in mind.

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
		logger.debug("Intialising Rate Limit Interceptor.")

		self.callLimit = callLimit # Initialise a property with the call threshold
		self.callCounter = 0 # Initialise a property for the number of calls to 0
		
	def intercept(self, method, request, context, methodName):
		'''This is the function that runs when the call is received.
		'''

		self.callCounter += 1 # Increment the number of active calls
		
		# Check if call counter exceeds the specified call limit
		if (self.callCounter > self.callLimit):
			logger.error("Concurrent request limit reached.")
			""" Decrement the number of active calls as the
			current call is not going to be made 
			"""
			self.callCounter -= 1 
			context.set_code(grpc.StatusCode.RESOURCE_EXHAUSTED)
			context.set_details('Concurrent request limit reached.')
			return None
		else:
    		# Invoke the requested method
			response =  method(request, context)
			''' Decrement the numebr of active calls as the current call has been
			processed '''
			self.callCounter -= 1

			# Return the response
			return response