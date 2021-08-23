FROM golang:alpine

# Install git.
# Git is required for fetching the dependencies.
RUN apk add --no-cache git

# Create project directory
RUN mkdir /project/
# Create services directory
RUN mkdir /project/services/
# Create a route analysis aggregaotr folder in the services directory
RUN mkdir /project/services/routeAnalysisAggregator/
# Create a program logs folder in the route analysis aggregaotr directory
RUN mkdir /project/services/routeAnalysisAggregator/program\ logs/

# Copy the 'proto' folder into the services directory
COPY services/routeAnalysisAggregator/proto/v1/ /project/services/routeAnalysisAggregator/proto/v1
# Copy the 'certification' folder into the project directory
COPY certification /project/certification
# Copy the service contents into the authenticationService directory
COPY services/routeAnalysisAggregator/configuration.yaml /project/services/routeAnalysisAggregator
COPY services/routeAnalysisAggregator/routeAnalysisAggregator.go /project/services/routeAnalysisAggregator
COPY services/routeAnalysisAggregator/routeAnalysisAggregator_test.go /project/services/routeAnalysisAggregator
COPY services/routeAnalysisAggregator/go.mod /project/services/routeAnalysisAggregator
COPY services/routeAnalysisAggregator/go.sum /project/services/routeAnalysisAggregator

WORKDIR /project/services/routeAnalysisAggregator/

# Fetch the dependecies
RUN go mod tidy

# Build the binary for authentication service
RUN go build -o ./routeAnalysisAggregator .

EXPOSE 50101
ENTRYPOINT ["./routeAnalysisAggregator"]