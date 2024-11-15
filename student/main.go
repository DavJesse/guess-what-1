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
	var dataSet, bef []int
	//var processed []int
	var lowerLimit, upperLimit int
	var content string
	var count int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := format.Atoi(scanner.Text())
		if err != "" {
			log.Fatal(err)
		}
		count++
		if count < 11 {
			bef = dataSet
			dataSet = append(dataSet, num)
		} else {
			bef = dataSet[1:]
			dataSet= append(dataSet[1:], num)
		}

		dataSet = maths.RemoveOutlier(dataSet)

		// Calculate average
		average := maths.Average(dataSet)

		upperLimit = int(math.Round(average)) + 85
		lowerLimit = int(math.Round(average)) - 85
		content += fmt.Sprintf("Number: %d\nAverage: %0.2f\nBefore Append: %d\nDataSet: %d\n%d - %d\n\n\n", num, average, bef, dataSet, lowerLimit, upperLimit)
		os.WriteFile("data.txt", []byte(content), 0o644)

		fmt.Printf("%d %d\n", lowerLimit, upperLimit)

	}
}
