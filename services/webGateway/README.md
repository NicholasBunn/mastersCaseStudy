# **Web Gateway**

## **Description**

This service acts as a reception area to access the system from the browser. Web requests, routed through the Envoy proxy are received and handled by this component. It handles user logins, returning authentication tokens, as well as routing user requests to the relevant backend components.

- The 'LoginService' service call queries an authentication service to log users in and provide them with their authentication token, which is required to access backend services.
- The 'RouteAnalysisAggregator' service call routes user requests to the Route Analysis Aggregator.
- The 'RoutePowerAggregator' service call rotues user requests to the Power Train Aggregator.
- The 'RouteMotionAggregator' service call routes user requests to the Vessel Motion Aggregator.

For information on the API design (calls and messages), check out the proto file at https://github.com/NicholasBunn/mastersCaseStudy/blob/main/services/webGateway/proto/v1/web_gateway_api_v1.proto

## **Prerequisites**

- Go
- make and/or Docker

## **Installation**

### **Local**

- Open a terminal instance and navigate to the directory in which you'd like to install the project.
- Run "git clone https://github.com/NicholasBunn/mastersCaseStudy.git" to clone the project onto your machine.
- Run "cd MastersCaseStudy" to change directory so that you are situated in "MastersCaseStudy".
- Run "cd services/webGateway" to change into this service's directory.
- Run "go mod tidy" to fetch and tidy up the required packages.
- Run "cd .." twice to jump back into the "MastersCaseStudy" directory.
- (Optional) Run "make testWebGateway" to test that all dependencies are installed properly and the service is running as expected. This will fail if you're running this before I've added mock tests.
- Run "make runWebGateway" to start up the service.

### **Docker**

- Open a terminal instance and navigate to the directory in which you'd like to install the project.
- Run "git clone https://github.com/NicholasBunn/mastersCaseStudy.git" to clone the project onto your machine.
- Run "cd MastersCaseStudy" to change directory so that you are situated in "MastersCaseStudy").
- Run "docker build . -f services/webGateway/Dockerfile -t web_gateway" to build the Docker image for this service.
- Run "docker run -p 50000:50000 --network=host web_gateway" to start up a Docker container with the service.

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
