# Standard library imports
import logging

# Third party imports
import grpc
from grpc_interceptor import ServerInterceptor
import jwt

# Logger setup
try:
	logger = logging.getLogger(__name__)
except:
    print("Unable to initialise log file, good luck :)")

class AuthenticationInterceptor(ServerInterceptor):
	
	def __init__(self, secretKey, tokenDuration, authenticatedMethods):
		logger.debug("Initialising authentication interceptor")

		self.secretKey = secretKey # Initialise a property with the secret key for decrypting the access token
		self.authenticatedMethods = authenticatedMethods # Initialise a property with the methods that require authentication
	
	def authorise(self, methodName, context):
		'''This function goes through a series of checks to verify that the user
		making a request is properly authenticated for that request.
		'''

		# Check if the requested method requires authentication
		accessibleRoles = self.authenticatedMethods.get(methodName)
		if (accessibleRoles == None):
			logger.info(f"Authentication is not required for {methodName}")
			return
			
		# Extract the metadata from the request
		try:
			metadata = dict(context.invocation_metadata())
		except:
			logger.debug("Failed to authenticate: metadata is not provided")
			raise grpc.StatusCode.PERMISSION_DENIED

		# Check if a JWT has been included in the metadata
		try:
			encodedToken = metadata["authorisation"]
		except:
			logger.debug("Failed to authenticate: JWT has not been provided")
			raise grpc.StatusCode.PERMISSION_DENIED

		# Check that the provided JWT is valid
		claims, err = self.verifyJWT(encodedToken) # Extract the user claims from the token
		if (err != None):
			logger.debug("Failed to authenticate: Provided JWT is invalid")
			raise err

		''' Check that the role of the user making the service call authenticates them for the 
		service being called
		'''
		for role in accessibleRoles:
			if (role == claims["role"]):
				logger.debug(f"Successfully authenticated request for {methodName}")
				return

		logger.debug('''Failed to authenticate: the user does not have permission to access the 
		requested service''')
		raise grpc.StatusCode.PERMISSION_DENIED

	def verifyJWT(self, accessToken):
		'''This function takes a JWT token as an input and decodes it to ensure that it is valid (according to the HS256 algorithm).
		'''

		try:
			token = jwt.decode(accessToken, self.secretKey, algorithms=["HS256"])
		except Exception as e:
			logger.debug(f"Invalid token: {e}")
			raise grpc.StatusCode.PERMISSION_DENIED
		
		return token, None

	def intercept(self, method, request, context, methodName):
		'''This is the function that runs when the call is received.
		'''
		
		logger.info("Starting server-side authentication interceptor")

		# Attempt to authorise the user for the request
		err = self.authorise(methodName, context)
		if err:
			context.set_code(err)
			return None
		
		# Return the result of the service/method call
		return method(request, context)