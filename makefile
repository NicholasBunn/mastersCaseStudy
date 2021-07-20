# Make commands relating to proto files
protoGenGo:
	# Route analysis aggregator
	protoc -I services --go_out=services --go-grpc_out=services services/routeAnalysisAggregator/proto/v1/route_analysis_aggregator_api_v1.proto

	# Ocean weather service
	protoc -I services --go_out=services --go-grpc_out=services services/oceanWeatherService/proto/v1/ocean_weather_service_api_v1.proto

	# Power train service
	protoc -I services --go_out=services --go-grpc_out=services services/powerTrainService/proto/v1/power_train_service_api_v1.proto

	# Process vibration service
	protoc -I services --go_out=services --go-grpc_out=services services/processVibrationService/proto/v1/process_vibration_service_api_v1.proto

	# Comfort service
	protoc -I services --go_out=services --go-grpc_out=services services/comfortService/proto/v1/comfort_service_api_v1.proto


protoGenPy:
	# Ocean weather service
	python3 -m grpc_tools.protoc -I=services/oceanWeatherService/proto/v1 --python_out=services/oceanWeatherService/proto/v1/generated  --grpc_python_out=services/oceanWeatherService/proto/v1/generated services/oceanWeatherService/proto/v1/ocean_weather_service_api_v1.proto # add 'from .' to line 5 of the _grpc.py file

	# Power train service
	python3 -m grpc_tools.protoc -I=services/powerTrainService/proto/v1 --python_out=services/powerTrainService/proto/v1/generated  --grpc_python_out=services/powerTrainService/proto/v1/generated services/powerTrainService/proto/v1/power_train_service_api_v1.proto # add 'from .' to line 5 of the _grpc.py file

	# Process vibration service
	python3 -m grpc_tools.protoc -I=services/processVibrationService/proto/v1 --python_out=services/processVibrationService/proto/v1/generated --grpc_python_out=services/processVibrationService/proto/v1/generated services/processVibrationService/proto/v1/process_vibration_service_api_v1.proto # add 'from .' to line 5 of the _grpc.py file

	# Comfort service
	python3 -m grpc_tools.protoc -I=services/comfortService/proto/v1 --python_out=services/comfortService/proto/v1/generated --grpc_python_out=services/comfortService/proto/v1/generated services/comfortService/proto/v1/comfort_service_api_v1.proto # add 'from .' to line 5 of the _grpc.py file

protoGenCSharp:
	# protoc -I=services/vesselMotionService/proto/v1 services/vesselMotionService/proto/v1/vessel_motion_service_api_v1.proto --csharp_out=services/vesselMotionService/proto/v1/generated # --grpc_out=services/vesselMotionService/proto/v1/generated --plugin=protoc-gen-grpc=tools\grpc_csharp_plugin.exe
	# NEED TO ADD A COMMAND HERE TO CREATE CLIENT STUBS FOR AGGREGATORS
	
# protoGenCPP:
# 	# Comfort service
# 	protoc -I=services/comfortService/proto/v1 --grpc_out=services/comfortService/proto/v1/generated --plugin=protoc-gen-grpc=/usr/local/bin/grpc_cpp_plugin services/comfortService/proto/v1/comfort_service_api_v1.proto
# 	protoc -I=services/comfortService/proto/v1 --cpp_out=services/comfortService/proto/v1/generated services/comfortService/proto/v1/comfort_service_api_v1.proto

protoGenAll:
	make protoGenGo; make protoGenPy; protoGenCSharp

protoCleanGo:
	find . -type f -name '*.pb.go' -delete

protoCleanPy:
	find . -type f -name '*_grpc.py' -delete
	find . -type f -name '*_pb2.py' -delete

# protoCleanCPP:
# 	find . -type f -name '*.pb.cc' -delete
# 	find . -type f -name '*.pb.h' -delete
# 	find . -type f -name '*.pb.o' -delete

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

testProcessVibrationService:
	/usr/bin/python3 /home/nic/Documents/Work/Masters/Code/mastersCaseStudy/services/processVibrationService/test_processVibrationService.py

testComfortService:
	/usr/bin/python3 /home/nic/Documents/Work/Masters/Code/mastersCaseStudy/services/comfortService/test_comfortService.py

testVesselMotionService:
	cd services/vesselMotionService; dotnet test; cd ..

testGo:
	go test ./...

testPy:
	make testOceanWeatherService
	make testPowerTrainService
	make testProcessVibrationService
	make testComfortService

testCSharp:
	make testVesselMotionService

# testCPP:

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

runProcessVibrationService:
	/usr/bin/python3 /home/nic/Documents/Work/Masters/Code/mastersCaseStudy/services/processVibrationService/processVibrationService.py

runComfortService:
	/usr/bin/python3 /home/nic/Documents/Work/Masters/Code/mastersCaseStudy/services/comfortService/comfortService.py

runRouteAnalysisAggregator:
	go run /home/nic/Documents/Work/Masters/Code/mastersCaseStudy/services/routeAnalysisAggregator/routeAnalysisAggregator.go

runPrometheus:

runPushGateway:

runPy:
	make runOceanWeatherService
	make runPowerTrainService
	make runProcessVibrationService
	make runComfortService

run CSharp:
	make runVesselMotionService

runAll:
	make runPy
	make runCSharp