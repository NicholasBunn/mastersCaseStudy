package main

import (
	"context"
	"fmt"
	"testing"
 
	serverPB "github.com/NicholasBunn/mastersCaseStudy/protoFiles/go/routeAnalysisAggregator/v1"

)

func TestAnalyseRoute(t *testing.T) {
	// var Tests = []struct {
	// 	name string
	// 	request *serverPB.AnalysisRequest
	// }{
	// 	{name: "Standard test",
	// 	request: &serverPB.AnalysisRequest{
	// 		UnixTime: {1608580803.0},
	// 		Latitude: {58.7984},
	// 		Longitude: {17.8081},
	// 		Heading: {15},
	// 		PropPitch: {0.26854406323815316},
	// 		MotorSpeed: {0.597549569477592},
	// 		SOG: {0.030389908256880732},
	// 		},
	// 	},
	// }

	t.Run("Testing standard stuff.", func(t *testing.T) {
		/* This tests some of the core stuff, I'll document this once I've properly written the tests
		*/

		server := server{}

		request := serverPB.AnalysisRequest{
			UnixTime: []float64{1608580803.0},
			Latitude: []float32{58.7984},
			Longitude: []float32{17.8081},
			Heading: []float32{15},
			PropPitch: []float32{0.26854406323815316},
			MotorSpeed: []float32{0.597549569477592},
			SOG: []float32{0.030389908256880732},
			}

		
		response, err := server.AnalyseRoute(context.Background(), &request)
		
		if (err != nil) {
			fmt.Println("SAD: ", err)
		}
		fmt.Println(response)
	})
}

func compareSlices32(sliceOne []float32, sliceTwo []float32) (bool) {
	for i, entry := range sliceOne {
        if entry != sliceTwo[i] {
            return false
        }
    }
    return true
}
