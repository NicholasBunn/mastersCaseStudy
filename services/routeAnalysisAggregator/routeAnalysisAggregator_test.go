package main

import (
	"context"
	"fmt"
	"testing"
 
	serverPB "github.com/NicholasBunn/mastersCaseStudy/protoFiles/go/routeAnalysisAggregator/v1"

)

func TestCalculateRelativeWindDirection(t *testing.T) {

	var Tests = []struct {
		windDirection 	[]float32
		heading 		[]float32
		expectedOutput 	[]float32
	}{
		{[]float32{30, 64, 130, 194, 298}, []float32{3, 54, 43, 102, 183}, []float32{27, 10, 87, 92, 115},
		},
		{[]float32{30, 64, 130, 194, 204}, []float32{32, 174, 350, 320, 298}, []float32{358, 250, 140, 234, 266},
		},
	}

	t.Run("Testing Route Analysis Aggregator: Unit Test: Positive Cases for Calculate Relative Wind Direction (Function Test)", func(t *testing.T) {
		/* This tests the core functionality of the Calculate Relative Wind Direction function, where the result is between 0 and 360 degrees (no conversion required).
		*/

		testCase := Tests[0]

		output, err := calculateRelativeWindDirection(testCase.windDirection, testCase.heading)
		if (err != nil){
			fmt.Println(err)
		} else {
			if compareSlices32(output, testCase.expectedOutput) != true {
				t.Error("Calculate relative wind direction failed with inputs: ", testCase.windDirection, ", ", testCase.heading, ". Expected ", testCase.expectedOutput, ", received ", output)
			} else {
				fmt.Println("Calculate Relative Wind Direction: Positive Test Passed")
			}
		}
	})
	
	t.Run("Testing Route Analysis Aggregator: Unit Test: Negative Cases for Calculate Relative Wind Direction (Function Test)", func(t *testing.T) {
		/* This tests the core functionality of the Calculate Relative Wind Direction function, where the result is less than 0 degrees (and thus, conversion is required).
		*/

		testCase := Tests[1]

		output, err := calculateRelativeWindDirection(testCase.windDirection, testCase.heading)
		if (err != nil){
			fmt.Println(err)
		} else {
			if compareSlices32(output, testCase.expectedOutput) != true {
				t.Error("Calculate relative wind direction failed with inputs: ", testCase.windDirection, ", ", testCase.heading, ". Expected ", testCase.expectedOutput, ", received ", output)
			} else {
				fmt.Println("Calculate Relative Wind Direction: Negative Test Passed")
			}
		}
	})
}

func TestCalculateRelativeWindSpeed( t *testing.T) {
	var Tests = []struct {
		windSpeed []float32
		relativeWindDirection []float32
		sog []float32
		expectedOutput 	[]float32
	}{
		{[]float32{10, 43, 32, 21, 18}, []float32{13, 52, 32, 83, 45}, []float32{14, 8, 10, 12, 13},[]float32{23.74370064785235, 34.473443439003304, 37.137539077005634, 14.559256211508096, 25.72792206135786},
		},
		{[]float32{10, 43, 32, 21, 18}, []float32{91, 102, 179, 157, 132}, []float32{14, 8, 10, 12, 13},[]float32{13.825475935627166, -0.9402027051636566, -21.99512624500452, -7.330601922501245, 0.95564908554055},
		},
		{[]float32{10, 43, 32, 21, 18}, []float32{182, 203, 265, 232, 226}, []float32{14, 8, 10, 12, 13},[]float32{4.006091729809043, -31.581708698454932, 7.21101623207494, -0.9288909818388245, 0.49614933173804765},
		},
		{[]float32{10, 43, 32, 21, 18}, []float32{274, 297, 352, 340, 319}, []float32{14, 8, 10, 12, 13},[]float32{14.697564737441255, 27.521591488800514, 41.688578199730244, 31.733545036504076, 26.5847724440099},
		},
		{[]float32{10, 43, 32, 21, 18}, []float32{0, 90, 180, 270, 360}, []float32{14, 8, 10, 12, 13},[]float32{24, 7.999999999999997, -22, 12.000000000000002, 31},
		},
		{[]float32{10, 43, 32, 21, 18}, []float32{364, 375, 457, 967, 634}, []float32{14, 8, 10, 12, 13},[]float32{14.697564737441255, 27.521591488800514, 41.688578199730244, 31.733545036504076, 26.5847724440099},
		},
	}

	t.Run("Testing Route Analysis Aggregator: Unit Test: First Quadrant Cases for Calculate Relative Wind Speed (Function Test)", func(t *testing.T) {
		/* This tests the core functionality of the Calculate Relative Wind Speed function, with inputs between 0 and 90 degrees (top right quadrant)
		*/
		testCase := Tests[0]

		output, err := calculateRelativeWindSpeed(testCase.windSpeed, testCase.relativeWindDirection, testCase.sog)
		if (err != nil){
			fmt.Println(err)
		} else {
			if compareSlices32(output, testCase.expectedOutput) != true {
				t.Error("Calculate relative wind speed failed with inputs: ", testCase.windSpeed, ", ", testCase.relativeWindDirection, ", ", testCase.sog, ". Expected ", testCase.expectedOutput, ", received ", output)			} else {
				fmt.Println("Calculate Relative Wind Speed: First Quadrant Test Passed")
			}
		}
	})

	t.Run("Testing Route Analysis Aggregator: Unit Test: Second Quadrant Cases for Calculate Relative Wind Speed (Function Test)", func(t *testing.T) {
		/* This tests the core functionality of the Calculate Relative Wind Speed function, with inputs between 90 and 180 degrees (bottom right quadrant)
		*/
		testCase := Tests[1]

		output, err := calculateRelativeWindSpeed(testCase.windSpeed, testCase.relativeWindDirection, testCase.sog)
		if (err != nil){
			fmt.Println(err)
		} else {
			if compareSlices32(output, testCase.expectedOutput) != true {
				t.Error("Calculate relative wind speed failed with inputs: ", testCase.windSpeed, ", ", testCase.relativeWindDirection, ", ", testCase.sog, ". Expected ", testCase.expectedOutput, ", received ", output)			} else {
				fmt.Println("Calculate Relative Wind Speed: Second Quadrant Test Passed")
			}
		}
	})

	t.Run("Testing Route Analysis Aggregator: Unit Test: Third Quadrant Cases for Calculate Relative Wind Speed (Function Test)", func(t *testing.T) {
		/* This tests the core functionality of the Calculate Relative Wind Speed function, with inputs between 180 and 270 degrees (bottom left quadrant)
		*/
		testCase := Tests[2]

		output, err := calculateRelativeWindSpeed(testCase.windSpeed, testCase.relativeWindDirection, testCase.sog)
		if (err != nil){
			fmt.Println(err)
		} else {
			if compareSlices32(output, testCase.expectedOutput) != true {
				t.Error("Calculate relative wind speed failed with inputs: ", testCase.windSpeed, ", ", testCase.relativeWindDirection, ", ", testCase.sog, ". Expected ", testCase.expectedOutput, ", received ", output)			} else {
				fmt.Println("Calculate Relative Wind Speed: Third Quadrant Test Passed")
			}
		}
	})

	t.Run("Testing Route Analysis Aggregator: Unit Test: Fourth Quadrant Cases for Calculate Relative Wind Speed (Function Test)", func(t *testing.T) {
		/* This tests the core functionality of the Calculate Relative Wind Speed function, with inputs between 270 and 360 degrees (front left quadrant)
		*/
		testCase := Tests[3]

		output, err := calculateRelativeWindSpeed(testCase.windSpeed, testCase.relativeWindDirection, testCase.sog)
		if (err != nil){
			fmt.Println(err)
		} else {
			if compareSlices32(output, testCase.expectedOutput) != true {
				t.Error("Calculate relative wind speed failed with inputs: ", testCase.windSpeed, ", ", testCase.relativeWindDirection, ", ", testCase.sog, ". Expected ", testCase.expectedOutput, ", received ", output)
			} else {
				fmt.Println("Calculate Relative Wind Speed: Fourth Quadrant Test Passed")
			}
		}
	})

	// Add in edge case tests (0, 90, 180, 270, 360)
	t.Run("Testing Route Analysis Aggregator: Unit Test: Edge Cases for Calculate Relative Wind Speed (Function Test)", func(t *testing.T) {
		/* This tests the core functionality of the Calculate Relative Wind Speed function, with inputs of 0, 90, 180, 270, and 360 degrees (edge cases)
		*/
		testCase := Tests[4]

		output, err := calculateRelativeWindSpeed(testCase.windSpeed, testCase.relativeWindDirection, testCase.sog)
		if (err != nil){
			fmt.Println(err)
		} else {
			if compareSlices32(output, testCase.expectedOutput) != true {
				t.Error("Calculate relative wind speed failed with inputs: ", testCase.windSpeed, ", ", testCase.relativeWindDirection, ", ", testCase.sog, ". Expected ", testCase.expectedOutput, ", received ", output)
			} else {
				fmt.Println("Calculate Relative Wind Speed: Edge Case Test Passed")
			}
		}
	})

	t.Run("Testing Route Analysis Aggregator: Unit Test: Incorrect Input Handling Cases for Calculate Relative Wind Speed (Function Test)", func(t *testing.T) {
		/* This tests the error-handling functionality of the Calculate Relative Wind Speed function, where inputs exceed 360 degrees
		*/
		testCase := Tests[5]

		output, err := calculateRelativeWindSpeed(testCase.windSpeed, testCase.relativeWindDirection, testCase.sog)
		if (err != nil){
			fmt.Println("Calculate Relative Wind Speed: Incorrect Input Value Passed")
		} else {
			t.Error("Calculate relative wind speed failed with inputs: ", testCase.windSpeed, ", ", testCase.relativeWindDirection, ". Expected failure, received ", output)		
		}
	})

}

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
