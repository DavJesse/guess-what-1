package main

import "math"

func average(numSlc []int) float64 {
	var numSum float64
	var aveDiv float64

	size := float64(len(numSlc))

	for _, num := range numSlc {
		numSum += float64(num)
	}

	aveDiv = numSum / size

	return aveDiv
}

func median(numSlc []int) float64 {
	var median float64
	var medSlc []int

	half := len(numSlc) / 2

	for i := 0; i < len(numSlc); i++ {
		for j := i + 1; j < len(numSlc); j++ {
			if numSlc[i] > numSlc[j] {
				numSlc[i], numSlc[j] = numSlc[j], numSlc[i]
			}
		}
	}

	if len(numSlc) > 2 {
		if len(numSlc)%2 != 0 {
			median = float64(numSlc[half])
		} else {
			medSlc = append(medSlc, numSlc[half-1], numSlc[half])
			median = average(medSlc)
		}
	} else {
		median = average(numSlc)
	}

	return median
}

func variance(numSlc []int, mean float64) float64 {
	var sum float64
	var sqr float64
	var result float64
	var div float64 = float64(len(numSlc) - 1)

	for _, n := range numSlc {
		sqr = square(float64(n) - mean)
		sum += sqr
	}

	result = sum / div
	return result
}

func standardDeviation(variance float64) float64 {
	result := math.Sqrt(variance)
	return result
}

func square(n float64) float64 {
	return n * n
}
