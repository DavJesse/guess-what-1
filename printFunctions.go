package main

import "os"

// Prints every string parsed
func printLn(str string) {
	os.Stdout.WriteString(str + "\n")
}

// Prints every strings parsed
func printF(str1, str2 string) {
	os.Stdout.WriteString(str1 + str2 + "\n")
}
