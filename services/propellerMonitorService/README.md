# **Propeller Monitor Service**

## **Description**

This service offers information about the propeller of the S.A. Agulhas II. This is a servitisation of the algorithm written by Brendon Nickerson for his doctoral thesis (https://scholar.sun.ac.za/handle/10019.1/110071). This is a typical example of a digital twin offered service, where any information requried by the service can be provided by the twin itself, with the information provided describing the real asset. It offers one service call, serving insight for - The 'EstimatePropellerLoad' service call provides insight for operational decision-making (As is described by https://www.researchgate.net/publication/332173693_Designing_Ship_Digital_Services) by providing an estimation of the ice-induced load on the propeller during ice passage of the S.A. Agulhas II.
decision making:

- The 'EstimatePropellerLoad' service call provides insight for operational decision-making (As is described by https://www.researchgate.net/publication/332173693_Designing_Ship_Digital_Services) by providing an estimation of the ice-induced load on the propeller during ice passage of the S.A. Agulhas II.

For information on the API design (calls and messages), check out the proto file at https://github.com/NicholasBunn/mastersCaseStudy/blob/main/services/propellerMonitorService/proto/v1/propeller_monitor_service_api_v1.proto

## **Prerequisites**

- Python3
- make and/or Docker

## **Installation**

### **Local**

- Open a terminal instance and navigate to the directory in which you'd like to install the project.
- Run "git clone https://github.com/NicholasBunn/mastersCaseStudy.git" to clone the project onto your machine.
- Run "cd MastersCaseStudy" to change directory so that you are situated in "MastersCaseStudy").
- Run "python3 -m pip install -r services/propellerMonitorService/requirements.txt" to install all the necessary dependencies.
- (Optional) Run "make test_propellerMonitorService.py" to test that all dependencies are installed properly and the service has all it needs.
- Run "make runPropellerMonitorService" to start up the service.

### **Docker**

- Open a terminal instance and navigate to the directory in which you'd like to install the project.
- Run "git clone https://github.com/NicholasBunn/mastersCaseStudy.git" to clone the project onto your machine.
- Run "cd MastersCaseStudy" to change directory so that you are situated in "MastersCaseStudy").
- Run "docker build . -f services/propellerMonitorService/Dockerfile -t propeller_monitor_service" to build the Docker image for this service.
- Run "docker run -e PUSHGATEWAYHOST: custom address -p 127.0.0.1:50055:50055/tcp propeller_monitor_service" to start up a Docker container with the service.

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
