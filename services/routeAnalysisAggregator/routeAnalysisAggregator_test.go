package main

import (
	"fmt"
	"testing"
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

	t.Run("Testing positive cases for Calculate Relative Wind Direction.", func(t *testing.T) {
		/* This tests the core functionality of the Calulcate Relative Wind Direction function, where the result is between 0 and 360 degrees (no conversion required).
		*/

		testCase := Tests[0]

		output, err := calculateRelativeWindDirection(testCase.windDirection, testCase.heading)
		if err != nil{
			fmt.Println(err)
		} else {
			if compareSlices(output, testCase.expectedOutput) != true {
				t.Error("Calculate relative wind direction failed with inputs: ", testCase.windDirection, ", ", testCase.heading, ". Expected ", testCase.expectedOutput, ", received ", output)
			} else {
				fmt.Println("Calcualte Relative Wind Direction: Positive Test Passed")
			}
		}
	})
	
	t.Run("Testing negative cases for Calculate Relative Wind Direction ", func(t *testing.T) {
		/* This tests the core functionality of the Calculate Relative Wind Direction function, where the result is less than 0 degrees (and thus, conversion is required).
		*/

		testCase := Tests[1]

		output, err := calculateRelativeWindDirection(testCase.windDirection, testCase.heading)
		if err != nil{
			fmt.Println(err)
		} else {
			if compareSlices(output, testCase.expectedOutput) != true {
				t.Error("Calculate relative wind direction failed with inputs: ", testCase.windDirection, ", ", testCase.heading, ". Expected ", testCase.expectedOutput, ", received ", output)
			} else {
				fmt.Println("Calculate Relative Wind Direction: Negative Test Passed")
			}
		}
	})
}

func compareSlices(sliceOne []float32, sliceTwo []float32) (bool) {
	for i, entry := range sliceOne {
        if entry != sliceTwo[i] {
            return false
        }
    }
    return true
}