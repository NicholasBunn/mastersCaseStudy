FROM golang:alpine

# Install git.
# Git is required for fetching the dependencies.
RUN apk add --no-cache git

# Create project directory
RUN mkdir /project/
# Create services directory
RUN mkdir /project/services/
# Create a web gateway folder in the services directory
RUN mkdir /project/services/webGateway/
# Create a program logs folder in the web gateway directory
RUN mkdir /project/services/webGateway/program\ logs/

# Copy the 'proto' folder into the services directory
COPY services/webGateway/proto/v1/ /project/services/webGateway/proto/v1
# Copy the 'certification' folder into the project directory
COPY certification /project/certification
# Copy the service contents into the webGateway directory
COPY services/webGateway/configuration.yaml /project/services/webGateway
COPY services/webGateway/webGateway.go /project/services/webGateway
COPY services/webGateway/webGateway_test.go /project/services/webGateway
COPY services/webGateway/go.mod /project/services/webGateway
COPY services/webGateway/go.sum /project/services/webGateway

WORKDIR /project/services/webGateway/

# Fetch the dependecies
RUN go mod tidy

# Build the binary for the web gateway
RUN go build -o ./webGateway .

EXPOSE 50000
ENTRYPOINT ["./webGateway"]