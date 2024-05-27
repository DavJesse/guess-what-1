package main

import (
	"math"
	"os"
)

func main() {
	// Exit function if too many or not enough arguments are given
	if len(os.Args) != 2 {
		printLn("This program only takes one file as input")
		return
	}

	// Capture the name of data file, store it in variable 'arg'
	arg := os.Args[1]

	// Check whether file has a '.txt' extension
	if !hasSuffix(arg, ".txt") {
		printLn("wrong file format!")
		printLn("Please enter a valid text file with a '.txt' extension")
		return
	}

	// Read file to extract data
	rawFile, err := os.ReadFile(arg)
	if err != nil {
		printF("math-skills encountered a problem in reading ", arg)
		printF("Check ", arg+" and re-submit a valid file")
		return
	}

	// Exist program ashould 'arg' be empty
	if len(rawFile) == 0 {
		printF(arg, " is empty!")
		printF("Populate ", arg+" and try again")
		return
	}

	// Store data extracted from file in the variable 'file'
	file := string(rawFile)

	// Split 'file' at '\n' and convert to []int
	dataSlc := splitString(file, "\n")
	numSlc := sliceTransfm(dataSlc)

	// Calculate average, median, variance, standard deviation
	average := average(numSlc)
	median := median(numSlc)
	variance := variance(numSlc, average)
	standardDeviation := standardDeviation(variance)

	// Round off the values calculated above to the next whole number
	ave := int(math.Round(average))
	med := int(math.Round(median))
	vrnc := int(math.Round(variance))
	sD := int(math.Round(standardDeviation))

	// Print average, median, variance, standard deviation values to terminal
	printLn("Average: " + Itoa(ave))
	printLn("Median: " + Itoa(med))
	printLn("Variance: " + Itoa(vrnc))
	printLn("Standard Deviation: " + Itoa(sD))
}
