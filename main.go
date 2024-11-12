package main

import (
	"math"
	"os"

	"math-skills/format"
	"math-skills/maths"
	"math-skills/stdout"
)

func main() {
	// Exit function if too many or not enough arguments are given
	if len(os.Args) != 2 {
		stdout.PrintLn("This program only takes one file as input")
		return
	}

	// Capture the name of data file, store it in variable 'arg'
	arg := os.Args[1]

	// Check whether file has a '.txt' extension
	if !format.HasSuffix(arg, ".txt") {
		stdout.PrintLn("wrong file format!")
		stdout.PrintLn("Please enter a valid text file with a '.txt' extension")
		return
	}

	// Read file to extract data
	rawFile, err := os.ReadFile(arg)
	if err != nil {
		stdout.PrintF("math-skills encountered a problem in reading ", arg)
		stdout.PrintF("Check ", arg+" and re-submit a valid file")
		return
	}

	// Exist program ashould 'arg' be empty
	if len(rawFile) == 0 {
		stdout.PrintF(arg, " is empty!")
		stdout.PrintF("Populate ", arg+" and try again")
		return
	}

	// Store data extracted from file in the variable 'file'
	file := string(rawFile)

	// Split 'file' at '\n' and convert to []int
	dataSlc := format.SplitString(file, "\n")
	numSlc := format.SliceTransfm(dataSlc)

	// Calculate average, median, variance, standard deviation
	average := maths.Average(numSlc)
	median := maths.Median(numSlc)
	variance := maths.Variance(numSlc, average)
	standardDeviation := maths.StandardDeviation(variance)

	// Round off the values calculated above to the next whole number
	ave := int(math.Round(average))
	med := int(math.Round(median))
	vrnc := int(math.Round(variance))
	sD := int(math.Round(standardDeviation))

	// Print average, median, variance, standard deviation values to terminal
	stdout.PrintLn("Average: " + format.Itoa(ave))
	stdout.PrintLn("Median: " + format.Itoa(med))
	stdout.PrintLn("Variance: " + format.Itoa(vrnc))
	stdout.PrintLn("Standard Deviation: " + format.Itoa(sD))
}
