# Make commands relating to proto files
protoGenGo:
	# Wave service
	# protoc --go_out=. --go-grpc_out=. apis/wave_service_api/v1/wave_service_api_v1.proto

protoGenPy:
	# Wave service
	python3 -m grpc_tools.protoc -I=services/waveService/proto/v1 --python_out=services/waveService/proto/v1/generated  --grpc_python_out=services/waveService/proto/v1/generated services/waveService/proto/v1/wave_service_api_v1.proto # add 'from .' to line 5 of the _grpc.py file

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
testWaveService:
	/usr/bin/python3 /home/nic/Documents/Work/Masters/Code/mastersCaseStudy/services/waveService/test_waveService.py

testGo:
	go test ./...

testPy:
	make testWaveService

testAll:
	make testPy

# Make commands relating to running the program
runWaveService:
	/usr/bin/python3 /home/nic/Documents/Work/Masters/Code/mastersCaseStudy/services/waveService/waveService.py
	
runPrometheus:

runPushGateway:

runPy:
	make runWaveService

runAll:
	make runPy