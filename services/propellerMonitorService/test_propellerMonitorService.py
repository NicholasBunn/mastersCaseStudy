# Standard library imports
import os
import unittest
import warnings

# Local application imports
import propellerMonitorService

# Third party imports
import numpy as np
from mat4py import loadmat


class InverseSolutionSolverTest(unittest.TestCase):
	''' This class is used to execute all tests on the solver class.
	'''

	def setUp(self):
		''' setUp
		'''

		warnings.simplefilter("ignore", ResourceWarning)

		self.mySolver = propellerMonitorService.InverseSolutionSolver()

	def test_setUpMatrices(self):
		self.mySolver.setUpMatrices()

		trueMassMatrix = np.loadtxt(open("services/propellerMonitorService/test data/massMatrix.csv", "rb"), delimiter=",")
		trueDampingMatrix = np.loadtxt(open("services/propellerMonitorService/test data/dampingMatrix.csv", "rb"), delimiter=",")
		trueStiffnessMatrix = np.loadtxt(open("services/propellerMonitorService/test data/stiffnessMatrix.csv", "rb"), delimiter=",")
		
		if (np.allclose(self.mySolver.massMatrix, trueMassMatrix, rtol=1e-05, atol=1e-08)):
			print("PASS")
		else:
			self.fail("Mass matrix incorrect")

		if (np.allclose(self.mySolver.dampingMatrix, trueDampingMatrix, rtol=1e-05, atol=1e-08)):
    			print("PASS")
		else:
			self.fail("Damping matrix incorrect")
		
		if (np.allclose(self.mySolver.stiffnessMatrix, trueStiffnessMatrix, rtol=1e-05, atol=1e-08)):
			print("PASS")
		else:
			self.fail("Stiffness matrix incorrect")

	def test_setUpVectors(self):
		inputData = loadmat("services/propellerMonitorService/test data/Case_2019_10_30_20_06_20.mat")

		trueMassMatrix = np.loadtxt(open("services/propellerMonitorService/test data/massMatrix.csv", "rb"), delimiter=",")
		trueDampingMatrix = np.loadtxt(open("services/propellerMonitorService/test data/dampingMatrix.csv", "rb"), delimiter=",")
		trueStiffnessMatrix = np.loadtxt(open("services/propellerMonitorService/test data/stiffnessMatrix.csv", "rb"), delimiter=",")
		trueForceVector = np.loadtxt(open("services/propellerMonitorService/test data/forceVector.csv", "rb"), delimiter=",")

		self.mySolver.massMatrix = trueMassMatrix
		self.mySolver.dampingMatrix = trueDampingMatrix
		self.mySolver.stiffnessMatrix = trueStiffnessMatrix

		self.mySolver.setUpVectors(inputData['Time_Final'], inputData['ForeTorqFinal'], inputData['rpmFinal'])

		if (np.allclose(self.mySolver.forceVector, trueForceVector, rtol=1e00, atol=1e-08)):
				print("PASS")
		else:
			self.fail("Force vector incorrect")

	def test_inverseInitialCond(self):
		
		trueMassMatrix = np.loadtxt(open("services/propellerMonitorService/test data/massMatrix.csv", "rb"), delimiter=",")
		trueDampingMatrix = np.loadtxt(open("services/propellerMonitorService/test data/dampingMatrix.csv", "rb"), delimiter=",")
		trueStiffnessMatrix = np.loadtxt(open("services/propellerMonitorService/test data/stiffnessMatrix.csv", "rb"), delimiter=",")
		trueForceVector = np.loadtxt(open("services/propellerMonitorService/test data/forceVector.csv", "rb"), delimiter=",")

		trueDIVector = np.loadtxt(open("services/propellerMonitorService/test data/DI.csv", "rb"), delimiter=",")
		trueAIVector = np.loadtxt(open("services/propellerMonitorService/test data/AI.csv", "rb"), delimiter=",")
		trueVIVector = np.loadtxt(open("services/propellerMonitorService/test data/VI.csv", "rb"), delimiter=",")

		self.mySolver.massMatrix = trueMassMatrix
		self.mySolver.dampingMatrix = trueDampingMatrix
		self.mySolver.stiffnessMatrix = trueStiffnessMatrix
		self.mySolver.forceVector = trueForceVector

		duration = 10 # Duration of simulation [seconds]
		self.mySolver.numTimeSteps = round(duration/self.mySolver.dt) # Number of time steps in this simulation
		timeVector = np.linspace(0, duration, self.mySolver.numTimeSteps)
		self.mySolver.numDegreesOfFreedom = max(np.shape(self.mySolver.massMatrix))

		DI, VI, AI = self.mySolver.inverseInitialCond(trueForceVector[-2,1], trueForceVector[-1,1])

		# print(DI)
		# print(trueDIVector)
		# print(VI)
		# print(trueVIVector)
		# print(AI)
		# print(trueAIVector)

	def test_jwhAlpha(self):
		trueMassMatrix = np.loadtxt(open("services/propellerMonitorService/test data/massMatrix.csv", "rb"), delimiter=",")
		trueDampingMatrix = np.loadtxt(open("services/propellerMonitorService/test data/dampingMatrix.csv", "rb"), delimiter=",")
		trueStiffnessMatrix = np.loadtxt(open("services/propellerMonitorService/test data/stiffnessMatrix.csv", "rb"), delimiter=",")
		trueForceVector = np.loadtxt(open("services/propellerMonitorService/test data/forceVector.csv", "rb"), delimiter=",")
		trueDIVector = np.loadtxt(open("services/propellerMonitorService/test data/DI.csv", "rb"), delimiter=",")
		trueAIVector = np.loadtxt(open("services/propellerMonitorService/test data/AI.csv", "rb"), delimiter=",")
		trueVIVector = np.loadtxt(open("services/propellerMonitorService/test data/VI.csv", "rb"), delimiter=",")

		trueQVector = np.loadtxt(open("services/propellerMonitorService/test data/qMatrix.csv", "rb"), delimiter=",")
		trueQDVector = np.loadtxt(open("services/propellerMonitorService/test data/qdMatrix.csv", "rb"), delimiter=",")
		trueQDDVector = np.loadtxt(open("services/propellerMonitorService/test data/qddMatrix.csv", "rb"), delimiter=",")

		self.mySolver.massMatrix = trueMassMatrix
		self.mySolver.dampingMatrix = trueDampingMatrix
		self.mySolver.stiffnessMatrix = trueStiffnessMatrix
		self.mySolver.forceVector = trueForceVector

		self.mySolver.numTimeSteps = 1200 # Number of time steps in this simulation
		# timeVector = np.linspace(0, duration, self.mySolver.numTimeSteps)
		self.mySolver.numDegreesOfFreedom = max(np.shape(self.mySolver.massMatrix))

		# Determine self.alphaM, self.alphaF, and self.gamma
		self.mySolver.alphaM = (1/2)*((3)/(1))
		self.mySolver.alphaF = 1/(1)
		self.mySolver.gamma = (1/2) + self.mySolver.alphaM - self.mySolver.alphaF

		# Constants for effective stiffnedd matrix
		self.mySolver.a1 = (self.mySolver.alphaM**2)/(self.mySolver.alphaF*(self.mySolver.gamma**2)*(self.mySolver.dt**2))
		self.mySolver.a2 = self.mySolver.alphaM/(self.mySolver.gamma*self.mySolver.dt)

		# Constants for effective force
		self.mySolver.a3 = self.mySolver.alphaM/(self.mySolver.alphaF*self.mySolver.gamma*self.mySolver.dt)
		self.mySolver.a4 = (self.mySolver.gamma-self.mySolver.alphaM)/(self.mySolver.gamma*self.mySolver.alphaF)
		self.mySolver.a5 = (self.mySolver.alphaF-1)/self.mySolver.alphaF
		self.mySolver.a6 = self.mySolver.alphaM/(self.mySolver.alphaF*(self.mySolver.gamma**2)*(self.mySolver.dt**2))
		self.mySolver.a7 = 1/(self.mySolver.alphaF*self.mySolver.gamma*self.mySolver.dt)
		self.mySolver.a8 = (self.mySolver.gamma-1)/self.mySolver.gamma
		self.mySolver.a9 = (self.mySolver.gamma-self.mySolver.alphaM)/(self.mySolver.alphaF*(self.mySolver.gamma**2)*self.mySolver.dt)
		self.mySolver.a10 = 1/(self.mySolver.gamma*self.mySolver.dt)

		self.mySolver.effectiveStiffness = self.mySolver.a1*self.mySolver.massMatrix + self.mySolver.a2*self.mySolver.dampingMatrix + self.mySolver.alphaF*self.mySolver.stiffnessMatrix
		
		q, qd, qdVer, qdd = self.mySolver.jwhAlpha(trueDIVector, trueVIVector, trueAIVector, trueForceVector)

		if (np.allclose(q, trueQVector, rtol=1e-05, atol=1e-20)):
			print("PASS")
		else:
			self.fail("q matrix incorrect")

		if (np.allclose(qd, trueQDVector, rtol=1e-50, atol=1e-06)):
			print("PASS")
		else:
			self.fail("qd matrix incorrect")

		if (np.allclose(qdd, trueQDDVector, rtol=1e-50, atol=1e-02)):
    			print("PASS")
		else:
			self.fail("qdd matrix incorrect")

	def test_calculateInternalTorqueAndThetas(self):
		trueQVector = np.loadtxt(open("services/propellerMonitorService/test data/qMatrix.csv", "rb"), delimiter=",").transpose()
		trueQDVector = np.loadtxt(open("services/propellerMonitorService/test data/qdMatrix.csv", "rb"), delimiter=",").transpose()
		trueQDDVector = np.loadtxt(open("services/propellerMonitorService/test data/qddMatrix.csv", "rb"), delimiter=",").transpose()

		trueThetaXVector = np.loadtxt(open("services/propellerMonitorService/test data/thetaXVector.csv", "rb"), delimiter=",")
		trueThetaXDVector = np.loadtxt(open("services/propellerMonitorService/test data/thetaXDVector.csv", "rb"), delimiter=",")
		trueThetaXDDVector = np.loadtxt(open("services/propellerMonitorService/test data/thetaXDDVector.csv", "rb"), delimiter=",")
		trueTorqueXVector = np.loadtxt(open("services/propellerMonitorService/test data/torqueXVector.csv", "rb"), delimiter=",")

		self.mySolver.numTimeSteps = 1200 # Number of time steps in this simulation

		thetaX, thetaXD, thetaXDD, torqueX = self.mySolver.calculateInternalTorqueAndThetas(trueQVector, trueQDVector, trueQDDVector)

		if (np.allclose(thetaX, trueThetaXVector, rtol=1e-05, atol=1e-08)):
    			print("PASS")
		else:
			self.fail("ThetaX vector incorrect")
			
		if (np.allclose(thetaXD, trueThetaXDVector, rtol=1e-05, atol=1e-08)):
    			print("PASS")
		else:
			self.fail("ThetaXD vector incorrect")

		if (np.allclose(thetaXDD, trueThetaXDDVector, rtol=1e-05, atol=1e-08)):
    			print("PASS")
		else:
			self.fail("ThetaXDD vector incorrect")

		if (np.allclose(torqueX, trueTorqueXVector, rtol=1e-05, atol=1e-08)):
    			print("PASS")
		else:
			self.fail("TorqueX vector incorrect")
        
	def test_calculateAngularVelAndAccel(self):
		trueQDVector = np.loadtxt(open("services/propellerMonitorService/test data/qdMatrix.csv", "rb"), delimiter=",").transpose()
		trueQDDVector = np.loadtxt(open("services/propellerMonitorService/test data/qddMatrix.csv", "rb"), delimiter=",").transpose()

		trueThetaX0DVector = np.loadtxt(open("services/propellerMonitorService/test data/thetaD0x.csv", "rb"), delimiter=",")
		trueThetaX0DDVector = np.loadtxt(open("services/propellerMonitorService/test data/thetaDD0x.csv", "rb"), delimiter=",")
		
		thetaX0D, thetaX0DD = self.mySolver.calculateAngularVelAndAccel(trueQDVector, trueQDDVector)

		if (np.allclose(thetaX0D, trueThetaX0DVector, rtol=1e-05, atol=1e-08)):
    			print("PASS")
		else:
			self.fail("ThetaX0D vector incorrect")
		
		if (np.allclose(thetaX0DD, trueThetaX0DDVector, rtol=1e-05, atol=1e-08)):
    			print("PASS")
		else:
			self.fail("ThetaX0DD vector incorrect")

	def test_extractIceLoad(self):
		trueQVector = np.loadtxt(open("services/propellerMonitorService/test data/qMatrix.csv", "rb"), delimiter=",").transpose()
		trueThetaX0DVector = np.loadtxt(open("services/propellerMonitorService/test data/thetaD0x.csv", "rb"), delimiter=",")

		trueQIceVector = np.loadtxt(open("services/propellerMonitorService/test data/qIceVector.csv", "rb"), delimiter=",")
		trueQMotorVector = np.loadtxt(open("services/propellerMonitorService/test data/qMotorVector.csv", "rb"), delimiter=",")
		trueQPropVector = np.loadtxt(open("services/propellerMonitorService/test data/qPropVector.csv", "rb"), delimiter=",")

		Qice, Qmotor, Qprop = self.mySolver.extractIceLoad(trueQVector, trueThetaX0DVector)

		if (np.allclose(Qice, trueQIceVector, rtol=1e-05, atol=1e-08)):
    			print("PASS")
		else:
			self.fail("Qice vector incorrect")

		if (np.allclose(Qmotor, trueQMotorVector, rtol=1e-05, atol=1e-08)):
    			print("PASS")
		else:
			self.fail("Qmotor vector incorrect")
		
		if (np.allclose(Qprop, trueQPropVector, rtol=1e-01, atol=1e-08)):
    			print("PASS")
		else:
			self.fail("QProp vector incorrect")

	# def test_solveSystem(self):
	# 	pass

if __name__ == '__main__':
	unittest.main()
