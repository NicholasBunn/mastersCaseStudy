FROM python

# Create project directory
RUN mkdir /project
# Create services directory
RUN mkdir /project/services
# Create an ocean weather service folder in the services directory
RUN mkdir /project/services/propellerMonitorService/
# Create an program logs folder in the propeller monitor service directory
RUN mkdir /project/services/propellerMonitorService/program\ logs

# Create variable for required files and test data (This is just because there are spaces in the file paths)
ENV TEST_DATA_PATH "services/propellerMonitorService/test data/"

# Copy the 'proto' folder into the services directory
COPY protoFiles/python/propellerMonitorService/v1/ project/protoFiles/python/propellerMonitorService/v1
# Copy the 'interceptor' folder into the project directory
COPY interceptors/python/ project/interceptors/python/
# Copy the 'certification' folder into the project directory
COPY certification/ /project/certification/
# Copy the test data into the propellerMonitorService directory
COPY ${TEST_DATA_PATH} project/${TEST_DATA_PATH}
# Copy the service contents into the propellerMonitorService directory
COPY services/propellerMonitorService/configuration.yaml project/services/propellerMonitorService
COPY services/propellerMonitorService/requirements.txt project/services/propellerMonitorService
COPY services/propellerMonitorService/smoothN.py project/services/propellerMonitorService
COPY services/propellerMonitorService/test_propellerMonitorService.py project/services/propellerMonitorService
COPY services/propellerMonitorService/propellerMonitorService.py project/services/propellerMonitorService

# Basically performing cd inside the image and setting the work directory
WORKDIR /project

# Update pip to avoid warnings
RUN python3 -m pip install --upgrade pip
# Tells docker to run pip install inside the image which will add
# all the required packages
RUN python3 -m pip install -r services/propellerMonitorService/requirements.txt

# Run tests on the service before starting it up
# RUN python3 services/propellerMonitorService/test_propellerMonitorService.py

# Tells Docker that we're going to run the microservice on port 50050
EXPOSE 50055
# Tells Docker how to run the microservice
ENTRYPOINT [ "python3", "services/propellerMonitorService/propellerMonitorService.py" ]
