# Standard library imports
import os
import sys
import time
import logging
import unittest

# Third party imports
import yaml
import keras
import numpy as np
import pandas as pd

# Local application imports
import powerTrainService

def loadConfigFile(filepath):
	with open(os.path.join(sys.path[0], filepath), "r") as f:
		config = yaml.safe_load(f)
	serverConfig = config["server"]
	return serverConfig

def importData(excelFileName):
	''' This function receives a filename ("filename.xlsx") as an input, reads it into a Pandas dataframe, and returns the generated dataFrame
	'''

	# Import ship and weather data for estimation
	dataSet = pd.read_excel(excelFileName, engine = "openpyxl") # This is a dataFrame

	return dataSet # NOTE: "dataSet" is a dataFrame

class PowerTrainServiceUnitTest(unittest.TestCase):

	def setUp(self):
		pass

	def tearDown(self):
		pass

	def test_structureData(self):
		''' This tests the structuring of data. It ensures that the data is fitted and shaped correctly
		'''

		print("Testing Power Train Service: Structure Data (Function Test)")

    	# Import test data
		testDataSet = importData("services/powerTrainService/test data/CMU_2019_2020_openWater.xlsx")

		# Structure test data
		structuredData = powerTrainService.structureData(testDataSet) # Save this as a class parameter to be used in the 'runModel' test

		self.assertEqual(structuredData[0,0], 0.597549569477592)
		self.assertEqual(structuredData[0,1], 0.60128600844793)
		self.assertEqual(structuredData[0,2], 0.26854406323815316)
		self.assertEqual(structuredData[0,3], 0.24256864992988544)
		self.assertEqual(structuredData[0,4], 0.030389908256880732)
		self.assertEqual(structuredData[0,5], 0.04178272980501393)
		self.assertEqual(structuredData[0,6], 0.07507507507507508)
		self.assertEqual(structuredData[0,7], 0.0)
		self.assertEqual(structuredData[0,8], 0.0)
		self.assertEqual(structuredData[0,9], 0.0)

	def test_loadModel(self):
		''' This tests the loading of models. It ensures that the models stored in the "required files" folder exist by their expected names, that the load function is operating effectively, and that the models are of the correct class (Keras sequential model class). NOTE: This function test that the models exist by the correct name, and not that the models are exactly the ones expected!
		'''

		print("Testing Power Train Service: Load Model (Function Test)")

		openWaterEstimationModel = powerTrainService.loadModel("OPENWATER") # Save this as a class parameter to be used in the 'runModel' test
		iceEstimationModel = powerTrainService.loadModel("ICE")
		unkownInputModel = powerTrainService.loadModel("UNKNOWN")

		self.assertIsInstance(openWaterEstimationModel, keras.engine.sequential.Sequential)
		self.assertIsInstance(iceEstimationModel, keras.engine.sequential.Sequential)
		self.assertIsInstance(unkownInputModel, keras.engine.sequential.Sequential)

	def test_runModel(self):
		''' This tests the running of the models. It ensures that the runModel function hasn't been updated in a way that breaks the original functionality.
		'''

		print("Testing Power Train Service: Run Model (Function Test)")

		model = keras.models.load_model("services/powerTrainService/required files/OpenWaterModel_R67.h5")
    
		testData = {'PortPropMotorSpeed': [0.597549569477592, 0.5991748807429235],
					'StbdPropMotorSpeed': [0.60128600844793, 0.6010362244160419],
					'PropellerPitchPort': [0.26854406323815316, 0.47962105209935035], 
					'PropellerPitchStbd': [0.24256864992988544, 0.4839979447640052],
					'SOG': [0.030389908256880732, 0.09518348623853208],
					'WindDirRel': [0.04178272980501393, 0.013927576601671309],
					'WindSpeed': [0.07507507507507508, 0.036036036036036036],
					'Beaufort number': [0.0, 0.0],
					'Wave direction': [0.0, 0.0],
					'Wave length': [0.0, 0.0]}

		testResults = powerTrainService.runModel(model, pd.DataFrame(testData))
		
		self.assertEqual(float(testResults[0,0]), 356.5527648925781)
		self.assertEqual(float(testResults[1,0]), 173.1065216064453)

class PowerTrainIntegrationTEst(unittest.TestCase):
	config = loadConfigFile("configuration.yaml")
	serverClass = powerTrainService.PowerTrainServiceServicer
	hostName = os.getenv(key = "POWERTRAINHOST", default="localhost")
	port = f'{hostName}:{config["port"]["myself"]}'

	def setUp(self):
		''' setUp is used to create a server instance to test each service call
		'''

		self.server = powerTrainService.grpc.server(powerTrainService.futures.ThreadPoolExecutor(max_workers=10))
		powerTrainService.power_train_service_api_v1_pb2_grpc.add_PowerTrainServiceServicer_to_server(self.serverClass(), self.server)
		self.server.add_insecure_port(self.port)
		self.server.start()
	
	def tearDown(self):
		''' tearDown is used to pull down the server instance that was used for testing
		'''

		self.server.stop(None)

	def test_PowerEstimate(self):
		''' This function tests the PowerEstimate-specific functionality. It ensures that the service call takes the correct inputs, processes them as is required for the model, and makes the correct estimates based on the provided inputs.
		'''

		print("Testing Power Train Service: Power Estimate (Service Call Test)")

		with powerTrainService.grpc.insecure_channel(self.port) as channel:
			stub = powerTrainService.power_train_service_api_v1_pb2_grpc.PowerTrainServiceStub(channel)
			response = stub.PowerEstimate(powerTrainService.power_train_service_api_v1_pb2.PowerTrainEstimateRequest(
				port_prop_motor_speed = [83.5450057983399, 83.7725067138672, 120.443740844727],
				stbd_prop_motor_speed = [84.4112548828125, 84.3762435913086, 120.522499084473],
				propeller_pitch_port = [-40.5200004577637, 0.0300000011920929, 95.3400039672852],
				propeller_pitch_stbd = [-46.2599983215332, 0.359999984502792, 92.3999938964844],
				sog = [0.545311111064, 1.707955555408, 13.756244443256],
				wind_direction_relative = [15.0, 5.0, 332],
				wind_speed = [2.5, 1.2, 13.6],
				beaufort_number = [0.0, 0.0, 3],
				wave_direction = [0.0, 0.0, 255],
				wave_length = [0.0, 0.0, 69.3333333333333],
				model_type = powerTrainService.power_train_service_api_v1_pb2.OPENWATER,
			))
			
		self.assertEqual(response.power_estimate[0], 415.4642639160156)
		self.assertEqual(response.power_estimate[1], -6.391994476318359)
		self.assertEqual(response.power_estimate[2], 3718.072509765625)


	def test_CostEstimate(self):
		pass

	def test_PowerTracking(self):
		pass

	def test_PowerEstimateEvaluation(self):
		pass

if __name__ == '__main__':
	unittest.main()
