package main

import (
	"fmt"
	"os"
)

func main() {
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
	fmt.Printf("%q\n", dataSlc)
}

func hasSuffix(str, subStr string) bool {
	var status bool
	diff := len(str) - len(subStr)

	if str[diff:] == subStr {
		status = true
	}
	return status
}

func splitString(str, sep string) []string {
	var result []string
	// var newStr string

	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); i++ {
			if str[i:j] == sep {
				// newStr = str[:i]
				result = append(result, str[:i])
				str = str[j+1:]
			}
		}
	}
	return result
}
