package main

func average(numSlc []int) int {
	var numSum int
	var numAve int

	for _, num := range numSlc {
		numSum += num
	}

	numAve = numSum / len(numSlc)

	return numAve
}
