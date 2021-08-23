FROM golang:alpine

# Install git.
# Git is required for fetching the dependencies.
RUN apk add --no-cache git

# Create project directory
RUN mkdir /project/
# Create services directory
RUN mkdir /project/services/
# Create an authentication service folder in the services directory
RUN mkdir /project/services/authenticationService/
# Create an program logs folder in the authentication service directory
RUN mkdir /project/services/authenticationService/program\ logs/

# Copy the 'proto' folder into the services directory
COPY services/authenticationService/proto/v1/ /project/services/authenticationService/proto/v1
# Copy the 'certification' folder into the project directory
COPY certification /project/certification
# Copy the service contents into the authenticationService directory
COPY services/authenticationService/configuration.yaml /project/services/authenticationService
COPY services/authenticationService/authenticationService.go /project/services/authenticationService
COPY services/authenticationService/authenticationService_test.go /project/services/authenticationService
COPY services/authenticationService/go.mod /project/services/authenticationService
COPY services/authenticationService/go.sum /project/services/authenticationService

WORKDIR /project/services/authenticationService/

# Fetch the dependecies
RUN go mod tidy

# Build the binary for authentication service
RUN go build -o ./authenticationService .

EXPOSE 50101
ENTRYPOINT ["./authenticationService"]