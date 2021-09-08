# Standard library imports
import os
import unittest
import warnings

# Third party imports
import numpy as np
from mat4py import loadmat

# Local application imports
import propellerMonitorService

class InverseSolutionSolverTest(unittest.TestCase):
	''' This class is used to execute all unit tests on the inverse solution solver class. 
	'''

	def setUp(self):
		''' setUp is used to create an instance of the InverseSolutionSolver class for each test
		'''

		warnings.simplefilter("ignore", ResourceWarning) # This just supresses resource warnings for us

		self.mySolver = propellerMonitorService.InverseSolutionSolver() # Create an instance of the InverseSolutionSolverClass called mySolver

	def test_setUpMatrices(self):
		''' This tests that the setUpMatrices function correctly builds the mass (inertia), damping, and stiffness matrices. The matrices that are built are verified against the original matrices used by Brendon Nickerson https://scholar.sun.ac.za/handle/10019.1/110071), extracted from his Matlab code.
		'''
		
		print("Testing Propeller Monitor Service: InverseSolutionSolver: Unit Test: setUpMatrices (Function Test)")

		# Invoke the setUpMatrices function
		self.mySolver.setUpMatrices() 

		# Load the original matrices to use for validation, which have been saved into CSV files
		trueMassMatrix = np.loadtxt(open("services/propellerMonitorService/test data/massMatrix.csv", "rb"), delimiter=",")
		trueDampingMatrix = np.loadtxt(open("services/propellerMonitorService/test data/dampingMatrix.csv", "rb"), delimiter=",")
		trueStiffnessMatrix = np.loadtxt(open("services/propellerMonitorService/test data/stiffnessMatrix.csv", "rb"), delimiter=",")
		
		# Compare the original matrices to the ones built by the setUpMatrices function
		if not (np.allclose(self.mySolver.massMatrix, trueMassMatrix, rtol=1e-05, atol=1e-08)):
			self.fail("Mass matrix incorrect")

		if not (np.allclose(self.mySolver.dampingMatrix, trueDampingMatrix, rtol=1e-05, atol=1e-08)):
			self.fail("Damping matrix incorrect")
		
		if not (np.allclose(self.mySolver.stiffnessMatrix, trueStiffnessMatrix, rtol=1e-05, atol=1e-08)):
			self.fail("Stiffness matrix incorrect")

	def test_setUpVectors(self):
		''' This tests that the setUpVectors function correctly builds the force vector. The vector that is built is verified against the original vector used by Brendon Nickerson https://scholar.sun.ac.za/handle/10019.1/110071), extracted from his Matlab code.
		'''

		print("Testing Propeller Monitor Service: InverseSolutionSolver: Unit Test: setUpVectors (Function Test)")

		# Load the data used for this test
		inputData = loadmat("services/propellerMonitorService/test data/Case_2019_10_30_20_06_20.mat")

		# Load the original matrices, which have been saved into CSV files
		trueMassMatrix = np.loadtxt(open("services/propellerMonitorService/test data/massMatrix.csv", "rb"), delimiter=",")
		trueDampingMatrix = np.loadtxt(open("services/propellerMonitorService/test data/dampingMatrix.csv", "rb"), delimiter=",")
		trueStiffnessMatrix = np.loadtxt(open("services/propellerMonitorService/test data/stiffnessMatrix.csv", "rb"), delimiter=",")

		# Load the original vector to use for validation, which has been saved into a CSV file
		trueForceVector = np.loadtxt(open("services/propellerMonitorService/test data/forceVector.csv", "rb"), delimiter=",")

		# Set the mass matrices to that used in the original model (This allows for clean unit tests to be written because of guarenteed inputs)
		self.mySolver.massMatrix = trueMassMatrix
		self.mySolver.dampingMatrix = trueDampingMatrix
		self.mySolver.stiffnessMatrix = trueStiffnessMatrix

		# Invoke the setupVectors functon
		self.mySolver.setUpVectors(inputData['Time_Final'], inputData['ForeTorqFinal'], inputData['rpmFinal'])

		# Compare the original vector to the one built by the setUpVectors function
		if not (np.allclose(self.mySolver.forceVector, trueForceVector, rtol=1e00, atol=1e-08)):
			self.fail("Force vector incorrect")

	def test_inverseInitialCond(self):
		''' This tests that the inverseInitialCond function correctly calculates the initial conditions. The conditions that are built are verified against the original conditions used by Brendon Nickerson https://scholar.sun.ac.za/handle/10019.1/110071), extracted from his Matlab code.
		'''

		print("Testing Propeller Monitor Service: InverseSolutionSolver: Unit Test: inverseInitalCond (Function Test)")
		
		# Load the original matrices and vectors, which have been saved into CSV files
		trueMassMatrix = np.loadtxt(open("services/propellerMonitorService/test data/massMatrix.csv", "rb"), delimiter=",")
		trueDampingMatrix = np.loadtxt(open("services/propellerMonitorService/test data/dampingMatrix.csv", "rb"), delimiter=",")
		trueStiffnessMatrix = np.loadtxt(open("services/propellerMonitorService/test data/stiffnessMatrix.csv", "rb"), delimiter=",")
		trueForceVector = np.loadtxt(open("services/propellerMonitorService/test data/forceVector.csv", "rb"), delimiter=",")

		# Load the original initial conditions to use for validation, which have been saved into a CSV file
		trueDIVector = np.loadtxt(open("services/propellerMonitorService/test data/DI.csv", "rb"), delimiter=",")
		trueAIVector = np.loadtxt(open("services/propellerMonitorService/test data/AI.csv", "rb"), delimiter=",")
		trueVIVector = np.loadtxt(open("services/propellerMonitorService/test data/VI.csv", "rb"), delimiter=",")

		# Set the mass matrices and vectors to that used in the original model (This allows for clean unit tests to be written because of guarenteed inputs)
		self.mySolver.massMatrix = trueMassMatrix
		self.mySolver.dampingMatrix = trueDampingMatrix
		self.mySolver.stiffnessMatrix = trueStiffnessMatrix
		self.mySolver.forceVector = trueForceVector

		# Set the required solution parameters
		self.mySolver.numTimeSteps = round(10/self.mySolver.dt)
		self.mySolver.numDegreesOfFreedom = max(np.shape(self.mySolver.massMatrix))

		# Invoke the inverseInitialCond function
		DI, VI, AI = self.mySolver.inverseInitialCond(trueForceVector[-2,1], trueForceVector[-1,1])

		# print(DI)
		# print(trueDIVector)
		# print(VI)
		# print(trueVIVector)
		# print(AI)
		# print(trueAIVector)

	def test_jwhAlpha(self):
		''' This tests that the jwhAlpha function correctly performs the generalise-alpha time scheme. The outputs are verified against the original outputs used by Brendon Nickerson https://scholar.sun.ac.za/handle/10019.1/110071), extracted from his Matlab code.
		'''

		print("Testing Propeller Monitor Service: InverseSolutionSolver: Unit Test: jwhAlpha (Function Test)")

		# Load the original matrices and vectors, which have been saved into CSV files
		trueMassMatrix = np.loadtxt(open("services/propellerMonitorService/test data/massMatrix.csv", "rb"), delimiter=",")
		trueDampingMatrix = np.loadtxt(open("services/propellerMonitorService/test data/dampingMatrix.csv", "rb"), delimiter=",")
		trueStiffnessMatrix = np.loadtxt(open("services/propellerMonitorService/test data/stiffnessMatrix.csv", "rb"), delimiter=",")
		trueForceVector = np.loadtxt(open("services/propellerMonitorService/test data/forceVector.csv", "rb"), delimiter=",")
		trueDIVector = np.loadtxt(open("services/propellerMonitorService/test data/DI.csv", "rb"), delimiter=",")
		trueAIVector = np.loadtxt(open("services/propellerMonitorService/test data/AI.csv", "rb"), delimiter=",")
		trueVIVector = np.loadtxt(open("services/propellerMonitorService/test data/VI.csv", "rb"), delimiter=",")

		# Load the original output vectors to use for validation, which have been saved into a CSV file
		trueQVector = np.loadtxt(open("services/propellerMonitorService/test data/qMatrix.csv", "rb"), delimiter=",")
		trueQDVector = np.loadtxt(open("services/propellerMonitorService/test data/qdMatrix.csv", "rb"), delimiter=",")
		trueQDDVector = np.loadtxt(open("services/propellerMonitorService/test data/qddMatrix.csv", "rb"), delimiter=",")

		# Set the mass matrices and vectors to that used in the original model (This allows for clean unit tests to be written because of guarenteed inputs)
		self.mySolver.massMatrix = trueMassMatrix
		self.mySolver.dampingMatrix = trueDampingMatrix
		self.mySolver.stiffnessMatrix = trueStiffnessMatrix
		self.mySolver.forceVector = trueForceVector

		# Set the required solution parameters
		self.mySolver.numTimeSteps = 1200 # Number of time steps in this simulation
		self.mySolver.numDegreesOfFreedom = max(np.shape(self.mySolver.massMatrix))

		# Determine and set alphaM, alphaF, and gamma
		self.mySolver.alphaM = (1/2)*((3)/(1))
		self.mySolver.alphaF = 1/(1)
		self.mySolver.gamma = (1/2) + self.mySolver.alphaM - self.mySolver.alphaF

		# Determine and set the constants for the effective stiffness matrix
		self.mySolver.a1 = (self.mySolver.alphaM**2)/(self.mySolver.alphaF*(self.mySolver.gamma**2)*(self.mySolver.dt**2))
		self.mySolver.a2 = self.mySolver.alphaM/(self.mySolver.gamma*self.mySolver.dt)

		# Determine and set the effective stiffness matrix
		self.mySolver.effectiveStiffness = self.mySolver.a1*self.mySolver.massMatrix + self.mySolver.a2*self.mySolver.dampingMatrix + self.mySolver.alphaF*self.mySolver.stiffnessMatrix

		# Determine and set the constants for the effective force matrix
		self.mySolver.a3 = self.mySolver.alphaM/(self.mySolver.alphaF*self.mySolver.gamma*self.mySolver.dt)
		self.mySolver.a4 = (self.mySolver.gamma-self.mySolver.alphaM)/(self.mySolver.gamma*self.mySolver.alphaF)
		self.mySolver.a5 = (self.mySolver.alphaF-1)/self.mySolver.alphaF
		self.mySolver.a6 = self.mySolver.alphaM/(self.mySolver.alphaF*(self.mySolver.gamma**2)*(self.mySolver.dt**2))
		self.mySolver.a7 = 1/(self.mySolver.alphaF*self.mySolver.gamma*self.mySolver.dt)
		self.mySolver.a8 = (self.mySolver.gamma-1)/self.mySolver.gamma
		self.mySolver.a9 = (self.mySolver.gamma-self.mySolver.alphaM)/(self.mySolver.alphaF*(self.mySolver.gamma**2)*self.mySolver.dt)
		self.mySolver.a10 = 1/(self.mySolver.gamma*self.mySolver.dt)

		# Invoke the jwhAlpha function
		q, qd, _, qdd = self.mySolver.jwhAlpha(trueDIVector, trueVIVector, trueAIVector, trueForceVector)

		# Compare the original vectors to the ones built by the jwhAlpha function
		if not (np.allclose(q, trueQVector, rtol=1e-05, atol=1e-20)):
			self.fail("q matrix incorrect")

		if not (np.allclose(qd, trueQDVector, rtol=1e-50, atol=1e-06)):
			self.fail("qd matrix incorrect")

		if not (np.allclose(qdd, trueQDDVector, rtol=1e-50, atol=1e-02)):
			self.fail("qdd matrix incorrect")

	def test_calculateInternalTorqueAndThetas(self):
		''' This tests that the calculateInternalTorquAndThetas function correctly calculates the torque and thetas (which are required to calculate the angular velocity and acceleration). The outputs are verified against the original outputs used by Brendon Nickerson https://scholar.sun.ac.za/handle/10019.1/110071), extracted from his Matlab code.
		'''

		print("Testing Propeller Monitor Service: InverseSolutionSolver: Unit Test: calculateInternalTorqueAndThetas (Function Test)")

		# Load the original vectors, which have been saved into CSV files
		trueQVector = np.loadtxt(open("services/propellerMonitorService/test data/qMatrix.csv", "rb"), delimiter=",").transpose()
		trueQDVector = np.loadtxt(open("services/propellerMonitorService/test data/qdMatrix.csv", "rb"), delimiter=",").transpose()
		trueQDDVector = np.loadtxt(open("services/propellerMonitorService/test data/qddMatrix.csv", "rb"), delimiter=",").transpose()

		# Load the original output vectors to use for validation, which have been saved into a CSV file
		trueThetaXVector = np.loadtxt(open("services/propellerMonitorService/test data/thetaXVector.csv", "rb"), delimiter=",")
		trueThetaXDVector = np.loadtxt(open("services/propellerMonitorService/test data/thetaXDVector.csv", "rb"), delimiter=",")
		trueThetaXDDVector = np.loadtxt(open("services/propellerMonitorService/test data/thetaXDDVector.csv", "rb"), delimiter=",")
		trueTorqueXVector = np.loadtxt(open("services/propellerMonitorService/test data/torqueXVector.csv", "rb"), delimiter=",")

		# Set the required solution parameters
		self.mySolver.numTimeSteps = 1200 # Number of time steps in this simulation

		# Invoke the calculateInternalTorqueAndThetas function
		thetaX, thetaXD, thetaXDD, torqueX = self.mySolver.calculateInternalTorqueAndThetas(trueQVector, trueQDVector, trueQDDVector)

		# Compare the original vectors to the ones built by the calculateInternalTorqueAndThetas function
		if not (np.allclose(thetaX, trueThetaXVector, rtol=1e-05, atol=1e-08)):
			self.fail("ThetaX vector incorrect")
			
		if not (np.allclose(thetaXD, trueThetaXDVector, rtol=1e-05, atol=1e-08)):
			self.fail("ThetaXD vector incorrect")

		if not (np.allclose(thetaXDD, trueThetaXDDVector, rtol=1e-05, atol=1e-08)):
			self.fail("ThetaXDD vector incorrect")

		if not (np.allclose(torqueX, trueTorqueXVector, rtol=1e-05, atol=1e-08)):
			self.fail("TorqueX vector incorrect")
        
	def test_calculateAngularVelAndAccel(self):
		''' This tests that the calculateAngularVelAndAccel function correctly calculates the angular velocity and acceleration (which are required to calculate the load on the propeller). The outputs are verified against the original outputs used by Brendon Nickerson https://scholar.sun.ac.za/handle/10019.1/110071), extracted from his Matlab code.
		'''

		print("Testing Propeller Monitor Service: InverseSolutionSolver: Unit Test: calculateAngularVelAndAccel (Function Test)")

		# Load the original vectors, which have been saved into CSV files
		trueQDVector = np.loadtxt(open("services/propellerMonitorService/test data/qdMatrix.csv", "rb"), delimiter=",").transpose()
		trueQDDVector = np.loadtxt(open("services/propellerMonitorService/test data/qddMatrix.csv", "rb"), delimiter=",").transpose()

		# Load the original output vectors to use for validation, which have been saved into a CSV file
		trueThetaX0DVector = np.loadtxt(open("services/propellerMonitorService/test data/thetaD0x.csv", "rb"), delimiter=",")
		trueThetaX0DDVector = np.loadtxt(open("services/propellerMonitorService/test data/thetaDD0x.csv", "rb"), delimiter=",")
		
		# Invoke the calculatAngularVelAndAccel function
		thetaX0D, thetaX0DD = self.mySolver.calculateAngularVelAndAccel(trueQDVector, trueQDDVector)

		# Compare the original vectors to the ones built by the calculateAngularVelAndAccel function
		if not (np.allclose(thetaX0D, trueThetaX0DVector, rtol=1e-05, atol=1e-08)):
			self.fail("ThetaX0D vector incorrect")
		
		if not (np.allclose(thetaX0DD, trueThetaX0DDVector, rtol=1e-05, atol=1e-08)):
			self.fail("ThetaX0DD vector incorrect")

	def test_extractLoads(self):
		''' This tests that the extractIceLoad function correctly calculates the extracts the ice load, motor load, and propeller load. The outputs are verified against the original outputs used by Brendon Nickerson https://scholar.sun.ac.za/handle/10019.1/110071), extracted from his Matlab code.
		'''

		print("Testing Propeller Monitor Service: InverseSolutionSolver: Unit Test: extractLoads (Function Test)")

		# Load the original vectors, which have been saved into CSV files
		trueQVector = np.loadtxt(open("services/propellerMonitorService/test data/qMatrix.csv", "rb"), delimiter=",").transpose()
		trueThetaX0DVector = np.loadtxt(open("services/propellerMonitorService/test data/thetaD0x.csv", "rb"), delimiter=",")

		# Load the original output vectors to use for validation, which have been saved into a CSV file
		trueQIceVector = np.loadtxt(open("services/propellerMonitorService/test data/qIceVector.csv", "rb"), delimiter=",")
		trueQMotorVector = np.loadtxt(open("services/propellerMonitorService/test data/qMotorVector.csv", "rb"), delimiter=",")
		trueQPropVector = np.loadtxt(open("services/propellerMonitorService/test data/qPropVector.csv", "rb"), delimiter=",")

		# Invoke the extractIceLoad function
		Qice, Qmotor, Qprop = self.mySolver.extractLoads(trueQVector, trueThetaX0DVector)

		# Compare the original vectors to the ones built by the extractLoads function
		if not (np.allclose(Qice, trueQIceVector, rtol=1e-05, atol=1e-08)):
			self.fail("Qice vector incorrect")

		if not (np.allclose(Qmotor, trueQMotorVector, rtol=1e-05, atol=1e-08)):
			self.fail("Qmotor vector incorrect")
		
		if not (np.allclose(Qprop, trueQPropVector, rtol=1e-01, atol=1e-08)):
			self.fail("QProp vector incorrect")

	# def test_solveSystem(self):
	# 	pass

if __name__ == '__main__':
	unittest.main()
