package main

import "os"

func printLn(str string) {
	os.Stdout.WriteString(str + "\n")
}
