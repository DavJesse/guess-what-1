package main

// Detect whether a string starts with a specific set of characters
func hasSuffix(str, subStr string) bool {
	var status bool
	diff := len(str) - len(subStr)

	//Check is string starts with substring
	if str[diff:] == subStr {
		status = true
	}

	return status
}

// Split a string at a separator
func splitString(str string, sep string) []string {
	var result []string
	var token string

	//Empty and populate 'token' with elements occuring before the separator
	for i := 0; i < len(str); i++ {
		if i <= len(str)-len(sep) && str[i:i+len(sep)] == sep {

			//Append 'token' to result
			result = append(result, token)

			//Empty 'token'
			token = ""

		} else {

			//Populate 'token' with elements occuring before the separator
			token += string(str[i])

		}
	}

	//Append the elemwnts after the last seperator
	result = append(result, token)

	//Purge empty strings that may occur after spliting the string
	for i := 0; i < len(result); i++ {
		if result[i] == "" {

			//If the empty string does not occur at the end of the slice
			if i < len(result)-1 {
				result = append(result[:i], result[i+1:]...)
				i--
			} else {

				//If the empty string occurs at the end of the slice
				result = result[:i]

			}
		}
	}

	return result
}

// Converts an integer to a string
func Itoa(num int) string {
	var div int
	var char int32
	var mod int
	var result string
	var numSlc []int

	//Turn negative numbers to positive integers for seemless conversion to ASCII
	if num < 0 {
		num *= -1
	}

	//Return the string "0" if the integer parsed is 0
	if num == 0 {
		result = "0"
	}

	//Set 'div' to 'num' to parse it through the while loop below
	div = num

	//As long as 'div' is greater than zero, find and store the remainder of 'div' divided by 10
	for div > 0 {

		//Capture the last digit of the integer and update 'div' for the next round of divisions
		mod = div % 10
		div /= 10

		//Append the last digit to a slice of ints
		numSlc = append(numSlc, mod)
	}

	//Range from through 'numSlc', in reverse
	for i := len(numSlc) - 1; i >= 0; i-- {

		//Convert each digit to its rune representation
		char = int32(numSlc[i]) + '0'

		//Concat the casted runes to 'result'
		result += string(char)
	}
	return result
}
