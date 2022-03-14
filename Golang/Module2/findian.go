// Packages
package main

// Imports
import (
	"fmt"
	"strings"
)

// Main Function
func main() {

	// Declaring Main scope variables
	var user_scan string
	search_result := "Not Found!"
	var err error

	fmt.Print("Please enter a string: ")
	fmt.Scanf("%s \n", &user_scan)

	if err != nil {
		fmt.Println("Error: ", err)
	} else { // All 3 runes must be found for a "Found!" ouput
		if strings.Contains(user_scan, "i") &&
			strings.Contains(user_scan, "a") &&
			strings.Contains(user_scan, "n") {
			search_result = "Found!"
		}
		fmt.Println(search_result)
	}

}
