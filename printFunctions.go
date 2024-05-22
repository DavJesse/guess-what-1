package main

import "os"

// Prints every sting parsed
func printLn(str string) {
	os.Stdout.WriteString(str + "\n")
}
