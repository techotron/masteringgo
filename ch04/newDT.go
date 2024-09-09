package main

import (
	"fmt"
	"errors"
)

// TreeLast is a generic type of a slice of any type. When you instantiate a variable of this type, you need to 
//  specify the type of the elements in the slice, eg: myIntSlice := TreeLast[int]{1, 2, 3} or TreeLast[string]{"aa", "bb"}
type TreeLast[T any] []T

// This is a method which will operate of TreeLast variables.
func (t TreeLast[T]) replaceLast(element T) (TreeLast[T], error) {
	if len(t) == 0 {
		return t, errors.New("This is empty!")
	}

	t[len(t) -1] = element
	return t, nil
}

func main() {
	tempStr := TreeLast[string]{"aa", "bb"}
	fmt.Println(tempStr)
	tempStr.replaceLast("cc")
	fmt.Println(tempStr)

	tempInt := TreeLast[int]{12, -3}
	fmt.Println(tempInt)
	tempInt.replaceLast(0)
	fmt.Println(tempInt)
}

