import math
import numpy as np
from mat4py import loadmat
import matplotlib.pyplot as plt
import smoothN

# outerD = 0.5 # Outer diamete of shaft [m]
# innerD = 0.175 # Inner diameter of shaft [m]

# rho = 7850 # Density of steel
# self.J = ( (math.pi)*( (outerD**4)-(innerD**4) )/32 ) # Polar modulus of shaft section [m^3]
# self.Jp = 1.347*(10**4) # self.Cp Propeller and hydrodynamic mass moment of inertia [kg.m^2]
# self.Jm = 4.415*(10**3) # Motor mass moment of inertia [kg.m^2]
# self.Js = 5.120*(10**2) # SKF coupling mass moment of inertia [kg.m^2]
# self.L = 29.473 # Length of shaft [m]
# self.G = 81*(10**9) # Shear modulus of elasticity [Pa]
# self.xt = 25 # Torque measurement location along the shaft [m] (This changes depending on voyage config!)
# self.Cp = 1.136*(10**5) # self.Cp propeller hydrodynamic damping [Nm.s/rad]
# Jh = 3.592*(10**3) # Hydropdynamic mass moment of inertia [kg.m^2]
# self.xv = 25 # Velocity measurement location along the sahft [m] (This changes depending on voyage config!)

class InverseSolutionSolver():

    def __init__(self):
        ## Solution Parameters
        self.fs = 600 # Sample rate [Hz]
        self.dt = 1/self.fs # Simulation time step size
        self.xt = 25 # Torque measurement location along the shaft [m] (This changes depending on voyage config!)
        self.xv = 25 # Velocity measurement location along the sahft [m] (This changes depending on voyage config!)

        ## Specifications

        # Shaft Properties
        self.outerD = 0.5 # Outer diamete of shaft [m]
        self.innerD = 0.175 # Inner diameter of shaft [m]
        self.G = 81*(10**9) # Shear modulus of elasticity [Pa]
        self.L = 29.473 # Length of shaft [m]
        self.rho = 7850 # Density of steel
        self.J = ( (math.pi)*( (self.outerD**4)-(self.innerD**4) )/32 ) # Polar modulus of shaft section [m^3]

        # Propeller Properties
        self.Cp = 1.136*(10**5) # self.Cp propeller hydrodynamic damping [Nm.s/rad]
        self.Jp = 1.347*(10**4) # self.Cp Propeller and hydrodynamic mass moment of inertia [kg.m^2]
        self.Jh = 3.592*(10**3) # Hydropdynamic mass moment of inertia [kg.m^2]

        # Motor Properties
        self.Jm = 4.415*(10**3) # Motor mass moment of inertia [kg.m^2]

        # Sleeve Coupling Properties
        self.Js = 5.120*(10**2) # SKF coupling mass moment of inertia [kg.m^2]        

    def setUpMatrices(self):
        # ________CALCULATE MASS MATRIX________

        self.massMatrix = np.zeros(104)

        # Populate diagonal with flexible mode inertias for uneven mode entries
        for i in range(1, 102, 2):
            self.massMatrix[i] = self.rho * self.J * (self.L/2)

        # Add diagonal terms for rigid mode
        self.massMatrix[0] = (self.rho*self.J*self.L) + self.Jp + self.Jm + self.Js

        # Generate full matrix from populated diagonal
        self.massMatrix = np.diag(self.massMatrix)

        # Populate additional inertial terms for rigid mode row/column
        for i in range(1, 102, 2):
            self.massMatrix[0,i] = self.massMatrix[0,i] + self.Jp - self.Jm + (self.Js*math.cos((i)*math.pi*21.525/self.L))
            self.massMatrix[i,0] = self.massMatrix[i,0] + self.Jp - self.Jm + (self.Js*math.cos((i)*math.pi*21.525/self.L))

        # Populate additional inertial terms for flexible mode rows/columns
        for i in range(1, 102, 2):
            for self.J in range(1, 102, 2):
                self.massMatrix[i,self.J] = self.massMatrix[i,self.J] + self.Jp + self.Jm + self.Js*math.cos((i)*math.pi*21.525/self.L)*math.cos((self.J)*math.pi*21.525/self.L)

        # Delete even rows and columns
        self.massMatrix = np.delete(self.massMatrix, range(2, self.massMatrix.shape[0]-2, 2), axis=0)
        self.massMatrix = np.delete(self.massMatrix, range(2, self.massMatrix.shape[1]-2, 2), axis=1)

        # Add in additional rows/columns for measurement equations 
        self.massMatrix[int(102/2+1),:] = 0
        self.massMatrix[int(102/2+2),:] = 0
        self.massMatrix[:,int(102/2+1)] = 0
        self.massMatrix[:,int(102/2+2)] = 0

        # ________CALCULATE STIFFNESS MATRIX________

        # Main diagonal of full size matrix (inlc. rows/columns for even modes)
        self.stiffnessMatrix = np.zeros(104)

        # Populate diagonal with flexible mode stiffness (sixtch term eq. 5.41) for uneven mode entries 
        for i in range(1, 102, 2):
            self.stiffnessMatrix[i] = self.G*self.J*(((i*math.pi)**2)/(2*self.L))

        self.stiffnessMatrix[0] = 1 # Placeholder so that this entry doesn't get removed in following line

        # Set last two rows equal to one so that they're not deleted, either
        self.stiffnessMatrix[-2:] = 1

        # Remove even modes (all those equal to 0)
        self.stiffnessMatrix = self.stiffnessMatrix[self.stiffnessMatrix != 0]

        # Generate matrix from the populated diagonal
        self.stiffnessMatrix = np.diag(self.stiffnessMatrix)

        # Set rigid term back to zero (this is what it should be)
        self.stiffnessMatrix[0] = 0

        # Add in torque measurment equation (5.38)
        stiffnessTorque = np.zeros(102)
        for i in range(0, 102-1, 2):
            stiffnessTorque[i+1] = -self.G*self.J*math.sin((i+1)*math.pi*self.xt/self.L)*(i+1)*math.pi/self.L

        stiffnessTorque[0] = 1 # Placeholder so that this entry doesn't get removed in following line

        # Remove even modes (all those equal to 0)
        stiffnessTorque = stiffnessTorque[stiffnessTorque != 0]

        # Set rigid term back to zero (this is what it should be)
        stiffnessTorque[0] = 0

        # Include torque measurement in second last row of stiffness matrix
        self.stiffnessMatrix[-2, 0:-2] = stiffnessTorque
        self.stiffnessMatrix[-1, :] = 0

        # Include Qice and Qmotor terms as last columns of stiffness matrix
        for i in range(0, 52):
            self.stiffnessMatrix[i, -1] = 1*(10**9)
            self.stiffnessMatrix[i, -2] = 1*(10**9)

        self.stiffnessMatrix[0, -1] = -1*(10**9)
        self.stiffnessMatrix[-2, -2] = 0

        # ________CALCULATE DAMPING MATRIX________

        # Generate matrix for uneven modes and extra 2 rows/columns
        self.dampingMatrix = np.zeros([54, 54])

        # Populate rigid and uneven flexible modes damping (third and fourth terms in eq. 5.40 and 5.41)
        for row in range(0, 52, 1):
            for column in range(0, 52, 1):
                self.dampingMatrix[row, column] = self.Cp

        # Create angular velocity measurement equation (eq. 5.39)
        dampingVelocity = np.zeros(102)
        dampingVelocity[0] = 1

        for i in range(1, 102, 2):
            dampingVelocity[i] = math.cos(i*math.pi*self.xv/self.L)

        # Remove even modes and raise/scale the matrix (all those equal to 0)
        dampingVelocity = dampingVelocity[dampingVelocity != 0] * (10**7)

        # Incorporate velocity measurement in last row of damping matrix
        for i in range(0, 52, 1):
            self.dampingMatrix[53,i] = dampingVelocity[i]

    def setUpVectors(self, timeStamps, torque, shaftSpeed):
        duration = round(len(timeStamps)/self.fs) # Duration of simulation [seconds]
        self.numTimeSteps = round(duration/self.dt) # Number of time steps in this simulation
        timeVector = np.linspace(0, duration, self.numTimeSteps)
        self.numDegreesOfFreedom = max(np.shape(self.massMatrix))

        # Initialise force vector
        self.forceVector = np.zeros([self.numDegreesOfFreedom+2, self.numTimeSteps])

        # Load and filter measured internal torque
        self.forceVector[-2, :] = smoothN.smoothn(yin=torque[0:self.numTimeSteps], s=200)[0][:,0]*1000 # This doesn't really seem to smooth it out all that much

        # Load and filter measured shaft angular velocity at Q1
        for index, rpm in enumerate(shaftSpeed):
            shaftSpeed[index] = rpm[0]*(2*math.pi/60)

        # THIS STILL NEEDS TO BE SCALED BY 10,000,000 BUT I'm RUNNING OUT OF MEMORY
        self.forceVector[-1, :] = shaftSpeed[0:self.numTimeSteps]

    def solveSystem(self):

        res = self.inverseInitialCond(self.forceVector[-2,1], self.forceVector[-1,1]*(10**7))
        
        print(res[0])
        print(res[1])
        print(res[2])

    def inverseInitialCond(self, initialMeasuredForce, initialMeasuredVel):
        F = np.zeros([self.numDegreesOfFreedom, self.numTimeSteps])

        F[-2, :] = initialMeasuredForce
        F[-1, :] = initialMeasuredVel

        rhoInf = 0

        # Determine alphaM, alphaF, and gamma
        alphaM = (1/2)*((3-rhoInf)/(1+rhoInf))
        alphaF = 1/(1+rhoInf)
        gamma = (1/2) + alphaM - alphaF

        # Initialise outputs
        u = np.zeros([self.numDegreesOfFreedom, self.numTimeSteps])
        ud = np.zeros([self.numDegreesOfFreedom, self.numTimeSteps])
        v = np.zeros([self.numDegreesOfFreedom, self.numTimeSteps])
        vd = np.zeros([self.numDegreesOfFreedom, self.numTimeSteps])

        # Set initial conditions
        u[:,1] = 0
        ud[:,1] = 0
        v[:,1] = 0
        vd[:,1] = 0

        # Constants for effective stiffnedd matrix
        a1 = (alphaM**2)/(alphaF*(gamma**2)*(self.dt**2))
        a2 = alphaM/(gamma*self.dt)

        # Constants for effective force
        a3 = alphaM/(alphaF*gamma*self.dt)
        a4 = (gamma-alphaM)/(gamma*alphaF)
        a5 = (alphaF-1)/alphaF
        a6 = alphaM/(alphaF*(gamma**2)*(self.dt**2))
        a7 = 1/(alphaF*gamma*self.dt)
        a8 = (gamma-1)/gamma
        a9 = (gamma-alphaM)/(alphaF*(gamma**2)*self.dt)
        a10 = 1/(gamma*self.dt)

        # Determine effective stiffness matrix
        effectiveStiffness = a1*self.massMatrix + a2*self.dampingMatrix + alphaF*self.stiffnessMatrix

        for i in range(0, self.numTimeSteps-1):
            # Determine the effective force for this timestep
            currF = F[:,i]
            nextF = F[:,i+1]

            effectiveForce = (alphaF*nextF + (1-alphaF)*currF - (1-alphaM)*self.massMatrix*vd[:,i] - (1-alphaF)*self.dampingMatrix*v[:,i] - (1-alphaF)*self.stiffnessMatrix*u[:,i] + alphaF*self.dampingMatrix*(a3*u[:,i] - a4*ud[:,i] - a5*v[:,i]) + alphaM*self.massMatrix*(a6*u[:,i] + a7*v[:,i] - a8*vd[:,i] - a9*ud[:,i]))[0]


            # Determine the displacement for the next time step
            u[:, i+1] = np.linalg.lstsq(effectiveStiffness, effectiveForce, rcond=-1)[0]            
            ud[:, i+1] = a10*(u[:,i+1] - u[:,i]) + a8*ud[:,i]
            v[:,i+1] = a3*(u[:,i+1] - u[:,i]) + a4*ud[:,i] + a5*v[:,i]
            vd[:,i+1] = a6*(u[:,i+1] - u[:,i]) - a7*v[:,i] + a8*vd[:,i] + a9*ud[:,i]

        DI = u[:,-1]
        VI = ud[:,-1]
        AI = vd[:,-1]

        return DI, VI, AI

    def jwhAlpha(Self):
        pass

mySolver = InverseSolutionSolver()

mySolver.setUpMatrices()

# print("Mass matrix: \n", mySolver.massMatrix)

inputData = loadmat("/home/nic/Documents/Work/Masters/Code/mastersCaseStudy/services/propellerMonitorService/test data/Case_2019_10_30_20_06_20.mat")

mySolver.setUpVectors(inputData['Time_Final'], inputData['ForeTorqFinal'], inputData['rpmFinal'])

mySolver.solveSystem()