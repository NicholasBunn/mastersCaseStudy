# Standard library imports
import sys
import os
import logging
import time
from concurrent import futures

# Third party imports
import grpc
import yaml
import numpy as np
import pandas as pd
from keras import models
from sklearn.preprocessing import MinMaxScaler

# Local application imports
import proto.v1.generated.power_train_service_api_v1_pb2 as power_train_service_api_v1_pb2
import proto.v1.generated.power_train_service_api_v1_pb2_grpc as power_train_service_api_v1_pb2_grpc
sys.path.append( os.path.dirname( os.path.dirname( os.path.dirname( os.path.abspath(__file__) ) ) ) )
import interceptors.python.metricInterceptor as metricInterceptor
import interceptors.python.authenticationInterceptor as authenticationInterceptor

def loadConfigFile(filepath):
	''' This function reads in a YAML configuration file. It takes the relative filepath as an input. It returns a dictionary (?) containing configuration variables.
	'''
    
	with open(os.path.join(sys.path[0], filepath), "r") as f:
		config = yaml.safe_load(f)
		serverConfig = config["server"]

	logging.debug("Succesfully loaded configuration file")
	return serverConfig

def structureData(dataSet):
	''' This function takes a (structured) dataFrame as an input, normalises and orders
	the data into the correct shape, as is required by the machine learning library, 
	before returning a numpy array containing the data.
	'''

	dataSet.shape # Shape the test data before accessing its parameters

	# ________NORMALISE THE DATA________
	# Transform par 1 - Port Propellor Speed (measured using the motor speed)
	scaler = MinMaxScaler()
	scaler.fit(dataSet['PortPropMotorSpeed'].values.reshape(-1,1))
	parameter1 = scaler.transform(dataSet['PortPropMotorSpeed'].values.reshape(-1,1))

	# Transform par 2 - Starboard Propellor Speed (measured using the motor speed)
	scaler.fit(dataSet['StbdPropMotorSpeed'].values.reshape(-1,1))
	parameter2 = scaler.transform(dataSet['StbdPropMotorSpeed'].values.reshape(-1,1))

	# Transform par 3 - Port Propellor Pitch
	scaler.fit(dataSet['PropellerPitchPort'].values.reshape(-1,1))
	parameter3 = scaler.transform(dataSet['PropellerPitchPort'].values.reshape(-1,1))

	# Transform par 4 - Starboard Propellor Pitch
	scaler.fit(dataSet['PropellerPitchStbd'].values.reshape(-1,1))
	parameter4 = scaler.transform(dataSet['PropellerPitchStbd'].values.reshape(-1,1))

	# Transform par 5 - Ship Speed Over Ground (SOG)
	scaler.fit(dataSet['SOG'].values.reshape(-1,1))
	parameter5 = scaler.transform(dataSet['SOG'].values.reshape(-1,1))

	# Transform par 6 - Wind Direction Relative to the Ship's Heading
	scaler.fit(dataSet['WindDirRel'].values.reshape(-1,1))
	parameter6 = scaler.transform(dataSet['WindDirRel'].values.reshape(-1,1))

	# Transform par 7 - Wind Speed
	scaler.fit(dataSet['WindSpeed'].values.reshape(-1,1))
	parameter7 = scaler.transform(dataSet['WindSpeed'].values.reshape(-1,1))

	# Transform par 8 - Beaufort Number
	scaler.fit(dataSet['Beaufort number'].values.reshape(-1,1))
	parameter8 = scaler.transform(dataSet['Beaufort number'].values.reshape(-1,1))

	# Transform par 9 - Wave Direction
	scaler.fit(dataSet['Wave direction'].values.reshape(-1,1))
	parameter9 = scaler.transform(dataSet['Wave direction'].values.reshape(-1,1))

	# Transform par 10 - Wave Length
	scaler.fit(dataSet['Wave length'].values.reshape(-1,1))
	parameter10 = scaler.transform(dataSet['Wave length'].values.reshape(-1,1))

	# ________SHAPE THE DATA FOR THE ML LIBRARY________
	X1 = np.reshape(parameter1,-1)	# Port propeller speed
	X2 = np.reshape(parameter2,-1)	# Starboard propeller speed
	X3 = np.reshape(parameter3,-1)	# Port propeller pitch
	X4 = np.reshape(parameter4,-1)	# Starboard propeller pitch
	X5 = np.reshape(parameter5,-1)	# SOG
	X6 = np.reshape(parameter6,-1)	# Relative wind direction
	X7 = np.reshape(parameter7,-1)	# Wind speed
	X8 = np.reshape(parameter8,-1)	# Beaufort number
	X9 = np.reshape(parameter9,-1)	# Wave direction
	X10 = np.reshape(parameter10,-1)	# Wave length

	# ________BUILD THE PARAMETERS________
	parameters = (X1, X2, X3, X4, X5, X6, X7, X8, X9, X10)

	modelInputs = np.transpose(parameters)

	modelInputs.shape

	return modelInputs

def loadModel(modelType):
	''' This function takes the filename of a model as an input, loads the model, and returns the Keras sequential model object. NOTE: The model that is called is passed the absolute path as opposed to only the model name.
	'''

	def modelSelector(argument):
		switcher = {
			0: "services/powerTrainService/required files/OpenWaterModel_R67.h5", # If no model is supplied, assume open water operation
			1: "services/powerTrainService/required files/OpenWaterModel_R67.h5", 
			2: "services/powerTrainService/required files/IceModel_R58.h5",
		}
		return switcher.get(argument, "services/powerTrainService/required files/OpenWaterModel_R67.h5") # Again, if no model is supplied, assume open water operation

	# MEEP do I actually use this switcher?
	try:
		workableModel = models.load_model(modelSelector(modelType))  # Import the model that was passed as an argument
		logging.debug(f"Successfully loaded {str(modelType)}")
		# MEEP "modelType" doesn't return the text representation
		return workableModel
	except ImportError as error:
		logging.error("Model file not available")
		raise error
	except IOError as error:
		logging.error("Invalid file")
		raise error
	except Exception as error:
		logging.error("Unaccounted error loading model")
		raise error

def runModel(myModel, modelInputs):
	''' This function takes a Keras sequential model object and the model's inputs as arguments. It uses these to generate a power prediction from the model, returning the power estimate.
	'''

	# Receive a power estimate by producing an estimate using the modelInputs set of input parameters
	estimatedPower = myModel.predict(modelInputs)

	return estimatedPower

def evaluateModel(myModel, modelInputs, actualPortMotorPower, actualStbdMotorPower):
	''' This function takes a Keras sequential model object, the model's inputs, and the actual power for evaluation as inputs. It evaluates the model's prediction against	the actual power, returning the real power.
	'''

	realPower = (actualPortMotorPower + actualStbdMotorPower)/2 # realPower holds the actual (average) power, as recorded by the MCU, used here to compare to the model's estimates

	# Evaluate the model's estimate against the actual power
	scores = myModel.evaluate(modelInputs, realPower, verbose=0)
	print("%s: %.2f%%" % (myModel.metrics_names[1], scores[1]))

	return scores

class PowerTrainServiceServicer(power_train_service_api_v1_pb2_grpc.PowerTrainServiceServicer):
	"""'PowerTrainService' offers four service calls that provide information about the power train of the vessel (namely power requirements and/or their assosciated costs).
	"""

	def PowerEstimate(self, request, context):
		"""The 'PowerEstimate' call provides foresight for tactical decision-making (As is described by https://www.researchgate.net/publication/332173693_Designing_Ship_Digital_Services) by providing power estimates for a requested route and sailing conditions. It structures the input data for a data-driven model, laods the model in, and produces a power estimate using it.
		"""

		logging.info("Received Power Estimate service call.")

		# Create the response message
		responseMessage = power_train_service_api_v1_pb2.PowerEstimateResponse()

		# Map the request message data to a dictionary
		inputData = {'PortPropMotorSpeed': request.port_prop_motor_speed, 
					'StbdPropMotorSpeed': request.stbd_prop_motor_speed, 
					'PropellerPitchPort': request.propeller_pitch_port, 
					'PropellerPitchStbd': request.propeller_pitch_stbd, 
					'SOG': request.sog, 
					'WindDirRel': request.wind_direction_relative, 
					'WindSpeed': request.wind_speed, 
					'Beaufort number': request.beaufort_number, 
					'Wave direction':  request.wave_direction, 
					'Wave length': request.wave_length}
		
		# ________PREPARE THE DATA________
		try:
			processedData = structureData(pd.DataFrame(inputData))
			logging.debug("Successfully structured data")
		except Exception as e:
			logging.debug(f"Failed to structure the data: \n{e}")
			context.set_code(grpc.StatusCode.INTERNAL)
			# context.set_details("bla bla")
			raise e

		# ________LOADING A PRE-TRAINED MODEL_______
		try:
			activeModel = loadModel(request.model_type)
			logging.debug("Successfully loaded model")
		except Exception as e:
			logging.debug(f"Failed to load model: \n{e}")
			context.set_code(grpc.StatusCode.INTERNAL)
			# context.set_details("bla bla")
			raise e

		# ________RUN THE LOADED MODEL_______
		try:
			estimatedPower = runModel(activeModel, pd.DataFrame(processedData))
			logging.debug("Succesfully ran the model")
		except Exception as e:
			logging.debug(f"Failed to run estimation model: \n{e}")
			context.set_code(grpc.StatusCode.INTERNAL)
			# context.set_details("bla bla")
			raise e

		# Populate response message
		responseMessage.unix_time.extend(request.unix_time)
		responseMessage.power_estimate.extend(estimatedPower)

		return responseMessage

	def CostEstimate(self, request, context):
		"""The 'CostEstimate' call provides foresight for tactical decision-making (As is described by https://www.researchgate.net/publication/332173693_Designing_Ship_Digital_Services) by providing cost estimates for a requested route and sailing conditions. It invokes the PowerEstimate service call to produce a power estimate, and then calculates the cost assosciated with the provided route's power requirement profile.
		"""

		logging.info("Received Power Estimate service call.")

		# Create the response message
		responseMessage = power_train_service_api_v1_pb2.CostEstimateResponse()

		startingCost = 100000 # Cost to start a voyage, in R
		hourlyCrewCost = 10000 # Cost of crew salaries per hour, in R
		fuelDensity = 0.8323 # Density of diesel, in kg/litre
		dieselPrice = 13 # Price of Diesel, in R/litler
		fuelConsumption = 179 # Fuel consumption of the S.A. Agulhas II, in g/kWh
		costPerkWh = (dieselPrice/fuelDensity)*(fuelConsumption/1000) # Cost of running the ship, in R/kWh

		requiredPowerSet = self.PowerEstimate(request, context)

		# Loop through requested points
		for count, powerEstimate in enumerate(requiredPowerSet.power_estimate):
			if(count != 0):
    			# For every requested point after the first point, work out how long it's been since the previous point. Assume constant operating conditions for that time period, calculate the cost for the power required over that period, and increment the total cost by that. 

				timeSpan = requiredPowerSet.unix_time[count] - requiredPowerSet.unix_time[count-1] # The time (in seconds) between the current timestamp and the previous one
				additionalCost = hourlyCrewCost*( timeSpan/3600 ) + ( costPerkWh*powerEstimate*(timeSpan/3600) ) # The additional cost incurred over the above time period

				# Increment the total cost 
				totalCost += additionalCost

				# Add the new values to the response message
				responseMessage.unix_time.append(requiredPowerSet.unix_time[count])
				responseMessage.power_estimate.append(powerEstimate)
				responseMessage.cost_estimate.append(additionalCost)
			else:
    			# For the first point, set the total cost as the starting cost
				totalCost = startingCost 

				# Add the new values to the response message
				responseMessage.unix_time.append(requiredPowerSet.unix_time[count])
				responseMessage.power_estimate.append(powerEstimate)
				responseMessage.cost_estimate.append(startingCost)

		responseMessage.total_cost = totalCost

		return responseMessage

	def PowerTracking(self, request, context):
		"""The 'Power Tracking' call provides insight for tactical and operational decision-making (As is described by https://www.researchgate.net/publication/332173693_Designing_Ship_Digital_Services) by providing real-time power use by the vessel.
		"""
		context.set_code(grpc.StatusCode.UNIMPLEMENTED)
		context.set_details('Method not implemented!')
		raise NotImplementedError('Method not implemented!')

	def PowerEstimateEvaluation(self, request, context):
		"""The 'Power Estimate Evaluation' call provdes hindsight for strategic decision-making (As is described by https://www.researchgate.net/publication/332173693_Designing_Ship_Digital_Services) by evaluating the accuracy of the models predictions.
		"""
		context.set_code(grpc.StatusCode.UNIMPLEMENTED)
		context.set_details('Method not implemented!')
		raise NotImplementedError('Method not implemented!')

def serve():
	''' This function creates a server with specified interceptors, registers the service calls offered by that server, and exposes
	the server over a specified port.
	'''

	# Create interceptor chain
	activeInterceptors = [
		metricInterceptor.MetricInterceptor(), 
		authenticationInterceptor.AuthenticationInterceptor(
			config["authentication"]["jwt"]["secretKey"], 
			config["authentication"]["jwt"]["tokenDuration"], 
			{config["authentication"]["accessLevel"]["name"]["powerEstimate"]: config["authentication"]["accessLevel"]["role"]["powerEstimate"]}
		)
	] # List containing the interceptors to be chained
	
	# Create a server to serve calls in its own thread
	server = grpc.server(
		futures.ThreadPoolExecutor(max_workers=10),
		interceptors = activeInterceptors
	)
	logging.debug("Successfully created server.")

	# Register an ocean weather service on the server
	power_train_service_api_v1_pb2_grpc.add_PowerTrainServiceServicer_to_server(PowerTrainServiceServicer(), server)
	logging.debug("Successfully registered power train service to server.")

	# Create an insecure connection on port
	powerTrainHost = os.getenv(key = "POWERTRAINHOST", default = "localhost") # Receives the hostname from the environmental variables (for Docker network), or defaults to localhost for local testing
	try:
		server.add_insecure_port(f'{powerTrainHost}:{config["port"]["myself"]}')
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

	logging.basicConfig(filename="services/powerTrainService/program logs/" + serviceName + ".log", format="%(asctime)s:%(name)s:%(levelname)s:%(module)s:%(funcName)s:%(message)s", level=logging.DEBUG, force=True)

	# ________SERVE REQUEST________
	serve() # Finish initialisation by serving the request
