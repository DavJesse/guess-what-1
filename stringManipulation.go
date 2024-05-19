package main

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
		if i <= len(str)-len(sep) && str[i:i+len(sep)] == sep {
			result = append(result, token)
			token = ""
		} else {
			token += string(str[i])
		}
	}
	result = append(result, token)

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

	if num < 0 {
		num *= -1
	}

	if num == 0 {
		result = "0"
	}

	div = num

	for div > 0 {
		mod = div % 10
		div /= 10

		numSlc = append(numSlc, mod)
	}

	for i := len(numSlc) - 1; i >= 0; i-- {
		char = int32(numSlc[i]) + '0'
		result += string(char)
	}
	return result
}
