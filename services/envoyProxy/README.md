# **Enovy Proxy**

## **Description**

The envoy proxy is used here to enable a webpage frontend to integrate with the system. Frontends, using grpc-web, can communicate with this proxy on port 8080. The proxy then "translates" the HTTP request sent by the frontend into an HTTP2 request, as used by the gRPC services, and routes the requests to the web gateway on port 50000.

For information on the configuration and setup used in this proxy, check out the yaml file at https://github.com/NicholasBunn/mastersCaseStudy/blob/main/services/envoyProxy/envoy.yaml

## **Prerequisites**

- Docker

## **Installation**

### **Docker**

- Open a terminal instance and navigate to the directory in which you'd like to install the project.
- Run "git clone https://github.com/NicholasBunn/mastersCaseStudy.git" to clone the project onto your machine.
- Run "cd MastersCaseStudy" to change directory so that you are situated in "MastersCaseStudy").
- Run "docker build . -f services/envoyProxy/Dockerfile -t web_proxy" to build the Docker image for envoy.
- Run "docker run -p 127.0.0.1:8080:8080/tcp web_proxy" to start up a Docker container with the service.
  docker run -d -p 8080:8080 -p 9901:9901 --network=host web_proxy

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
