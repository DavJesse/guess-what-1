package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"

	"student/format"
	"student/maths"
)

func main() {
	var dataSet []int
	var workData, processed []int
	// var containsOutlier bool
	var lowerLimit, upperLimit int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := format.Atoi(scanner.Text())
		if err != "" {
			log.Fatal(err)
		}
		dataSet = append(dataSet, num)

		if len(dataSet) > 9 {
			workData = dataSet[len(dataSet)-10:]
			// fmt.Printf("Before outlier removal: %d\n", workData)

			processed, _ = maths.RemoveOutlier(workData)
			// fmt.Printf("After outlier removal: %d\n", processed)
		} else {
			workData = dataSet
			processed = workData
		}

		// Calculate average
		average := maths.Average(processed)

		// if containsOutlier {
		// 	upperLimit = int(maths.Median(processed)) + 85
		// 	lowerLimit = int(maths.Median(processed)) - 85
		// } else {
		upperLimit = int(math.Round(average)) + 85
		lowerLimit = int(math.Round(average)) - 85
		// }

		fmt.Printf("%d %d\n", lowerLimit, upperLimit)

	}
}
