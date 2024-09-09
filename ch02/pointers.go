package main

import "fmt"

type aStructure struct {
	field1 complex128
	field2 int
}

func processPointer(x *float64) {
	*x = *x * *x
}

// Note: go is smart enough to know that temp is a point that escapes the function and therefore allocates to the heap memory, not the stack
func returnPointer(x float64) *float64 {
	temp := 2 * x
	return &temp
}

func bothPointers(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}


func main() {
	var f float64 = 12.123
	fmt.Println("Memory address of f:", &f)	// Memory address of f: <hex address 123>

	// Pointer to f
	fP := &f
	fmt.Println("Memory address of fP:", fP)	// Memory address of fP: <hex address 123>
	fmt.Println("Value of fP:", *fP)			// Value of fP: 12.123

	// The value of f changes
	processPointer(fP)
	fmt.Printf("Value of f: %.2f\n", f)		// Value of f: 146.97

	// The value of f does not change
	x := returnPointer(f)
	fmt.Printf("Value of x: %.2f\n", *x)		// Value of x: 293.93

	// The value of f does not change
	xx := bothPointers(fP)
	fmt.Printf("Value of xx: %.2f\n", *xx)	// Value of xx: 293.93

	// Check for empty structure
	var k *aStructure
	// This is nil because currently, k points to nowhere
	fmt.Println(k)		// <nil>

	// Therefore you are allowed to do this:
	if k == nil {
		k = new(aStructure)
	}

	fmt.Printf("%+v\n", k)		// &{field1:(0:0i) field2:0}
	// This may seem like a redundant check, but it is useful to avoid a nil pointer dereference which would crash the program
	if k != nil {
		fmt.Println("k is not nil!")
	}
}