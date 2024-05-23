package main

import (
	"math"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		printLn("This program only takes one file as input")
		return
	}
	arg := os.Args[1]

	if !hasSuffix(arg, ".txt") {
		printLn("wrong file format!")
		printLn("Please enter a valid text file with a '.txt' extension")
		return
	}

	rawFile, err := os.ReadFile(arg)
	if err != nil {
		printF("Math-skills encountered a problem in reading ", arg)
		printF("Check " + arg + "and re-submit a valid file\n")
		return
	}

	file := string(rawFile)
	dataSlc := splitString(file, "\n")

	numSlc := sliceTransfm(dataSlc)

	average := average(numSlc)
	median := median(numSlc)
	variance := variance(numSlc, average)
	standardDeviation := standardDeviation(variance)

	ave := int(math.Round(average))
	med := int(math.Round(median))
	vrnc := int(math.Round(variance))
	sD := int(math.Round(standardDeviation))

	printLn("Average: " + Itoa(ave))
	printLn("Median: " + Itoa(med))
	printLn("Variance: " + Itoa(vrnc))
	printLn("Standard Deviation: " + Itoa(sD))
}
