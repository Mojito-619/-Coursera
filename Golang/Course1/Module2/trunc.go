// Packages
package main

// Imports
import "fmt"

// Main Function
func main() {

	// Declaring Main scope variables
	var user_scan float32
	var err error

	for i := 0; i < 2; i++ {
		fmt.Print("Please enter a floating point number: ")
		fmt.Scanf("%f \n", &user_scan)

		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Printf("Your float is: %f \n", user_scan)
			var trunc_float int = int(user_scan)
			fmt.Printf("Your truncated float is: %d \n", trunc_float)
		}
	}
}
