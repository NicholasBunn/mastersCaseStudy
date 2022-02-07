# Standard library imports
import logging
import time

# Third party imports
import grpc
from grpc_interceptor import ClientCallDetails, ClientInterceptor

# Logger setup
try:
	logger = logging.getLogger(__name__)
except:
	print("Unable to initialise log file, good luck :)")

class RetryInterceptor(ClientInterceptor):
	
	def __init__(self, waitTime, maxRetries):
		logger.debug("Intialising Retry Interceptor.")

		self.failureCount = 0 # Initialise a property representing the number of failed calls to zero
		self.waitTime = waitTime # Initialise a property for how long to wait between the first failure
		self.maxRetries = maxRetries # Initialise a property with the maximum number of retries
		
	def intercept(
        self,
        method,
        request_or_iterator,
        call_details: ClientCallDetails,
    ):
		'''This is the function that runs when the call is received.
		'''

		while self.failureCount < (self.maxRetries - 1):
			logger.info(f"Trying to make a call, call number: {self.failureCount}")

			# Attempt to make the call
			response =  method(request_or_iterator, call_details)

			# Evaluate whether the response has an error
			if(hasattr(response, "_state") and (response._state.code)):
				print(f"Call attempt failed")
				time.sleep((self.waitTime/100)*(2**self.failureCount)) # Wait function that implements exponential backoff based on the number of failed calls
				self.failureCount += 1 # Increment the failed call counter
			else:
				return response
		return method(request_or_iterator, call_details)

			

		