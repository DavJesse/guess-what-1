package main

import (
	"os"
)

func Atoi(str string) (int, string) {
	var result int
	var rtnStr string

	var digit int
	sign := 1

	for i, r := range str {

		// Flag signs that may appear before a number
		if i == 0 && (r == '-' || r == '+') {
			if r == '-' {
				sign = -1
			}
			continue
		}

		// identify non-numeric characters
		if r < '0' || r > '9' {
			if r == '.' {
				rtnStr = "This program only takes whole numbers as data input"
			} else {
				rtnStr = "Your data set contains non-numeric charaters"
			}
		}

		// Calculate the digit value of each character and eventually result
		digit = int(r - '0')
		result = result*10 + digit
	}

	result *= sign

	return result, rtnStr
}

func sliceTransfm(strSlc []string) []int {
	var numSlc []int
	var num int
	var str string

	for _, v := range strSlc {
		num, str = Atoi(v)

		// Print error messages from invalid strings
		if str != "" {
			printLn(str)
			os.Exit(0)
		} else {
			numSlc = append(numSlc, num)
		}
	}
	return numSlc
}

func sortSlice(numSlc []int) []int {
	for i := 0; i < len(numSlc); i++ {
		for j := i + 1; j < len(numSlc); j++ {
			// Should an element be adjacent to a bigger one, swap them
			if numSlc[i] > numSlc[j] {
				numSlc[i], numSlc[j] = numSlc[j], numSlc[i]
			}
		}
	}

	return numSlc
}
