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
	numSlc = sliceTransfm(dataSlc)

	fmt.Printf("%q\n", numSlc)
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

	for i := 0; i < len(str); i++ {
		if i < len(str)-len(sep) && str[i:i+len(sep)] == sep {
			result = append(result, token)
			token = ""
		} else {
			token += string(str[i])
		}
	}
	result = append(result, token)
	fmt.Printf("%q\n", result)
	return result
}

func Atoi(str string) (int, string) {
	var result int
	var rtnStr string
	// var intSlc []int
	var digit int
	sign := 1

	for i, r := range str {
		if i == 0 && (r == '-' || r == '+') {
			if r == '-' {
				sign = -1
			}
			continue
		}
		if r < '0' || r > '9' {
			rtnStr = "Your data set contains non-numeric charaters"
		}
		digit = int(r - '0')
		result = result*10 + digit
	}
	result *= sign
	//fmt.Println(result)

	return result, rtnStr
}

func sliceTransfm(strSlc []string) []int {
	var numSlc []int
	var num int
	var str string
	for _, v := range strSlc {
		num, str = Atoi(v)
		if str != "" {
			printLn(str)
			os.Exit(0)
		} else {
			numSlc = append(numSlc, num)
		}
	}
	fmt.Println(numSlc)
	return numSlc
}

func printLn(str string) {
	os.Stdout.WriteString(str + "\n")
}
