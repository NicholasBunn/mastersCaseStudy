package main

import (
	"context"
	"fmt"
	"testing"
 
	serverPB "github.com/NicholasBunn/mastersCaseStudy/protoFiles/go/webGateway/v1"
)

func TestLogin(t *testing.T) {

	t.Run("Testing standard stuff.", func(t *testing.T) {
		/* This tests some of the core stuff, I'll document this once I've properly written the tests
		*/

		server := loginServer{}

		request := serverPB.LoginRequest{
			Username: "admin",
			Password: "adminPassword",
		}

		response, err := server.Login(context.Background(), &request)

		if (err != nil) {
			fmt.Println("We made a whoopsy: ", err)
		}

		fmt.Println(response, err)
	})
}

func TestRouteAnalysis(t *testing.T) {

	t.Run("Testing standard stuff.", func(t *testing.T) {
		/* This tests some of the core stuff, I'll document this once I've properly written the tests
		*/

		server := routeAnalysisServer{}

		request := serverPB.RouteAnalysisRequest {
			UnixTime: []float64 {1608580803.0},
			Latitude: []float32 {58.7984},
			Longitude: []float32 {17.8081},
			Heading: []float32 {180.0},
			PropPitch: []float32 {0.34654325785352354},
			MotorSpeed: []float32 {0.60128600844793},
			SOG: []float32 {0.030389908256880732},
		}

		response, err := server.RouteAnalysis(context.Background(), &request)

		if (err != nil) {
			fmt.Println("Whoopsy no2: ", err)
		}

		fmt.Println(response, err)
	})
}

func TestRoutePower(t *testing.T) {

	t.Run("Testing standard stuff.", func(t *testing.T) {
		/* This tests some of the core stuff, I'll document this once I've properly written the tests
		*/

		server := routePowerServer{}

		request := serverPB.RoutePowerRequest {
			UnixTime: []float64 {1608580803.0},
			Latitude: []float32 {58.7984},
			Longitude: []float32 {17.8081},
			Heading: []float32 {180.0},
			PropPitch: []float32 {0.34654325785352354},
			MotorSpeed: []float32 {0.60128600844793},
			SOG: []float32 {0.030389908256880732},
		}

		response, err := server.RoutePower(context.Background(), &request)

		if (err != nil) {
			fmt.Println("Whoopsy no3: ", err)
		}

		fmt.Println(response, err)
	})
}

func TestRouteMotion(t *testing.T) {

	t.Run("Testing standard stuff.", func(t *testing.T) {
		/* This tests some of the core stuff, I'll document this once I've properly written the tests
		*/

		server := routeMotionServer{}

		request := serverPB.RouteMotionRequest {
			UnixTime: []float64 {1608580803.0},
			Latitude: []float32 {58.7984},
			Longitude: []float32 {17.8081},
			Heading: []float32 {180.0},
			PropPitch: []float32 {0.34654325785352354},
			MotorSpeed: []float32 {0.60128600844793},
			SOG: []float32 {0.030389908256880732},
		}

		response, err := server.RouteMotion(context.Background(), &request)

		if (err != nil) {
			fmt.Println("Whoopsy no4: ", err)
		}

		fmt.Println(response, err)
	})
}