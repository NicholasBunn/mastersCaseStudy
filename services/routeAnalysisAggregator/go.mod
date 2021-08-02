module github.com/NicholasBunn/MastersCaseStudy/services/routeAnalysisAggregator

go 1.16

require (
	github.com/NicholasBunn/mastersCaseStudy/serviceSupport/go v0.0.0-20210723115057-7d9b8fd5749f
	github.com/NicholasBunn/mastersCaseStudy/services/routeAnalysisAggregator/proto/v1/generated/oceanWeatherService v0.0.0-20210721084728-4937374dc0f4
	github.com/NicholasBunn/mastersCaseStudy/services/routeAnalysisAggregator/proto/v1/generated/powerTrainService v0.0.0-20210721094440-3e843ee14182
	github.com/NicholasBunn/mastersCaseStudy/services/routeAnalysisAggregator/proto/v1/generated/routeAnalysisAggregator v0.0.0-20210721093602-8cf34122b0d5
	github.com/NicholasBunn/mastersCaseStudy/services/routeAnalysisAggregator/proto/v1/generated/vesselMotionService v0.0.0-20210802082852-289fee8cd766 // indirect
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/golang/protobuf v1.5.2
	google.golang.org/grpc v1.39.0
)
