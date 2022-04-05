/*
Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s = ½at^2 + vot + so

Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial
displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should
take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume
the following values for acceleration, initial velocity, and initial
displacement: a = 10, vo = 2, so = 1. I can use the
following statement to call GenDisplaceFn() to
generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to
print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print
the displacement after 5 seconds.

fmt.Println(fn(5))
*/

// Packages
package main

// Imports
import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Main Function
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Attempting to calculate displacement")
	fmt.Println("Please enter acceleration a: ")
	scanner.Scan()
	acceleration, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Println("Please enter initial velocity v0: ")
	scanner.Scan()
	initial_velocity, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Println("Please enter initial displacement s0: ")
	scanner.Scan()
	initial_displacement, _ := strconv.ParseFloat(scanner.Text(), 64)

	Displacement_function := GenDisplaceFn(acceleration, initial_velocity, initial_displacement)
	fmt.Println("Displacement model generated.")

	fmt.Println("Please enter travel time t: ")
	scanner.Scan()
	travel_time, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Printf("Displacement after %v seconds is %v meters.",
		travel_time, Displacement_function(travel_time))
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
	return fn
}
