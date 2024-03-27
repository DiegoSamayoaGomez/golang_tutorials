package main

import (
	"fmt"
)

// A simple introduction to arrays, slice, maps and loop
func main() {
	/*
	   ARRAY
	   - Fixed length (cant´t change it later)
	   - Same type (can´t mix int and float e.g)
	   - Indexable (can be accessed any of its parts)
	   - Contiguous in memory (related with pointers)
	*/
	// var intArr [3]int32 //Normal array declaration
	//var intArr [3]int32 = [3]int32{1, 2, 3} //Auto initialize an array
	intArr := [3]int32{1, 2, 3} //The shortest way to write an array

	//intArr[1] = 123 //Change a value in that position, not resize it
	fmt.Println(intArr)
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	/*
		SLICES
		Similar or based in arrays, but more powerful
		- Allows to be resized
	*/
	fmt.Println("SLICE ->")
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Printf("The lenght is %v with capacity %v", len(intSlice), cap(intSlice)) //cap -> capacity
	//Append values (increase size of the slice by adding new items)
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("The lenght is %v with capacity %v", len(intSlice), cap(intSlice)) //cap -> capacity

	//Append data from other slice
	intSlice2 := []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	//Declare the slice´s length and it´s capacity by using MAKE
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	/*
		MAPS
		Also known as glossaries*/

	//String: uint -> age:8
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Diego": 23, "Daniel": 14}
	fmt.Println(myMap2)
	fmt.Println(myMap["Adam"])

	/*
			LOOPS
		Well, loop because in GO there only exists the for loop
	*/
	for i := 0; i < 5; i++ {
		fmt.Println("This is a loop")
	}
	for name, age := range myMap2 {
		fmt.Printf("Name %v , Age %v ", name, age)
	}
}
