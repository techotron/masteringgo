package main

import (
	"fmt"
	"os"
	"strconv"
)

// Note: the slices package now has a Delete method (since 1.21)
// 2 techniques for deleting from a slice:
// 1. split the slice into 2 at the given index (excluding it)
// 2. copy the last element of the slice to the index and truncate the slice by 1
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need an integer value.")
		return
	}
	index := arguments[1]
	i, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Using index", i)
	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	bSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Original slice:", aSlice)

	if i > len(aSlice) {
		fmt.Println("Cannot delete element", i)
		return
	}

	// Technique 1
	// The ... operator auto expands the slice into individual elements so that it's elements
	// can be appended to aSlice[:1] one by one.
	aaSlice := append(aSlice[:i], aSlice[i+1:]...)
	fmt.Println("After 1st deletion:", aaSlice)

	// Technique 2
	bSlice[i] = bSlice[len(bSlice)-1]
	bSlice = bSlice[:len(bSlice)-1]
	fmt.Println("After 2nd deletion:", bSlice)
}