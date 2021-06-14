# Make commands relating to proto files
protoGenGo:
	# Wave service
	# protoc --go_out=. --go-grpc_out=. apis/waveServiceAPI/v1/waveServiceAPI_v1.proto

protoGenPy:
	# Wave service
	python3 -m grpc_tools.protoc -I=services/waveService/proto/v1 --python_out=services/waveService/proto/v1/generated  --grpc_python_out=services/waveService/proto/v1/generated services/waveService/proto/v1/waveServiceAPI_v1.proto

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
testGo:
	go test ./...

testPy:

# Make commands relating to running the program
runWaveService:
	python3 /home/nic/Documents/Work/Masters/Code/mastersCaseStudy/services/waveService/waveService.py
	
runPrometheus:

runPushGateway:
