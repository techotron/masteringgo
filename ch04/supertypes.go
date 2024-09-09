package main

import "fmt"

type AnotherInt int

// Support all types which are subtypes of int
type AllInts interface {
	~int
}

func AddElements[T AllInts](s []T) T {
	sum := T(0)
	for _, v := range s {
		sum = sum + v
	} 
	return sum
}

func main() {
	fmt.Println(AddElements([]AnotherInt{1,2,3,4}))
}