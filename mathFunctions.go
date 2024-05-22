package main

import (
	"math"
)

// Calculates the average value of integers presented in a slice of integers([]int)
func average(numSlc []int) float64 {
	var numSum float64
	var aveDiv float64

	size := float64(len(numSlc))

	// Range over the slice, calculating the total value of the elements
	for _, num := range numSlc {
		numSum += float64(num)
	}

	// Calculate average by dividing the sum of the elements by the number of elements
	aveDiv = numSum / size

	return aveDiv
}

// Calculates the median value of integers presented in a slice of integers([]int)
func median(numSlc []int) float64 {
	var median float64
	var medSlc []int

	//Establish a reference index for the median value
	half := len(numSlc) / 2

	numSlc = sortSlice(numSlc)

	//Establish whether the number of elements in the slice is even or not
	if len(numSlc) > 2 {

		//If even, the median value will be at index 'half'
		if len(numSlc)%2 != 0 {
			median = float64(numSlc[half])
		} else {
			//if even the median will be the average of the middle values at indices 'half-1' and 'half'
			medSlc = append(medSlc, numSlc[half-1], numSlc[half])
			median = average(medSlc)
		}
	} else {

		//If only two items in the slice, the median is the average
		median = average(numSlc)
	}

	return median
}

func variance(numSlc []int, mean float64) float64 {
	var sum float64
	var sqr float64
	var result float64
	var div float64 = float64(len(numSlc))

	//Calculate sum of the square of the differences between the value of each element and the mean
	for _, n := range numSlc {
		sqr = square(float64(n) - mean)
		sum += sqr
	}

	//Divide sum of differences by the umber of elements in the slice
	result = sum / div

	return result
}

// Calculates the Standard Deviation of data set
func standardDeviation(variance float64) float64 {

	//Calculate square root of variance
	result := math.Sqrt(variance)
	return result
}

// Calculates the square of any integer
func square(n float64) float64 {
	return n * n
}
