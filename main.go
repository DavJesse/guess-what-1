package main

import (
	"fmt"
	"os"
)

func main() {
	var numSlc []int

	if len(os.Args) != 2 {
		os.Stdout.WriteString("This program only takes one file as input\n")
		return
	}
	arg := os.Args[1]

	if !hasSuffix(arg, ".txt") {
		os.Stdout.WriteString("wrong file format!\n")
		os.Stdout.WriteString("Please enter a valid text file with a '.txt' extension\n")
		return
	}

	rawFile, err := os.ReadFile(arg)
	if err != nil {
		os.Stdout.WriteString("Math-skills encountered a problem in reading " + arg + "\n")
		os.Stdout.WriteString("Check " + arg + "and re-submit a valid file\n")
		return
	}

	file := string(rawFile)

	// os.Stdout.WriteString(file + "\n")
	dataSlc := splitString(file, "\n")
	fmt.Printf("data slice: %q\n", dataSlc)
	numSlc = sliceTransfm(dataSlc)

	fmt.Println("number slice: ", numSlc)

	average := average(numSlc)

	printLn("Average: " + Itoa(average))
}
