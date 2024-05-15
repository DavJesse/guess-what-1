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

func splitString(str string, sep string) []string {
	var result []string
	var token string

	for i := 0; i<len(str); i++ {
		if i < len(str) - len(sep) && str[i:i+len(sep)] == sep {
			result = append(result, token)
			token = ""
		} else {
			token += string(str[i])
		}
	}
	result = append(result, token)
	return result
}

func Atoi(str string) (int, string) {
	var rtnStr string
	var intSlc []int
	var digit int
	sign := 1

	if str[0] == '-' {
		sign = -1
	}

	for _, r := range str {
		if r >= '0' && r <= '9' {
			digit = int(r - '0')
			intSlc = append(intSlc, digit)
		} else {
			rtnStr = "Data in file contains non-integer characters"
		}
	}
	for _, v := range intSlc {
		
	}
}
