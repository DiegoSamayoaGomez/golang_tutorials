package main

//A summary of some value types

import "fmt"

func main() {
	//INT
	var intNum int = 32767
	fmt.Println(intNum)

	//FLOAT
	//Must define wheter it´s 32 or 64
	var floatNum float64 = 123456789.9
	fmt.Println(floatNum)

	//Can´t combine value types, must connvert them first
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	// -> here's the conversion
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	//Division and remainer
	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	//STRING
	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)
	fmt.Println(len(myString)) //Prints the number in ASCII o BINARY IDK

	//BOOLEAN
	var myBooelan bool = true
	fmt.Println(myBooelan)

	//---SHORT VARIABLE DECLARATION---
	//GO can infer the value type if we declare them with some data

	//Notice that "var" isn´t necessary
	myVar := "Text"
	fmt.Println(myVar)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	//CONSTANTS
	const myConst string = "Const value"
	fmt.Println(myConst)
}
