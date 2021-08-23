FROM python:3.8.5

# Create project directory
RUN mkdir /project
# Create services directory
RUN mkdir /project/services
# Create an ocean weather service folder in the services directory
RUN mkdir /project/services/comfortService/
# Create an program logs folder in the ocean weather service directory
RUN mkdir /project/services/comfortService/program\ logs

# Copy the 'proto' folder into the services directory
COPY protoFiles/python/comfortService/v1/ project/protoFiles/python/comfortService/v1
# Copy the 'interceptor' folder into the project directory
COPY interceptors/python project/interceptors/python
# Copy the 'certification' folder into the project directory
COPY certification /project/certification
# Copy the service contents into the comfortService directory
COPY services/comfortService/configuration.yaml project/services/comfortService
COPY services/comfortService/requirements.txt project/services/comfortService
COPY services/comfortService/test_comfortService.py project/services/comfortService
COPY services/comfortService/comfortService.py project/services/comfortService

# Basically performing cd inside the image and setting the work directory
WORKDIR /project

# Update pip to avoid warnings
RUN python3 -m pip install --upgrade pip
# Tells docker to run pip install inside the image which will add
# all the required packages
RUN python3 -m pip install -r services/comfortService/requirements.txt

# Run tests on the service before starting it up
RUN python3 services/comfortService/test_comfortService.py

# Tells Docker that we're going to run the microservice on port 50051
EXPOSE 50053
# Tells Docker how to run the microservice
ENTRYPOINT [ "python3", "services/comfortService/comfortService.py" ]
