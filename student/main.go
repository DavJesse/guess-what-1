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
	var processed, window []int
	var lowerLimit, upperLimit int
	var content string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := format.Atoi(scanner.Text())
		if err != "" {
			log.Fatal(err)
		}

		dataSet = append(dataSet, num)

		windowSize := 10
		end := len(dataSet)
		start := end - windowSize

		if start < 0 {
			start = 0
		}
		window = dataSet[start:end]

		processed = maths.RemoveOutlier(window)

		// Calculate average
		average := maths.Average(processed)

		upperLimit = int(math.Round(average)) + 85
		lowerLimit = int(math.Round(average)) - 85
		content += fmt.Sprintf("Number: %d\nAverage: %0.2f\nProcessed: %d\n%d - %d\n\n\n", num, average, processed, lowerLimit, upperLimit)
		os.WriteFile("data.txt", []byte(content), 0o644)

		fmt.Printf("%d %d\n", lowerLimit, upperLimit)

	}
}
