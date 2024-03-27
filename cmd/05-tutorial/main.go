package main

// Some simple examples of how strings works in GO

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	var indexed = myString[0]

	fmt.Println(myString)
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		//It prints some weird number
		//That´s because GO uses UTF-8 to represent strings
		//Some old programming languages uses ASCII
		fmt.Println(i, v)
	}

	//RUNES
	//Its an equivalent of int32, it has more memory space but it knows how to dinamically allocate it
	fmt.Println("Rune")
	var myRune = []rune("résumé")
	fmt.Println(myRune)

	//STRING CONCATENATION
	// We can use GO´s string builder to make it easier
	var strSlice = []string{"D", "I", "E", "G", "O"}
	var strBuilder strings.Builder //Imported strings library
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v\n", catStr)

}
