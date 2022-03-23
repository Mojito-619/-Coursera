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
	// Declaring  variables
	type person struct {
		fname string
		lname string
	}

	person_slice := make([]person, 0)

	// Obtain desired filename
	fmt.Print("Name of file you want to read (including filetype): ")
	filename_scanner := bufio.NewScanner(os.Stdin)
	filename_scanner.Scan()
	filename := filename_scanner.Text()

	// Opens selected file
	f_object, err := os.Open(filename)
	check(err)
	defer f_object.Close()

	// Scan each line of the file iteratively
	f_scanner := bufio.NewScanner(f_object)
	for f_scanner.Scan() {
		// Separate both fname and lname into separate fields
		full_name := strings.Fields(f_scanner.Text())
		// Ensure we don't go over 20 chars
		fname := truncString(full_name[0], 20)
		lname := truncString(full_name[1], 20)
		// Make the individual struct and append it to the slice
		person_slice = append(person_slice, person{fname, lname})
	}
	for _, x := range person_slice {
		fmt.Println(x)
	}
}

// Function to ensure 20 char max string
func truncString(input_string string, max_char int) string {
	if len(input_string) > max_char {
		return string(input_string[0:max_char])
	} else {
		return input_string
	}
}

// Function for error management
func check(err error) {
	if err != nil {
		panic(err)
	}
}
