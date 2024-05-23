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

		//Flag signs that may appear before a number
		if i == 0 && (r == '-' || r == '+') {
			if r == '-' {
				//If '-' update sign to '-1'
				sign = -1
			}
			continue
		}
		if r < '0' || r > '9' {

			//Identify data sets with decimal points
			if r == '.' {
				rtnStr = "This program only takes whole numbers as data input"
			} else {

				//ICapture non-numeric characters and assign an appropriate error message
				rtnStr = "Your data set contains non-numeric charaters"
			}
		}

		//Calculate the digit value of each character
		digit = int(r - '0')

		//Move the digit by a single decimal place to add another digit
		result = result*10 + digit
	}

	//Multiply result by sign to reflect the number's status as a posititive or negative number
	result *= sign

	return result, rtnStr
}

func sliceTransfm(strSlc []string) []int {
	var numSlc []int
	var num int
	var str string
	//Range over each element in the slice, coverting the strings to integers
	for _, v := range strSlc {
		num, str = Atoi(v)

		//Print error messages from invalid strings
		if str != "" {
			printLn(str)
			os.Exit(0)
		} else {

			//Append the converted intgers to 'numSlc'
			numSlc = append(numSlc, num)
		}
	}
	return numSlc
}

// Sort the elements of a slice in ascending order
func sortSlice(numSlc []int) []int {

	for i := 0; i < len(numSlc); i++ {
		for j := i + 1; j < len(numSlc); j++ {

			//Should an element be adjacent to a bigger one, swap them
			if numSlc[i] > numSlc[j] {
				numSlc[i], numSlc[j] = numSlc[j], numSlc[i]
			}
		}
	}

	return numSlc
}
