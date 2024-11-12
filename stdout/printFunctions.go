package stdout

import "os"

// Prints every string parsed
func PrintLn(str string) {
	os.Stdout.WriteString(str + "\n")
}

// Prints every strings parsed
func PrintF(str1, str2 string) {
	os.Stdout.WriteString(str1 + str2 + "\n")
}
