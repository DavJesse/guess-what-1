package main

func average(numSlc []int) int {
	var numSum int
	var aveDiv int
	var aveMod int

	for _, num := range numSlc {
		numSum += num
	}

	aveDiv = numSum / len(numSlc)
	aveMod = numSum % len(numSlc)

	if float64(aveMod) >= float64(len(numSlc))/2 {
		aveDiv++
	}

	return aveDiv
}

func median(numSlc []int) int {
	var median int
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
			median = numSlc[half+1]
		} else {
			medSlc = append(medSlc, numSlc[half-1], numSlc[half])
			median = average(medSlc)
		}
	} else {
		median = average(numSlc)
	}

	return median
}
