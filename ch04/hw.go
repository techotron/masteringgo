package main

import "fmt"

// Example of using generics with slices. PrintSlice will accept a slice of any type.
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	PrintSlice([]int{1, 2, 3, 4, 5})
	PrintSlice([]string{"one", "two", "three", "four", "five"})
	PrintSlice([]float64{1.1, 2.2, 3.3, 4.4, 5.5})
}