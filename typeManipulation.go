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
	return numSlc
}
