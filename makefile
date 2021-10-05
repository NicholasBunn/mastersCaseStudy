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

	# Propeller monitor service
	python3 -m grpc_tools.protoc -I=services/propellerMonitorService/proto/v1 --python_out=protoFiles/python/propellerMonitorService/v1  --grpc_python_out=protoFiles/python/propellerMonitorService/v1   services/propellerMonitorService/proto/v1/propeller_monitor_service_api_v1.proto # add 'from .' to line 5 of the _grpc.py file

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

	# Propeller monitor service
	protoc -I=services/propellerMonitorService/proto/v1/ propeller_monitor_service_api_v1.proto --js_out=import_style=commonjs:protoFiles/javaScript/propellerMonitorService/v1 --grpc-web_out=import_style=commonjs,mode=grpcwebtext:protoFiles/javaScript/propellerMonitorService/v1

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
	python3 services/oceanWeatherService/test_oceanWeatherService.py

testPowerTrainService:
	python3 services/powerTrainService/test_powerTrainService.py

testProcessVibrationService:
	python3 services/processVibrationService/test_processVibrationService.py

testComfortService:
	python3 services/comfortService/test_comfortService.py

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
	python3 services/oceanWeatherService/oceanWeatherService.py
	
runPowerTrainService:
	python3 services/powerTrainService/powerTrainService.py

runVesselMotionService:
	cd services/vesselMotionService; dotnet run; cd ..

runProcessVibrationService:
	python3 services/processVibrationService/processVibrationService.py

runComfortService:
	python3 services/comfortService/comfortService.py

runRouteAnalysisAggregator:
	cd services/routeAnalysisAggregator;	go run routeAnalysisAggregator.go

runPowerTrainAggregator:
	cd services/powerTrainAggregator;	go run powerTrainAggregator.go

runVesselMotionAggregator:
	cd services/vesselMotionAggregator;	go run vesselMotionAggregator.go

runAuthenticationService:
	cd services/authenticationService; go run authenticationService.go

runPropellerMonitorService:
	python3 services/propellerMonitorService/propellerMonitorService.py

runPrometheus:
	docker run -p 127.0.0.1:9090:9090 prometheus

runPushGateway:
	docker run 
runFrontend:
	cd services/webFrontend; node app.js

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
	make runPropellerMonitorService &

runCSharp:
	make runVesselMotionService &
	
runDockerOnly:
	docker run -d -p 10000:10000 --network=host web_proxy &
	docker run -d -p 127.0.0.1:3306:3306/tcp --name user_database -e MYSQL_ROOT_PASSWORD="supersecret" user_database &
	make runPrometheus &
	make runPushGateway &

runAll:
	make runGo &
	make runPy &
	make runCSharp &
	make runDockerOnly &
	make runFrontend &

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