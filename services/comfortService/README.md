# **Process Vibration Service**

## **Description**

This service quantifies human comfort on maritime vessels according to SANS 2631-1. It offers a single service call.

- The 'Comfort Rating' call provides foresight for tactical decision-making by providing a comfort rating for a proposed route, based on estimated vibrations on board.

For information on the API design (calls and messages), check out the proto file at https://github.com/NicholasBunn/mastersCaseStudy/blob/main/services/comfortService/proto/v1/comfort_service_api_v1.proto

## **Prerequisites**

- Python3
- make and/or Docker

## **Installation**

### **Local**

- Open a terminal instance and navigate to the directory in which you'd like to install the project.
- Run "git clone https://github.com/NicholasBunn/mastersCaseStudy.git" to clone the project onto your machine.
- Run "cd MastersCaseStudy" to change directory so that you are situated in "MastersCaseStudy").
- Run "python3 -m pip install -r services/comfortService/requirements.txt" to install all the necessary dependencies.
- (Optional) Run "make testComfortService" to test that all dependencies are installed properly and the service has all it needs.
- Run "make runComfortService" to start up the service.

### **Docker**

- Open a terminal instance and navigate to the directory in which you'd like to install the project.
- Run "git clone https://github.com/NicholasBunn/mastersCaseStudy.git" to clone the project onto your machine.
- Run "cd MastersCaseStudy" to change directory so that you are situated in "MastersCaseStudy").
- Run "docker build . -f services/comfortService/Dockerfile -t comfort_service" to build the Docker image for this service.
- Run "docker run -p 127.0.0.1:50053:50053/tcp comfort_service" to start up a Docker container with the service.

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
