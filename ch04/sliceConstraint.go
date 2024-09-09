package main

import "fmt"

// All 3 generic functions are equivalent but have a variation in syntax. 


// Define a type called S which is a slice of any subtype of E. E is an interface (any) type here.
func f1[S interface{ ~[]E }, E interface{}](x S) int {
	return len(x)
}

func f2[S ~[]E, E interface{}](x S) int {
	return len(x)
}

func f3[S ~[]E, E any](x S) int {
	return len(x)
}


func main() {
	fmt.Println(f1([]int{1,2,3,4}))
	fmt.Println(f2([]int{1,2,3,4}))
	fmt.Println(f3([]int{1,2,3,4}))
}