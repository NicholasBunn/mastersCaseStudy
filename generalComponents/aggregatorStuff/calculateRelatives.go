package aggregator

func calculateRelativeWindDirection(windDirection []float32, heading []float32) ([]float32, error) {
	/* This function takes the wind direction and vessel heading as inputs. Using these, it 
	calculates and returns the wind direction relative to the vessels direction.
	*/

	var relativeWindDirection []float32
	var tempRelativeWindDirection float32

	for index, windDirectionInstance := range windDirection {
		tempRelativeWindDirection = windDirectionInstance - heading[index]

		// Check whether the relative wind direction is negative and add 360 degrees until it is positive so that all directions returned are on the same coordinate system.
		for (tempRelativeWindDirection < 0) {
			tempRelativeWindDirection += 360
		}
		relativeWindDirection = append(relativeWindDirection, tempRelativeWindDirection)
	}

	return relativeWindDirection, nil
}

func calculateRelativeWindSpeed(windSpeed []float32, relativeWindDirection []float32, sog []float32) ([]float32, error) {
	/* This function takes the wind speed, wind direction, vessel speed, and vessel direction as inputs. Uisng these, it calculates and returns the wind speed relative to the vessel's speed.
	*/

	var relativeWindSpeed []float32

	for index, windSpeedInstance := range windSpeed {

		// Decompose the wind speed to get the component of wind acting head on to the ship and add that to the ship's SOG, done on a case-by-case basis depending on where the wind is incident on the ship. PS: There's a lot of type conversion going on here so it looks a bit messy, but it's just basic trigonometry
		if (relativeWindDirection[index] < 90) {
			relativeWindSpeed = append(relativeWindSpeed, float32(float64(sog[index]) + float64(windSpeedInstance)*math.Cos(float64(relativeWindDirection[index])*math.Pi/180)))
		} else if (relativeWindDirection[index] < 180) {
			relativeWindSpeed = append(relativeWindSpeed, float32(float64(sog[index]) - float64(windSpeedInstance)*math.Cos((180 - float64(relativeWindDirection[index]))*math.Pi/180)))
		} else if (relativeWindDirection[index] < 270) {
			relativeWindSpeed = append(relativeWindSpeed, float32(float64(sog[index]) - float64(windSpeedInstance)*math.Cos((float64(relativeWindDirection[index])-180)*math.Pi/180)))
		} else if (relativeWindDirection[index] <= 360) {
			relativeWindSpeed = append(relativeWindSpeed, float32(float64(sog[index]) + float64(windSpeedInstance)*math.Cos((360 - float64(relativeWindDirection[index]))*math.Pi/180)))
		} else {
			return nil, fmt.Errorf("Provided relative wind direction is out of range (greater than 360 degrees)")
		}
	}

	return relativeWindSpeed, nil
}