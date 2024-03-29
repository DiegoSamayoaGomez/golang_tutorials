package main

import (
	"fmt"
)

/*
	GENERICS
	Used to make "templates" I guess, use slices to cover many cases without repeating so much code
*/

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice[int](intSlice))

	var float32Slice = []float32{1, 2, 3}
	fmt.Println(sumSlice[float32](float32Slice))
}
func sumSlice[T int | float32 | float64](slice []T) T { //Here T is passed as a Type, but it covers the value type of int/float32/float64
	var sum T //Type that could be many values
	for _, v := range slice {
		sum += v
	}
	return sum
}
