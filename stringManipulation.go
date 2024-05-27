package main

func hasSuffix(str, subStr string) bool {
	var status bool
	diff := len(str) - len(subStr)

	// Check if string starts with substring
	if str[diff:] == subStr {
		status = true
	}

	return status
}

func splitString(str string, sep string) []string {
	var result []string
	var token string

	for i := 0; i < len(str); i++ {
		if i <= len(str)-len(sep) && str[i:i+len(sep)] == sep {

			// Append 'token' to result
			result = append(result, token)

			// Empty 'token'
			token = ""

		} else {
			// Populate 'token' with elements occuring before the separator
			token += string(str[i])
		}
	}

	// Append the elemwnts after the last seperator
	result = append(result, token)

	// Purge empty strings that may occur after spliting the string
	for i := 0; i < len(result); i++ {
		if result[i] == "" {
			if i < len(result)-1 {
				result = append(result[:i], result[i+1:]...)
				i--
			} else {
				result = result[:i]
			}
		}
	}

	return result
}

func Itoa(num int) string {
	var div int
	var char int32
	var mod int
	var result string
	var numSlc []int

	// Turn negative numbers to positive integers for seemless conversion to ASCII
	if num < 0 {
		num *= -1
	}

	// Return the string "0" if the integer parsed is 0
	if num == 0 {
		result = "0"
	}

	div = num

	for div > 0 {
		mod = div % 10
		div /= 10

		// Append the last digit to a slice of ints
		numSlc = append(numSlc, mod)
	}

	for i := len(numSlc) - 1; i >= 0; i-- {

		// Convert each digit to its rune representation and concat to result
		char = int32(numSlc[i]) + '0'
		result += string(char)
	}
	return result
}
