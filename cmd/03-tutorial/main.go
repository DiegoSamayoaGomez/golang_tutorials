package main

import (
	"errors"
	"fmt"
)

func main() {
	printValue := "I´m an argument"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 0

	var result, remainer, err = intDivision(numerator, denominator)
	/*
		if err != nil {
			fmt.Print(err.Error())
		} else if remainer == 0 {
			fmt.Printf("The result of the integer division is %v", result)
		} else {
			fmt.Printf("The result of the integer division is %v with remainder %v ", result, remainer)
		}
	*/
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainer == 0:
		fmt.Printf("The result of the integer division is %v ", result)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v ", result, remainer)
	}

}

func printMe(printValue string) {
	fmt.Println("Hi, I´m inside a function", printValue)
}

// The "int" outside the parenthesis means that the return value
// is gonna be an integer
// If there are more than one result value, we must put the return value types inside
// parenthesis (int, int)
func intDivision(numerator int, denominator int) (int, int, error) {

	//Error handling (Kinda)
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainer int = numerator % denominator
	return result, remainer, err
}
