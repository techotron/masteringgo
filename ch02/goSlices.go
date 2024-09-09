package main

import "fmt"

func main() {
	aSlice := []float64{}
	fmt.Println(aSlice, len(aSlice), cap(aSlice))
	aSlice = append(aSlice, 1234.56)
	aSlice = append(aSlice, -34.0)
	fmt.Println(aSlice, len(aSlice), cap(aSlice))

	t := make([]int, 4)
	t[0] = -1
	t[1] = -2
	t[2] = -3
	t[3] = -4
	fmt.Println(fmt.Sprintf("slice: %v, len: %v, cap: %v", t, len(t), cap(t)))
	t = append(t, -5)
	fmt.Println(fmt.Sprintf("slice: %v, len: %v, cap: %v", t, len(t), cap(t)))
}