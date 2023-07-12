package main

import (
	"fmt"
	"time"

	"github.com/montanaflynn/stats"
)

// Prepares the data in the correct format for regression.
func data_prep() ([]stats.Coordinate, []stats.Coordinate, []stats.Coordinate, []stats.Coordinate) {

	x1 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	x2 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	x3 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	x4 := []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8}
	y1 := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}
	y2 := []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74}
	y3 := []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73}
	y4 := []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89}

	// regression1 data
	regData1 := make([]stats.Coordinate, len(x1))
	for i := 0; i < len(x1); i++ {
		regData1[i] = stats.Coordinate{
			X: x1[i],
			Y: y1[i],
		}
	}

	// regression2 data
	regData2 := make([]stats.Coordinate, len(x2))
	for i := 0; i < len(x2); i++ {
		regData2[i] = stats.Coordinate{
			X: x2[i],
			Y: y2[i],
		}
	}

	// regression3 data
	regData3 := make([]stats.Coordinate, len(x3))
	for i := 0; i < len(x3); i++ {
		regData3[i] = stats.Coordinate{
			X: x3[i],
			Y: y3[i],
		}
	}

	// regression4 data
	regData4 := make([]stats.Coordinate, len(x4))
	for i := 0; i < len(x4); i++ {
		regData4[i] = stats.Coordinate{
			X: x4[i],
			Y: y4[i],
		}
	}

	return regData1, regData2, regData3, regData4
}

// Performs regression, then calculates the slope-intercept for each regression.
func regression(data []stats.Coordinate) (float64, float64) {
	reg, _ := stats.LinearRegression(data)

	var x1, y1, x2, y2 float64

	x1 = reg[0].X
	y1 = reg[0].Y

	// In all of these data sets, the 8th element is distinct from the first element.
	// This is a rather crude implementation, but I wanted to see if it would work.
	x2 = reg[7].X
	y2 = reg[7].Y

	slope := (y2 - y1) / (x2 - x1)
	intercept := y1 - slope*x1

	return intercept, slope
}

// Creates the results for all regressions.
func results() error {
	regData1, regData2, regData3, regData4 := data_prep()

	intercept1, slope1 := regression(regData1)
	intercept2, slope2 := regression(regData2)
	intercept3, slope3 := regression(regData3)
	intercept4, slope4 := regression(regData4)

	fmt.Printf("Regression %d, Int: %f, Slope: %f\n\n", 1, intercept1, slope1)
	fmt.Printf("Regression %d, Int: %f, Slope: %f\n\n", 2, intercept2, slope2)
	fmt.Printf("Regression %d, Int: %f, Slope: %f\n\n", 3, intercept3, slope3)
	fmt.Printf("Regression %d, Int: %f, Slope: %f\n\n", 4, intercept4, slope4)

	return nil
}

func main() {

	t0 := time.Now()
	results()
	t1 := time.Now()

	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}
