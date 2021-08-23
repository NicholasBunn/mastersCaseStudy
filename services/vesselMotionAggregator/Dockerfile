FROM golang:alpine

# Install git.
# Git is required for fetching the dependencies.
RUN apk add --no-cache git

# Create project directory
RUN mkdir /project/
# Create services directory
RUN mkdir /project/services/
# Create a vessel motion aggregator folder in the services directory
RUN mkdir /project/services/vesselMotionAggregator/
# Create a program logs folder in the vessel motion aggregator directory
RUN mkdir /project/services/vesselMotionAggregator/program\ logs/

# Copy the 'proto' folder into the services directory
COPY services/vesselMotionAggregator/proto/v1/ /project/services/vesselMotionAggregator/proto/v1
# Copy the 'certification' folder into the project directory
COPY certification /project/certification
# Copy the service contents into the authenticationService directory
COPY services/vesselMotionAggregator/configuration.yaml /project/services/vesselMotionAggregator
COPY services/vesselMotionAggregator/vesselMotionAggregator.go /project/services/vesselMotionAggregator
COPY services/vesselMotionAggregator/vesselMotionAggregator_test.go /project/services/vesselMotionAggregator
COPY services/vesselMotionAggregator/go.mod /project/services/vesselMotionAggregator
COPY services/vesselMotionAggregator/go.sum /project/services/vesselMotionAggregator

WORKDIR /project/services/vesselMotionAggregator/

# Fetch the dependecies
RUN go mod tidy

# Build the binary for authentication service
RUN go build -o ./vesselMotionAggregator .

EXPOSE 50103
ENTRYPOINT ["./vesselMotionAggregator"]