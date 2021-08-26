# Make commands relating to proto files
protoGenGo:
	# Ocean weather service
	protoc -I services --go_out=protoFiles/go --go-grpc_out=protoFiles/go services/oceanWeatherService/proto/v1/ocean_weather_service_api_v1.proto

	# Power train service
	protoc -I services --go_out=protoFiles/go --go-grpc_out=protoFiles/go services/powerTrainService/proto/v1/power_train_service_api_v1.proto

	# Process vibration service
	protoc -I services --go_out=protoFiles/go --go-grpc_out=protoFiles/go services/processVibrationService/proto/v1/process_vibration_service_api_v1.proto

	# Comfort service
	protoc -I services --go_out=protoFiles/go --go-grpc_out=protoFiles/go services/comfortService/proto/v1/comfort_service_api_v1.proto

	# Vessel motion service
	protoc -I services --go_out=protoFiles/go --go-grpc_out=protoFiles/go services/vesselMotionService/proto/v1/vessel_motion_service_api_v1.proto

	# Route analysis aggregator
	protoc -I services --go_out=protoFiles/go --go-grpc_out=protoFiles/go services/routeAnalysisAggregator/proto/v1/route_analysis_aggregator_api_v1.proto

	# Power train aggregator
	protoc -I services --go_out=protoFiles/go --go-grpc_out=protoFiles/go services/powerTrainAggregator/proto/v1/power_train_aggregator_api_v1.proto

	# Vessel motion aggregator
	protoc -I services --go_out=protoFiles/go --go-grpc_out=protoFiles/go services/vesselMotionAggregator/proto/v1/vessel_motion_aggregator_api_v1.proto

	# Authentication service
	protoc -I services --go_out=protoFiles/go --go-grpc_out=protoFiles/go services/authenticationService/proto/v1/authentication_service_api_v1.proto

	# Web gateway
	protoc -I services --go_out=protoFiles/go --go-grpc_out=protoFiles/go services/webGateway/proto/v1/web_gateway_api_v1.proto

protoGenPy:
	# Ocean weather service
	python3 -m grpc_tools.protoc -I=services/oceanWeatherService/proto/v1 --python_out=protoFiles/python/oceanWeatherService/v1  --grpc_python_out=protoFiles/python/oceanWeatherService/v1   services/oceanWeatherService/proto/v1/ocean_weather_service_api_v1.proto # add 'from .' to line 5 of the _grpc.py file

	# Power train service
	python3 -m grpc_tools.protoc -I=services/powerTrainService/proto/v1 --python_out=protoFiles/python/powerTrainService/v1  --grpc_python_out=protoFiles/python/powerTrainService/v1   services/powerTrainService/proto/v1/power_train_service_api_v1.proto # add 'from .' to line 5 of the _grpc.py file

	# Process vibration service
	python3 -m grpc_tools.protoc -I=services/processVibrationService/proto/v1 --python_out=protoFiles/python/processVibrationService/v1  --grpc_python_out=protoFiles/python/processVibrationService/v1   services/processVibrationService/proto/v1/process_vibration_service_api_v1.proto # add 'from .' to line 5 of the _grpc.py file

	# Comfort service
	python3 -m grpc_tools.protoc -I=services/comfortService/proto/v1 --python_out=protoFiles/python/comfortService/v1  --grpc_python_out=protoFiles/python/comfortService/v1   services/comfortService/proto/v1/comfort_service_api_v1.proto # add 'from .' to line 5 of the _grpc.py file

protoGenCSharp:
	# Vessel motion service
	# protoc -I=services/vesselMotionService/proto/v1 services/vesselMotionService/proto/v1/vessel_motion_service_api_v1.proto --csharp_out=services/vesselMotionService/proto/v1/generated # --grpc_out=services/vesselMotionService/proto/v1/generated --plugin=protoc-gen-grpc=tools\grpc_csharp_plugin.exe
	# NEED TO ADD A COMMAND HERE TO CREATE CLIENT STUBS FOR AGGREGATORS
	
# protoGenCPP:
# 	# Comfort service
# 	protoc -I=services/comfortService/proto/v1 --grpc_out=services/comfortService/proto/v1/generated --plugin=protoc-gen-grpc=/usr/local/bin/grpc_cpp_plugin services/comfortService/proto/v1/comfort_service_api_v1.proto
# 	protoc -I=services/comfortService/proto/v1 --cpp_out=services/comfortService/proto/v1/generated services/comfortService/proto/v1/comfort_service_api_v1.proto

protoGenJS:
	# Web gateway
	protoc -I=services/webGateway/proto/v1/ web_gateway_api_v1.proto --js_out=import_style=commonjs:protoFiles/javaScript/webGateway/v1 --grpc-web_out=import_style=commonjs,mode=grpcweb:protoFiles/javaScript/webGateway/v1

	# Route analysis aggregator
	protoc -I=services/routeAnalysisAggregator/proto/v1/ route_analysis_aggregator_api_v1.proto --js_out=import_style=commonjs:protoFiles/javaScript/routeAnalysisAggregator/v1 --grpc-web_out=import_style=commonjs,mode=grpcweb:protoFiles/javaScript/routeAnalysisAggregator/v1

	# Power train aggregator
	protoc -I=services/powerTrainAggregator/proto/v1/ power_train_aggregator_api_v1.proto --js_out=import_style=commonjs:protoFiles/javaScript/powerTrainAggregator/v1 --grpc-web_out=import_style=commonjs,mode=grpcweb:protoFiles/javaScript/powerTrainAggregator/v1

	# Vessel motion aggregator
	protoc -I=services/vesselMotionAggregator/proto/v1/ vessel_motion_aggregator_api_v1.proto --js_out=import_style=commonjs:protoFiles/javaScript/vesselMotionAggregator/v1 --grpc-web_out=import_style=commonjs,mode=grpcweb:protoFiles/javaScript/vesselMotionAggregator/v1

	# Authentication service
	protoc -I=services/authenticationService/proto/v1/ authentication_service_api_v1.proto --js_out=import_style=commonjs:protoFiles/javaScript/authenticationService/v1 --grpc-web_out=import_style=commonjs,mode=grpcweb:protoFiles/javaScript/authenticationService/v1

protoGenAll:
	make protoGenGo; make protoGenPy; protoGenCSharp; make protoGenJS

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

testRouteAnalysisAggregator:
	cd services/routeAnalysisAggregator; go test; cd ..

testPowerTrainAggregator:
	cd services/powerTrainAggregator; go test; cd ..

testVesselMotionAggregator:
	cd services/vesselMotionAggregator; go test; cd ..

testAuthenticationStuff:
	cd generalComponents/authenticationStuff; go test; cd ..
	
testAuthenticationService:
	cd services/authenticationService; go test; cd ..

testWebGateway:
	cd services/webGateway; go test; cd ..

testGo:
	make testAuthenticationStuff
	make testAuthenticationService
	# make testRouteAnalysisAggregator
	# make testPowerTrainAggregator

testPy:
	make testOceanWeatherService
	make testPowerTrainService
	make testProcessVibrationService
	make testComfortService

testCSharp:
	make testVesselMotionService

testAll:
	make testPy
	make testCSharp
	make testGo

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
	cd /home/nic/Documents/Work/Masters/Code/mastersCaseStudy/services/routeAnalysisAggregator;	go run routeAnalysisAggregator.go

runPowerTrainAggregator:
	cd /home/nic/Documents/Work/Masters/Code/mastersCaseStudy/services/powerTrainAggregator;	go run powerTrainAggregator.go

runVesselMotionAggregator:
	cd /home/nic/Documents/Work/Masters/Code/mastersCaseStudy/services/vesselMotionAggregator;	go run vesselMotionAggregator.go

runAuthenticationService:
	cd /home/nic/Documents/Work/Masters/Code/mastersCaseStudy/services/authenticationService; go run authenticationService.go

runPrometheus:

runPushGateway:

runServiceLayer:
	make runOceanWeatherService &
	make runPowerTrainService &
	make runVesselMotionService &
	make runProcessVibrationService &
	make runComfortService &

runAggregatorLayer:
	make runRouteAnalysisAggregator &
	make runPowerTrainAggregator &
	make runVesselMotionAggregator &

runGatewayLayer:
	make runAuthenticationService &
	
runGo:
	make runRouteAnalysisAggregator &
	make runPowerTrainAggregator &
	make runVesselMotionAggregator &
	make runAuthenticationService &

runPy:
	make runOceanWeatherService &
	make runPowerTrainService &
	make runProcessVibrationService &
	make runComfortService &

runCSharp:
	make runVesselMotionService &
	
runDockerOnly:
	docker run -d -p 8080:8080 --network=host web_proxy &
	docker run -d -p 127.0.0.1:3306:3306/tcp --name user_database -e MYSQL_ROOT_PASSWORD="supersecret" user_database &

runAll:
	make runGo &
	make runPy &
	make runCSharp &
	make runDockerOnly &

killGo:
	sudo killall vesselMotionAgg &
	sudo killall routeAnalysisAg &
	sudo killall authenticationS &
	sudo killall routeAnalysisAg &
	sudo killall powerTrainAggre &

killPy:
	sudo killall python3

killCSharp:
	sudo killall vesselMotionSer

killAll:
	make killGo
	make killPy
	make killCSharp