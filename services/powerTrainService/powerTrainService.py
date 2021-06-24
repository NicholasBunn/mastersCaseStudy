# Standard library imports
import sys
import os
import logging
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

def loadConfigFile(filepath):
	with open(os.path.join(sys.path[0], filepath), "r") as f:
		config = yaml.safe_load(f)
		serverConfig = config["server"]
	return serverConfig

def structureData(dataSet):
	''' This function takes a (structured) dataFrame as an input, normalises and orders
	the data into the correct shape, as is required by the machine learning library, 
	before returning a numpy array containing the data
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
	''' This function takes the filename of a model as an input, loads the model, and returns the model object. NOTE: The model that is called is passed the absolute path as opposed to only the model name
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
	''' This function takes a model object and the model's inputs as arguments. It uses these to generate a power prediction from the model, returning the power estimate.
	'''

	# Receive a power estimate by producing an estimate using the modelInputs set of input parameters
	estimatedPower = myModel.predict(modelInputs)

	return estimatedPower

class PowerTrainServiceServicer(power_train_service_api_v1_pb2_grpc.PowerTrainServiceServicer):
	"""'Power Train Service; offers four service calls that provide information about the power train of the vessel (namely power requirements and their assosciated costs)
	"""

	def PowerEstimate(self, request, context):
		"""The 'Power Estimate' call provides foresight for tactical decision-making by providing power estimates for a requested route and sailing conditions
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
		"""The 'Cost Estimate' call provides foresight for tactical decision-making by providing cost estimates for a requested route and sailing conditions
		"""
		hourlyCrewCost = 10000 # Cost of crew salaries per hour, in R
		fuelDensity = 0.8323 # Density of diesel, in kg/litre
		dieselPrice = 13 # Price of Diesel, in R/litler
		fuelConsumption = 179 # Fuel consumption of the S.A. Agulhas II, in g/kWh
		costPerkW = (dieselPrice/fuelDensity)*(fuelConsumption/1000) # Cost of running the ship, in R/kWh


		# for unixTime, portPropMotorSpeed, sbtdPropMotorSpeed, propPitchPort, propPitchStbd, sog, relativeWindDirection, windSpeed, beaufortNumber, waveDirection, waveLength, modelType in zip(request.time_and_data, request.port_prop_motor_speed, request.stbd_prop_motor_speed, request.propeller_pitch_port, request.propeller_pitch_stbd, request.sog, request.wind_direction_relative, request.wind_speed, request.beaufort_number, request.wave_direction, request.wave_length, request.model_type):
    	# 	pass

		totalCost = hourlyCrewCost*(timeInHours) + (costPerkWh*powerInkW*timeInHours)

		context.set_code(grpc.StatusCode.UNIMPLEMENTED)
		context.set_details('Method not implemented!')
		raise NotImplementedError('Method not implemented!')

	def PowerTracking(self, request, context):
		"""The 'Power Tracking' call provides insight for tactical and operational decision-making by providing real-time power use by the vessel
		"""
		context.set_code(grpc.StatusCode.UNIMPLEMENTED)
		context.set_details('Method not implemented!')
		raise NotImplementedError('Method not implemented!')

	def PowerEstimateEvaluation(self, request, context):
		"""The 'Power Estimate Evaluation' call provdes hindsight for strategic decision-making by evaluating the accuracy of the models predictions
		"""
		context.set_code(grpc.StatusCode.UNIMPLEMENTED)
		context.set_details('Method not implemented!')
		raise NotImplementedError('Method not implemented!')

def serve():
	''' This function creates a server with specified interceptors, registers the service calls offered by that server, and exposes
	the server over a specified port. The connection to this port is secured with server-side TLS encryption. 
	'''

	# Create a server to serve calls in its own thread
	server = grpc.server(
		futures.ThreadPoolExecutor(max_workers=10),
		# interceptors = activeInterceptors
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

	logging.basicConfig(filename="services/powerTrainService/program logs/" + serviceName + ".log", format="%(asctime)s:%(name)s:%(levelname)s:%(module)s:%(funcName)s:%(message)s", level=logging.DEBUG)

	# ________SERVE REQUEST________
	serve() # Finish initialisation by serving the request
