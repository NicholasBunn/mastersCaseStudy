# Make commands relating to proto files
protoGenGo:

protoGenPy:

protoGenAll:

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
runPrometheus:

runPushGateway:
