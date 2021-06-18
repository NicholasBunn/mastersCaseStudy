# Make commands relating to proto files
protoGenGo:
	# Ocean weather service
	# protoc --go_out=. --go-grpc_out=. apis/ocean_weather_service_api/v1/ocean_weather_service_api_v1.proto

protoGenPy:
	# Ocean weather service
	python3 -m grpc_tools.protoc -I=services/oceanWeatherService/proto/v1 --python_out=services/oceanWeatherService/proto/v1/generated  --grpc_python_out=services/oceanWeatherService/proto/v1/generated services/oceanWeatherService/proto/v1/ocean_weather_service_api_v1.proto # add 'from .' to line 5 of the _grpc.py file

protoGenAll:
	make protoGenGo; make protoGenPy

protoCleanGo:
	find . -type f -name '*.pb.go' -delete

protoCleanPy:
	find . -type f -name '*_grpc.py' -delete
	find . -type f -name '*_pb2.py' -delete

protoCleanAll:
	make protoCleanGo; make protoCleanPy

# Make commands related to TLS certificates
certGen:
	cd certification; ./gen.sh; cd ..


certClean:
	cd certification; rm *.pem; rm *.srl; cd ..

# Make commands relating to testing the program
testOceanWeatherService:
	/usr/bin/python3 /home/nic/Documents/Work/Masters/Code/mastersCaseStudy/services/oceanWeatherService/test_oceanWeatherService.py

testGo:
	go test ./...

testPy:
	make testOceanWeatherService

testAll:
	make testPy

# Make commands relating to running the program
runOceanWeatherService:
	/usr/bin/python3 /home/nic/Documents/Work/Masters/Code/mastersCaseStudy/services/oceanWeatherService/oceanWeatherService.py
	
runPrometheus:

runPushGateway:

runPy:
	make runOceanWeatherService

runAll:
	make runPy