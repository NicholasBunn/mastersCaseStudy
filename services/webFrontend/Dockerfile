FROM node:14

# Create project directory
RUN mkdir /project/
# Create services directory
RUN mkdir /project/services/
# Create an web frontend folder in the services directory
RUN mkdir /project/services/webFrontend

WORKDIR /project/services/webFrontend

COPY ./services/webFrontend/package*.json ./
RUN npm install

# COPY ./services/webFrontend/dist/main.js ./dist/
# COPY ./services/webFrontend/public ./public
# COPY ./services/webFrontend/app.js ./
# COPY ./services/webFrontend/index.html ./
# COPY ./services/webFrontend/index.js ./
COPY ./services/webFrontend/ ./
# Copy the 'certification' folder into the project directory
COPY certification /project/certification

RUN ls

EXPOSE 8001
CMD [ "node", "app.js"]

#docker build . -f services/webFrontend/Dockerfile -t web_frontend
#docker run --network=host web_frontend