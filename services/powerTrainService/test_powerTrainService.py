# Standard library imports
import os
import sys
import time
import unittest

# Third party imports
import yaml
import keras
import numpy as np
import pandas as pd

# Local application imports
import powerTrainService

# def loadConfigFile(filepath):
# 	with open(os.path.join(sys.path[0], filepath), "r") as f:
# 		config = yaml.safe_load(f)
# 	serverConfig = config["server"]
# 	return serverConfig

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
		openWaterEstimationModel = powerTrainService.loadModel("OPENWATER") # Save this as a class parameter to be used in the 'runModel' test
		iceEstimationModel = powerTrainService.loadModel("ICE")
		unkownInputModel = powerTrainService.loadModel("UNKNOWN")

		self.assertIsInstance(openWaterEstimationModel, keras.engine.sequential.Sequential)
		self.assertIsInstance(iceEstimationModel, keras.engine.sequential.Sequential)
		self.assertIsInstance(unkownInputModel, keras.engine.sequential.Sequential)

	def test_runModel(self):
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

if __name__ == '__main__':
	unittest.main()
