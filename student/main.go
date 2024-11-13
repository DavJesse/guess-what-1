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
	var workData []int
	var lowerLimit, upperLimit int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := format.Atoi(scanner.Text())
		if err != "" {
			log.Fatal(err)
		}
		dataSet = append(dataSet, num)

		if len(dataSet) > 14 {
			workData = dataSet[len(dataSet)-15:]
		} else {
			workData = dataSet
		}

		// Calculate average and standard deviation
		average := maths.Average(workData)
		variance := maths.Variance(workData, average)
		standardDeviation := maths.StandardDeviation(variance)

if len(workData) < 5 {
	upperLimit = 200
	lowerLimit = 0
} else {

	// Establish upper and lower limits
	upperLimit = int(math.Round(average + 2.5*standardDeviation))
	lowerLimit = int(math.Round(average - 3.9*standardDeviation))

	if upperLimit < 200 {
		upperLimit = 200
	}

	if lowerLimit < 100 {
		lowerLimit = 100
	}

}


		fmt.Printf("%d %d\n", lowerLimit, upperLimit)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}