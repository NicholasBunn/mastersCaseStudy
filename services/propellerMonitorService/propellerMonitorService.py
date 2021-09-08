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
            for j in range(1, 102, 2):
                self.massMatrix[i,j] = self.massMatrix[i,j] + self.Jp + self.Jm + self.Js*math.cos((i)*math.pi*21.525/self.L)*math.cos((j)*math.pi*21.525/self.L)

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
        self.forceVector = np.zeros([self.numDegreesOfFreedom, self.numTimeSteps])

        # Load and filter measured internal torque
        self.forceVector[-2, :] = smoothN.smoothn(yin=torque[0:self.numTimeSteps], s=200)[0][:,0]*1000 # This doesn't really seem to smooth it out all that much

        # Load and filter measured shaft angular velocity at Q1
        for index, rpm in enumerate(shaftSpeed):
            shaftSpeed[index] = rpm[0]*(2*math.pi/60)

        # THIS STILL NEEDS TO BE SCALED BY 10,000,000 BUT I'm RUNNING OUT OF MEMORY
        self.forceVector[-1,:] = shaftSpeed[0:self.numTimeSteps]
        self.forceVector[-1,:] = self.forceVector[-1,:]*(10**7) # Could move this line above

    def solveSystem(self):

        DI, VI, AI = self.inverseInitialCond(self.forceVector[-2,1], self.forceVector[-1,1])
        
        # print(DI)
        # print(VI)
        # print(AI)

        self.forceVector[-1,:] = self.forceVector[-1,:]*(10**7) # Could move this line above

        q, qd, qdVer, qdd = self.jwhAlpha(DI, VI, AI, self.forceVector)

        print(q[:4,:4])
        print(q[-4:,:4])
        print(q[:4,-4:])
        print(q[-4:,-4:])
        # print(qd)
        # print(qdVer) # Verification matrix
        # print(qdd)

    def inverseInitialCond(self, initialMeasuredForce, initialMeasuredVel):
        F = np.zeros([self.numDegreesOfFreedom, self.numTimeSteps], dtype=np.float64)

        F[-2, :] = initialMeasuredForce
        F[-1, :] = initialMeasuredVel

        rhoInf = 0

        # Determine self.alphaM, self.alphaF, and self.gamma
        self.alphaM = (1/2)*((3-rhoInf)/(1+rhoInf))
        self.alphaF = 1/(1+rhoInf)
        self.gamma = (1/2) + self.alphaM - self.alphaF

        # Initialise outputs
        u = np.zeros([self.numDegreesOfFreedom, self.numTimeSteps])
        ud = np.zeros([self.numDegreesOfFreedom, self.numTimeSteps])
        v = np.zeros([self.numDegreesOfFreedom, self.numTimeSteps])
        vd = np.zeros([self.numDegreesOfFreedom, self.numTimeSteps])

        # Set initial conditions
        u[:,0] = 0
        ud[:,0] = 0
        v[:,0] = 0
        vd[:,0] = 0

        # Constants for effective stiffnedd matrix
        self.a1 = (self.alphaM**2)/(self.alphaF*(self.gamma**2)*(self.dt**2))
        self.a2 = self.alphaM/(self.gamma*self.dt)

        # Constants for effective force
        self.a3 = self.alphaM/(self.alphaF*self.gamma*self.dt)
        self.a4 = (self.gamma-self.alphaM)/(self.gamma*self.alphaF)
        self.a5 = (self.alphaF-1)/self.alphaF
        self.a6 = self.alphaM/(self.alphaF*(self.gamma**2)*(self.dt**2))
        self.a7 = 1/(self.alphaF*self.gamma*self.dt)
        self.a8 = (self.gamma-1)/self.gamma
        self.a9 = (self.gamma-self.alphaM)/(self.alphaF*(self.gamma**2)*self.dt)
        self.a10 = 1/(self.gamma*self.dt)

        # Determine effective stiffness matrix
        self.effectiveStiffness = self.a1*self.massMatrix + self.a2*self.dampingMatrix + self.alphaF*self.stiffnessMatrix

        # effectiveForce = np.ndarray(54, dtype=np.float64)
        for i in range(0, self.numTimeSteps-1):
            # Determine the effective force for this timestep
            effectiveForce = (
                self.alphaF*F[:,i+1] + (1-self.alphaF)*F[:,i] - \
                (1-self.alphaM)*np.dot(self.massMatrix, vd[:,i]) - (1-self.alphaF)*np.dot(self.dampingMatrix,v[:,i]) - (1-self.alphaF)*np.dot(self.stiffnessMatrix,u[:,i]) + \
                self.alphaF*(np.dot(self.dampingMatrix, self.a3*u[:,i]) - np.dot(self.dampingMatrix, self.a4*ud[:,i]) - np.dot(self.dampingMatrix, self.a5*v[:,i])) + \
                self.alphaM*(np.dot(self.massMatrix, self.a6*u[:,i]) + np.dot(self.massMatrix, self.a7*v[:,i]) - np.dot(self.massMatrix, self.a8*vd[:,i]) - np.dot(self.massMatrix, self.a9*ud[:,i]))
                )

            # Determine the displacement for the next time step
            u[:, i+1] = np.linalg.lstsq(self.effectiveStiffness, effectiveForce, rcond=-1)[0]    

            # Determine the velocity and acceleration for the next time step        
            ud[:, i+1] = self.a10*u[:,i+1] - self.a10*u[:,i] + self.a8*ud[:,i]
            v[:,i+1] = self.a3*u[:,i+1] - self.a3*u[:,i] + self.a4*ud[:,i] + self.a5*v[:,i]
            vd[:,i+1] = self.a6*u[:,i+1] - self.a6*u[:,i] - self.a7*v[:,i] + self.a8*vd[:,i] + self.a9*ud[:,i]

        DI = u[:,-1]
        VI = ud[:,-1]
        AI = vd[:,-1]

        return DI, VI, AI

    def jwhAlpha(self, DI, VI, AI, F):

        # Initialise outputs
        u = np.zeros([self.numDegreesOfFreedom, self.numTimeSteps])
        ud = np.zeros([self.numDegreesOfFreedom, self.numTimeSteps])
        v = np.zeros([self.numDegreesOfFreedom, self.numTimeSteps])
        vd = np.zeros([self.numDegreesOfFreedom, self.numTimeSteps])

        # Set initial conditions
        u[:,0] = DI
        ud[:,0] = VI
        v[:,0] = VI
        vd[:,0] = AI

        for i in range(0, self.numTimeSteps-1, 1):
            
            # Determine the effective force for this timestep
            effectiveForce = (
                self.alphaF*F[:,i+1] + (1-self.alphaF)*F[:,i] - \
                (1-self.alphaM)*np.dot(self.massMatrix, vd[:,i]) - (1-self.alphaF)*np.dot(self.dampingMatrix,v[:,i]) - (1-self.alphaF)*np.dot(self.stiffnessMatrix,u[:,i]) + \
                self.alphaF*(np.dot(self.dampingMatrix, self.a3*u[:,i]) - np.dot(self.dampingMatrix, self.a4*ud[:,i]) - np.dot(self.dampingMatrix, self.a5*v[:,i])) + \
                self.alphaM*(np.dot(self.massMatrix, self.a6*u[:,i]) + np.dot(self.massMatrix, self.a7*v[:,i]) - np.dot(self.massMatrix, self.a8*vd[:,i]) - np.dot(self.massMatrix, self.a9*ud[:,i]))
                )

            # Determine the displacement for the next time step
            u[:, i+1] = np.linalg.lstsq(self.effectiveStiffness, effectiveForce, rcond=-1)[0]    

            # Determine the velocity and acceleration for the next time step        
            ud[:, i+1] = self.a10*(u[:,i+1] - u[:,i]) + self.a8*ud[:,i]
            v[:,i+1] = self.a3*(u[:,i+1] - u[:,i]) + self.a4*ud[:,i] + self.a5*v[:,i]
            vd[:,i+1] = self.a6*(u[:,i+1] - u[:,i]) - self.a7*v[:,i] + self.a8*vd[:,i] + self.a9*ud[:,i]

        return u, ud, v, vd

    def calculateInternalTorqueAndThetas(self, q, qd, qdd):
        # Shaft angular displacement at torque measurement location
        thetaX = q[:,0]
        for i in range(0, 51, 1):
            thetaX = thetaX + np.dot( math.cos((2*(i+1)-1)*math.pi*self.xt/self.L), q[:,i+1] )

        thetaX = thetaX - thetaX[0]

        # Shaft angular velocity at torque measurement location
        thetaXD = qd[:,0]
        for i in range(0, 51, 1):
            thetaXD = thetaXD + np.dot( math.cos((2*(i+1)-1)*math.pi*self.xt/self.L), qd[:,i+1] )

        # Shaft angular velocity at torque measurement location
        thetaXDD = qdd[:,0]
        for i in range(0, 51, 1):
            thetaXDD = thetaXDD + np.dot( math.cos((2*(i+1)-1)*math.pi*self.xt/self.L), qdd[:,i+1] )

        torqueX = np.zeros(self.numTimeSteps)
        for i in range(0, 51, 1):
            torqueX = torqueX - (
                np.dot(
                    (self.G*self.J*math.sin((2*(i+1)-1)*math.pi*self.xt/self.L)*(2*(i+1)-1)*math.pi/self.L), q[:,i+1]
                    )
                )
    
        return thetaX, thetaXD, thetaXDD, torqueX
    
    def calculateAngularVelAndAccel(self, qd, qdd):
        # Shaft angular velocity at propeller
        thetaX0D = qd[:,0]
        for i in range(0, 51, 1):
            thetaX0D = thetaX0D + np.dot( math.cos((2*(i+1)-1)*math.pi*0/self.L), qd[:,i+1] )

        # Shaft angular acceleration at propeller
        thetaX0DD = qdd[:,0]
        for i in range(0, 51, 1):
            thetaX0DD = thetaX0DD + np.dot( math.cos((2*(i+1)-1)*math.pi*0/self.L), qdd[:,i+1] )

        return thetaX0D, thetaX0DD
    
    def extractIceLoad(self, q, thetaX0D):
        Qice = q[:, -2]*(10**9)
        Qmotor = q[:,-1]*(10**9)
        Qprop = (thetaX0D*self.Cp/1000) + (Qice/1000)

        return Qice, Qmotor, Qprop

if __name__=='__main__':
    mySolver = InverseSolutionSolver()

    mySolver.setUpMatrices()

    # print("Mass matrix: \n", mySolver.massMatrix)

    inputData = loadmat("services/propellerMonitorService/test data/Case_2019_10_30_20_06_20.mat")

    mySolver.setUpVectors(inputData['Time_Final'], inputData['ForeTorqFinal'], inputData['rpmFinal'])

    mySolver.solveSystem()