// Packages
package main

// Imports
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Main Function
func main() {

	// Declaring Main scope variables
	slice1 := make([]int, 0, 3) //Empty slice of 3

	for i := 0; true; i++ {
		fmt.Print("Please enter an integer or X to quit: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		user_scan := scanner.Text()

		// Our only exit clause of upper case X
		if user_scan == "X" {
			return
		}

		// Converting our input into an int
		user_scan_int, err := strconv.Atoi(user_scan)
		if err != nil {
			fmt.Printf("WARNING. '%s' will be interpreted as 0: \n", user_scan)
		}

		// Add our new User Input to the slice
		slice1 = append(slice1, user_scan_int)

		// Sort the slice using a the stable func
		sort.SliceStable(slice1, func(i, j int) bool { return slice1[i] < slice1[j] })

		// Print the sorted slice
		fmt.Printf("Your sorted slice: %d \n", slice1)
	}
}
