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
	var dataSet []int = make([]int, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := format.Atoi(scanner.Text())
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

		fmt.Printf("%d --> the standard input\n", num)
		fmt.Printf("%d %d    --> the range for the next input, in this case for the number %d\n", lowerLimit, upperLimit, num)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
