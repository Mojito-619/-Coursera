// Packages
package main

// Imports
import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Main Function
func main() {

	// Declaring Main scope variables
	person := make(map[string]string)

	fmt.Print("Please enter your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name_scan := scanner.Text()
	person["name"] = name_scan

	fmt.Print("Please enter your adress: ")
	scanner.Scan()
	address_scan := scanner.Text()
	person["address"] = address_scan

	json_person, err := json.Marshal(person)
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
	} else {
		fmt.Println(string(json_person))
	}
}
