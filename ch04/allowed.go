package main

import "fmt"

func Same[T comparable](a, b T) bool {
	// Or
	// return a == b
	if a == b {
		return true
	}
	return false
}

func main() {
	fmt.Println(Same(10, 10)) // true
	fmt.Println(Same(10, 20)) // false
	fmt.Println(Same("A", "A")) // true
	fmt.Println(Same("A", "B")) // false
	// The following would case a compile error
	// fmt.Println(Same(10, "B"))
	// fmt.Println(Same([]int{1,2}, []int{1,3}))
}