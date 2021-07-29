# **Power Train Service**

## **Description**
This service offers information about the power train of the S.A. Agulhas II. The original research producing the models used here was carried out by Gerhard Durandt (https://scholar.sun.ac.za/handle/10019.1/109321). It offers four service calls, each serving different temporal aspects (foresight, insight, and hindsight) and value-spaces (tactical and strategic):
- The 'PowerEstimate' call provides foresight for tactical decision-making (As is described by https://www.researchgate.net/publication/332173693_Designing_Ship_Digital_Services) by providing power estimates for a requested route and sailing conditions.
- The 'CostEstimate' call provides foresight for tactical decision-making (As is described by https://www.researchgate.net/publication/332173693_Designing_Ship_Digital_Services) by providing cost estimates for a requested route and sailing conditions.
- The 'Power Tracking' call provides insight for tactical and operational decision-making (As is described by https://www.researchgate.net/publication/332173693_Designing_Ship_Digital_Services) by providing real-time power use by the vessel.
- The 'Power Estimate Evaluation' call provdes hindsight for strategic decision-making (As is described by https://www.researchgate.net/publication/332173693_Designing_Ship_Digital_Services) by evaluating the accuracy of the models predictions.

For information on the API design (calls and messages), check out the proto file at https://github.com/NicholasBunn/mastersCaseStudy/blob/main/services/powerTrainService/proto/v1/power_train_service_api_v1.proto

## **Prerequisites**
- Python3
- make and/or Docker

## **Installation**
### **Local**
- Open a terminal instance and navigate to the directory in which you'd like to install the project.
- Run "git clone https://github.com/NicholasBunn/mastersCaseStudy.git" to clone the project onto your machine.
- Run "cd MastersCaseStudy" to change directory so that you are situated in "MastersCaseStudy").
- Run "python3 -m pip install -r services/powerTrainService/requirements.txt" to install all the necessary dependencies.
- (Optional) Run "make testPowerTrainService" to test that all dependencies are installed properly and the service has all it needs.
- Run "make runPowerTrainService" to start up the service.

### **Docker**
- Open a terminal instance and navigate to the directory in which you'd like to install the project.
- Run "git clone https://github.com/NicholasBunn/mastersCaseStudy.git" to clone the project onto your machine.
- Run "cd MastersCaseStudy" to change directory so that you are situated in "MastersCaseStudy").
- Run "docker build . -f services/powerTrainService/DockerFile -t power_train_service" to build the Docker image for this service (This could take a while, Tensorflow is a big boi).
- Run "docker run -p 127.0.0.1:50051:50051/tcp power_train_service" to start up a Docker container with the service.

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