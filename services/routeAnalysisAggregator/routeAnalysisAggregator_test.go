package main

import (
	"fmt"
	"testing"
)

func TestCalculateRelativeWaveDirection(t *testing.T) {

	var Tests = []struct {
		windDirection 	[]float64
		heading 		[]float64
		expectedOutput 	[]float64
	}{
		{[]float64{30, 64, 130, 194, 298}, []float64{3, 54, 43, 102, 183}, []float64{27, 10, 87, 92, 115},
		},
		{[]float64{30, 64, 130, 194, 298}, []float64{32, 174, 35, 320, 204}, []float64{27, 150, 87, 92, 115},
		},
	}

	t.Run("Testing positive cases of Calculate Relative Wave Direction.", func(t *testing.T) {
		testCase := Tests[0]

		fmt.Println("BLABLA")
		output, err := calculateRelativeWaveDirection(testCase.windDirection, testCase.heading)
		if err != nil{
			fmt.Println(err)
		} else {
			if compareSlices(output, testCase.expectedOutput) != true {
				t.Error("calculate relative wave direction failed with inputs: ", testCase.windDirection, ", ", testCase.heading, ". Expected ", testCase.expectedOutput, ", received ", output)
			} else {
				fmt.Println("Passed")
			}
		}
	})
	
	t.Run("TBLA", func(t *testing.T) {
		testCase := Tests[1]
		fmt.Println(calculateRelativeWaveDirection(testCase.windDirection, testCase.heading))
	})
}

func compareSlices(sliceOne []float64, sliceTwo []float64) (bool) {
	for i, entry := range sliceOne {
        if entry != sliceTwo[i] {
            return false
        }
    }
    return true
}