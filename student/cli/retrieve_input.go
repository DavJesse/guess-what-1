package cli

import (
	"bufio"
	"log"
	"os"
)

// Extract user input from standaord input
func ExtractInput() string {
	var result string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		result := scanner.Text()
		if result != "" {
			return result
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
