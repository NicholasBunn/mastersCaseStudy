package main

import (
	"context"
	"fmt"
	"testing"
 
	serverPB "github.com/NicholasBunn/mastersCaseStudy/protoFiles/go/powerTrainAggregator/v1"
)

func TestEstimateRoute(t *testing.T) {
	t.Run("Testing standard stuff.", func(t *testing.T) {
		/* This tests some of the core stuff, I'll document this once I've properly written the tests
		*/
	
		server := server{}

		request := serverPB.PTEstimateRequest {
			UnixTime: []float64{1608811845.0, 1608812145.0, 1609157745},
			Latitude: []float32{58.7984, 58.8084, 58.8123},
			Longitude: []float32{17.8081, 17.8081, 17,9123},
			Heading: []float32{15, 20, 12},
			PropPitch: []float32{-40.5200004577637, 51.3299980163574, 95.3400039672852},
			MotorSpeed: []float32{83.5450057983399, 104.089996337891, 120.443740844727},
			SOG: []float32{0.545311111064, 2.973488888632, 13.756244443256},
		}

		response, err := server.EstimatePowerTrain(context.Background(), &request)
		
		if (err != nil) {
			fmt.Println("SAD: ", err)
		}
		fmt.Println(response)
	})
}