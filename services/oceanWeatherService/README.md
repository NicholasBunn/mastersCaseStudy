# **Ocean Weather Service**

## **Description**
This service offers information about marine weather conditions. It offers two service calls, each serving different temporal aspects (foresight and hindsight) and value-spaces (tactical and strategic):
- The 'OceanWeatherPrediction' service call provides foresight for tactical decision-making (As is described by https://www.researchgate.net/publication/332173693_Designing_Ship_Digital_Services) by providing future ocean weather conditions along a requested route.
- The 'OceanWeatherHistory' call provides hindsight for stategic decision-making (As is described by https://www.researchgate.net/publication/332173693_Designing_Ship_Digital_Services) by providing historical ocean weather conditions that the ship would have encountered along a requested route.

For information on the API design (calls and messages), check out the proto file at https://github.com/NicholasBunn/mastersCaseStudy/blob/main/services/oceanWeatherService/proto/v1/ocean_weather_service_api_v1.proto

## **Prerequisites**

## **Installation**
### **Local**
- Open a terminal instance and navigate to the directory in which you'd like to install the project.
- Run "git clone https://github.com/NicholasBunn/mastersCaseStudy.git" to clone the project onto your machine.
- Run "cd MastersCaseStudy" to change directory so that you are situated in "MastersCaseStudy").
- Run "python3 -m pip install -r services/oceanWeatherService/requirements.txt" to install all the necessary dependencies.
- (Optional) Run "make testOceanWeatherService" to test that all dependencies are installed properly and the service has all it needs.
- Run "make runOceanWeatherService" to start up the service.

### **Docker##
- Open a terminal instance and navigate to the directory in which you'd like to install the project.
- Run "git clone https://github.com/NicholasBunn/mastersCaseStudy.git" to clone the project onto your machine.
- Run "cd MastersCaseStudy" to change directory so that you are situated in "MastersCaseStudy").
- Run "docker build . -f services/oceanWeatherService/DockerFile -t ocean_weather_service" to build the Docker image for this service.
- Run "docker run -p 127.0.0.1:50050:50050/tcp ocean_weather_service" to start up a Docker container with the service.


## **Contributing**
This project forms part of my masters thesis, and as such collaboration is not currently offered. If you'd like to build on this project in your research, or get a bit more information on things feel free to drop me or either of my supervisors a message. You can find my supervisor's contact details at:
- https://sites.google.com/view/mad-research-group
- https://svrg.sun.ac.za/

## **License**
I'm still investigating what licenses this project falls under. I suppose it's research output so it's governed by Stellenbosch University, but I've thrown it up here to help anyone who may be building on my work in future projects.

## **Citation**
This software can be referenced by citing my masters thesis (to be published). Please keep an eye on my Research Gate profile (https://www.researchgate.net/profile/Nicholas-Bunn-2) for it's release.

## **Contact**
Drop me a message on either of the following:
- Email: nicholasbunn04@gmail.com
- LinkedIn: https://www.linkedin.com/in/nicholasbunn/