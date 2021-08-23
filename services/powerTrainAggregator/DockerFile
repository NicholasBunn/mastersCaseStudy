FROM golang:alpine

# Install git.
# Git is required for fetching the dependencies.
RUN apk add --no-cache git

# Create project directory
RUN mkdir /project/
# Create services directory
RUN mkdir /project/services/
# Create a power train aggregator folder in the services directory
RUN mkdir /project/services/powerTrainAggregator/
# Create a program logs folder in the power train aggregator directory
RUN mkdir /project/services/powerTrainAggregator/program\ logs/

# Copy the 'proto' folder into the services directory
COPY services/powerTrainAggregator/proto/v1/ /project/services/powerTrainAggregator/proto/v1
# Copy the 'certification' folder into the project directory
COPY certification /project/certification
# Copy the service contents into the authenticationService directory
COPY services/powerTrainAggregator/configuration.yaml /project/services/powerTrainAggregator
COPY services/powerTrainAggregator/powerTrainAggregator.go /project/services/powerTrainAggregator
COPY services/powerTrainAggregator/powerTrainAggregator_test.go /project/services/powerTrainAggregator
COPY services/powerTrainAggregator/go.mod /project/services/powerTrainAggregator
COPY services/powerTrainAggregator/go.sum /project/services/powerTrainAggregator

WORKDIR /project/services/powerTrainAggregator/

# Fetch the dependecies
RUN go mod tidy

# Build the binary for authentication service
RUN go build -o ./powerTrainAggregator .

EXPOSE 50102
ENTRYPOINT ["./powerTrainAggregator"]