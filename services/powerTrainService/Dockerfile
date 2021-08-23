FROM python:3.8.5

# Create project directory
RUN mkdir /project
# Create services directory
RUN mkdir /project/services
# Create an ocean weather service folder in the services directory
RUN mkdir /project/services/powerTrainService/
# Create an program logs folder in the ocean weather service directory
RUN mkdir /project/services/powerTrainService/program\ logs

# Create variable for required files and test data (This is just because there are spaces in the file paths)
ENV REQUIRED_FILE_PATH "services/powerTrainService/required files/"
ENV TEST_DATA_PATH "services/powerTrainService/test data/"

# Copy the 'proto' folder into the services directory
COPY protoFiles/python/powerTrainService/v1/ project/protoFiles/python/powerTrainService/v1
# Copy the 'interceptor' folder into the project directory
COPY interceptors/python project/interceptors/python
# Copy the 'certification' folder into the project directory
COPY certification /project/certification
# Copy the required files into the powerTrainService directory
COPY ${REQUIRED_FILE_PATH} project/${REQUIRED_FILE_PATH}
# Copy the test data into the powerTrainService directory
COPY ${TEST_DATA_PATH} project/${TEST_DATA_PATH}
# Copy the service contents into the powerTrainService directory
COPY services/powerTrainService/configuration.yaml project/services/powerTrainService
COPY services/powerTrainService/requirements.txt project/services/powerTrainService
COPY services/powerTrainService/test_powerTrainService.py project/services/powerTrainService
COPY services/powerTrainService/powerTrainService.py project/services/powerTrainService

# Basically performing cd inside the image and setting the work directory
WORKDIR /project

# Update pip to avoid warnings
RUN python3 -m pip install --upgrade pip
# Tells docker to run pip install inside the image which will add
# all the required packages
RUN python3 -m pip install -r services/powerTrainService/requirements.txt

# Run tests on the service before starting it up
RUN python3 services/powerTrainService/test_powerTrainService.py

# Tells Docker that we're going to run the microservice on port 50051
EXPOSE 50051
# Tells Docker how to run the microservice
ENTRYPOINT [ "python3", "services/powerTrainService/powerTrainService.py" ]
