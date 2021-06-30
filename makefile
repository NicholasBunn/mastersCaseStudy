# Make commands relating to proto files
protoGenGo:
	# Ocean weather service
	# protoc --go_out=. --go-grpc_out=. apis/ocean_weather_service_api/v1/ocean_weather_service_api_v1.proto

protoGenPy:
	# Ocean weather service
	python3 -m grpc_tools.protoc -I=services/oceanWeatherService/proto/v1 --python_out=services/oceanWeatherService/proto/v1/generated  --grpc_python_out=services/oceanWeatherService/proto/v1/generated services/oceanWeatherService/proto/v1/ocean_weather_service_api_v1.proto # add 'from .' to line 5 of the _grpc.py file

	# Power train service
	python3 -m grpc_tools.protoc -I=services/powerTrainService/proto/v1 --python_out=services/powerTrainService/proto/v1/generated  --grpc_python_out=services/powerTrainService/proto/v1/generated services/powerTrainService/proto/v1/power_train_service_api_v1.proto # add 'from .' to line 5 of the _grpc.py file

protoGenCSharp:
	# protoc -I=services/vesselMotionService/proto/v1 services/vesselMotionService/proto/v1/vessel_motion_service_api_v1.proto --csharp_out=services/vesselMotionService/proto/v1/generated # --grpc_out=services/vesselMotionService/proto/v1/generated --plugin=protoc-gen-grpc=tools\grpc_csharp_plugin.exe
	# NEED TO ADD A COMMAND HERE TO CREATE CLIENT STUBS FOR AGGREGATORS
	
protoGenCPP:
	protoc -I=services/comfortService/proto/v1 --grpc_out=services/comfortService/proto/v1/generated --plugin=protoc-gen-grpc=/usr/local/bin/grpc_cpp_plugin services/comfortService/proto/v1/comfort_service_api_v1.proto
	protoc -I=services/comfortService/proto/v1 --cpp_out=services/comfortService/proto/v1/generated services/comfortService/proto/v1/comfort_service_api_v1.proto

protoGenAll:
	make protoGenGo; make protoGenPy; protoGenCSharp; protoGenCPP

protoCleanGo:
	find . -type f -name '*.pb.go' -delete

protoCleanPy:
	find . -type f -name '*_grpc.py' -delete
	find . -type f -name '*_pb2.py' -delete

protoCleanCPP:
	find . -type f -name '*.pb.cc' -delete
	find . -type f -name '*.pb.h' -delete
	find . -type f -name '*.pb.o' -delete

protoCleanAll:
	make protoCleanGo; make protoCleanPy; make protoCleanCPP

# Make commands related to TLS certificates
certGen:
	cd certification; ./gen.sh; cd ..

certClean:
	cd certification; rm *.pem; rm *.srl; cd ..

# Make commands relating to testing the program
testOceanWeatherService:
	/usr/bin/python3 /home/nic/Documents/Work/Masters/Code/mastersCaseStudy/services/oceanWeatherService/test_oceanWeatherService.py

testPowerTrainService:
	/usr/bin/python3 /home/nic/Documents/Work/Masters/Code/mastersCaseStudy/services/powerTrainService/test_powerTrainService.py

testVesselMotionService:
	cd services/vesselMotionService; dotnet test; cd ..

testGo:
	go test ./...

testPy:
	make testOceanWeatherService
	make testPowerTrainService

testCSharp:
	make testVesselMotionService

testAll:
	make testPy
	make testCSharp

# Make commands relating to running the program
runOceanWeatherService:
	/usr/bin/python3 /home/nic/Documents/Work/Masters/Code/mastersCaseStudy/services/oceanWeatherService/oceanWeatherService.py
	
runPowerTrainService:
	/usr/bin/python3 /home/nic/Documents/Work/Masters/Code/mastersCaseStudy/services/powerTrainService/powerTrainService.py

runVesselMotionService:
	cd services/vesselMotionService; dotnet run; cd ..

runComfortService:
	cd services/comfortService; make build; ./comfortService

runPrometheus:

runPushGateway:

runPy:
	make runOceanWeatherService
	make runPowerTrainService

runAll:
	make runPy