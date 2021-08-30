# **Masters Case Study**

## **Description**

This is a repository for the code relating to my masters thesis case study. This case study was used to verify the deisgn of my architecture, which has been designed to aggregate information from various digital twins and digital services on maritime vessels, each developed by isolated domain experts. This case study is carried out on the S.A. Agulhas II, South Africa's Polar Research and Resupply Vessel.

## **Prerequisites**

If you would like to run project on your local machine, you will need Go, Python3, C# (using .NET 5.0), and prometheus. In addition to this, you will require the packages required by each service (this should be documented by the respective services).

If you don't already have all of this setup, I recommend installing Docker and running each service in it's own container - as they have been designed. Just follow the Installation guide provided with each service.

## **Installation**

### **Local**

- Open a terminal instance and navigate to the directory in which you'd like to install the project.
- Run "git clone https://github.com/NicholasBunn/mastersCaseStudy.git" to clone the project onto your machine.
- Run "cd MastersCaseStudy" to change directory so that you are situated in "MastersCaseStudy".
- Go through each service's README and follow the local installation instructions from this step onwards.

### **Docker**

- Open a terminal instance and navigate to the directory in which you'd like to install the project.
- Run "git clone https://github.com/NicholasBunn/mastersCaseStudy.git" to clone the project onto your machine.
- Run "cd MastersCaseStudy" to change directory so that you are situated in "MastersCaseStudy").
- Run "docker-compose build" to build the Docker images for each service.
- Run "docker-compose run" to start up all services in the Docker network.

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
