# **User Database**

## **Description**

This database stores user details (username, hashed password, priviledges). The authentication service uses this to log users in.

For more information on the table structure, check out the table desing at https://github.com/NicholasBunn/mastersCaseStudy/blob/main/services/userDatabase/scripts/createTable.sql

## **Prerequisites**

- Docker

## **Installation**

### **Local**

- This database runs in a Docker container, I'm not going to to add instructions for running local DB instances

### **Docker**

- Open a terminal instance and navigate to the directory in which you'd like to install the project.
- Run "git clone https://github.com/NicholasBunn/mastersCaseStudy.git" to clone the project onto your machine.
- Run "cd MastersCaseStudy" to change directory so that you are situated in "MastersCaseStudy").
- Run "docker build . -f services/userDatabase/Dockerfile -t user_database" to build the Docker image for this service.
- Run "docker run -d -p 127.0.0.1:3306:3306/tcp --name user_database -e MYSQL_ROOT_PASSWORD="supersecret" user_database" to start up a Docker container with the service.

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
