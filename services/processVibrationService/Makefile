LDFLAGS = -L/usr/local/lib `pkg-config --libs protobuf grpc++`\
           -Wl,--no-as-needed -lgrpc++_reflection -Wl,--as-needed\
           -ldl

CXX = g++
CPPFLAGS += `pkg-config --cflags protobuf grpc`
CXXFLAGS += -std=c++11

GRPC_CPP_PLUGIN = grpc_cpp_plugin
GRPC_CPP_PLUGIN_PATH ?= `which $(GRPC_CPP_PLUGIN)`

build: processVibrationService

processVibrationService: proto/v1/generated/process_vibration_service_api_v1.pb.o proto/v1/generated/process_vibration_service_api_v1.grpc.pb.o processVibrationService.o
	$(CXX) $^ $(LDFLAGS) -o $@

