// Packages
package main

// Imports
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Main Function
func main() {

	// Declaring Main scope variables
	search_result := "Not Found!"
	var err error

	// Because fmt.Scan separates strings into args with space, we can use bufio
	fmt.Print("Please enter a string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	// Make analysis case insensitive
	line = strings.ToLower(line)

	if err != nil {
		fmt.Println("Error: ", err)
	} else { // All 3 runes must be found for a "Found!" ouput
		if strings.HasPrefix(line, "i") &&
			strings.Contains(line, "a") &&
			strings.HasSuffix(line, "n") {
			search_result = "Found!"
		}
		fmt.Println(search_result)
	}

}
