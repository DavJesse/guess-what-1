package maths

import (
	"math"

	"student/format"
)

func Average(numSlc []int) float64 {
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

func Median(numSlc []int) float64 {
	var median float64
	var medSlc []int

	// Establish a reference index for the median value
	half := len(numSlc) / 2

	numSlc = format.SortSlice(numSlc)

	if len(numSlc) > 2 {
		// If even, the median value will be at index 'half'
		if len(numSlc)%2 != 0 {
			median = float64(numSlc[half])
		} else {
			// if even the median will be the average of the middle values at indices 'half-1' and 'half'
			medSlc = append(medSlc, numSlc[half-1], numSlc[half])
			median = Average(medSlc)
		}
	} else {
		// If one or two items in the slice, the median is the average
		median = Average(numSlc)
	}

	return median
}

func Variance(numSlc []int, mean float64) float64 {
	var sum float64
	var sqr float64
	var result float64
	div := float64(len(numSlc))

	// Calculate sum of the square of the differences between the value of each element and the mean
	for _, n := range numSlc {
		sqr = Square(float64(n) - mean)
		sum += sqr
	}

	// Divide sum of differences by the umber of elements in the slice
	result = sum / div

	return result
}

func StandardDeviation(variance float64) float64 {
	// Calculate square root of variance
	result := math.Sqrt(variance)
	return result
}

func Square(n float64) float64 {
	return n * n
}

func RemoveOutlier(data []int) []int {
	var result []int
	median := Median(data) // Use median as a base for normality

	for _, val := range data {
		
		if float64(val) > median*1.5 || float64(val) < median*0.5 {
			continue
		}
		result = append(result, val)

	}
	return result
}