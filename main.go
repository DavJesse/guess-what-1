package main

import (
	"log"
	"math"

	"math-skills/cli"
	"math-skills/format"
	"math-skills/maths"
	"math-skills/stdout"
)

func main() {
	var dataSet []int

	input := cli.ExtractInput()
	num, err := format.Atoi(input)
	if err != "" {
		log.Fatal(err)
	}

	dataSet = append(dataSet, num)

	// Calculate average and standard deviation
	average := maths.Average(dataSet)
	variance := maths.Variance(dataSet, average)
	standardDeviation := maths.StandardDeviation(variance)

	// Establish upper and lower limits
	upperLimit := int(math.Round(average + 2*standardDeviation))
	lowerLimit := int(math.Round(average - 2*standardDeviation))

	
	stdout.PrintLn(format.Itoa(upperLimit) + " - " + format.Itoa(lowerLimit))
}
