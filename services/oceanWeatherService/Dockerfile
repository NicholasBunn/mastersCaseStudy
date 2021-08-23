FROM python

# Create project directory
RUN mkdir /project
# Create services directory
RUN mkdir /project/services
# Create an ocean weather service folder in the services directory
RUN mkdir /project/services/oceanWeatherService/
# Create an program logs folder in the ocean weather service directory
RUN mkdir /project/services/oceanWeatherService/program\ logs

# Copy the 'proto' folder into the services directory
COPY protoFiles/python/oceanWeatherService/v1/ project/protoFiles/python/oceanWeatherService/v1
# Copy the 'interceptor' folder into the project directory
COPY interceptors/python/ project/interceptors/python/
# Copy the 'certification' folder into the project directory
COPY certification/ /project/certification/
# Copy the service contents into the oceanWeatherService directory
COPY services/oceanWeatherService/configuration.yaml project/services/oceanWeatherService
COPY services/oceanWeatherService/requirements.txt project/services/oceanWeatherService
COPY services/oceanWeatherService/test_oceanWeatherService.py project/services/oceanWeatherService
COPY services/oceanWeatherService/oceanWeatherService.py project/services/oceanWeatherService

# Basically performing cd inside the image and setting the work directory
WORKDIR /project

# Update pip to avoid warnings
RUN python3 -m pip install --upgrade pip
# Tells docker to run pip install inside the image which will add
# all the required packages
RUN python3 -m pip install -r services/oceanWeatherService/requirements.txt

# Run tests on the service before starting it up
RUN python3 services/oceanWeatherService/test_oceanWeatherService.py

# Tells Docker that we're going to run the microservice on port 50050
EXPOSE 50050
# Tells Docker how to run the microservice
ENTRYPOINT [ "python3", "services/oceanWeatherService/oceanWeatherService.py" ]
