/*
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing. The BubbleSort()
function should modify the slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.
*/

// Packages
package main

// Imports
import (
	"fmt"
)

// Main Function
func main() {
	arr := make([]int, 10)
	fmt.Printf("Enter 10 integers value: ")
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Printf("Your input array was %v\n", arr)
	BubbleSort(arr)
	fmt.Printf("Your sorted array is %v\n", arr)
}

func Swap(slice []int, index int) {
	// Swaps slice[index] with slice[index+1]
	buffer := slice[index]
	slice[index] = slice[index+1]
	slice[index+1] = buffer
}

func BubbleSort(slice []int) {
	// Bubble sorts the input. It does not return anything as it directly modifies
	// the argument as it passed as a pointer

	//The inner j loop will iterate through the whole slice
	//The outer i loop will decrease the scope until the last index
	slice_length := len(slice)
	for i := 0; i < slice_length-1; i++ {
		for j := 0; j < slice_length-i-1; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}
	}
}
