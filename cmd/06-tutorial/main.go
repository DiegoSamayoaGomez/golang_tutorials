package main

import "fmt"

// An introduction to structs and methods

// STRUCT
// It´s like defining our own value type
type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

// METHOD
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

// INTERFACE
type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {

		fmt.Println("You can make it there")
	} else {
		fmt.Println("Need to fuel up first")
	}
}

func main() {

	//When we haven´t defined some mpg and gallons, al their values are 0
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15} //We can also ommit the field´s name and just assign their values (25, 15)
	var myEngine2 electricEngine = electricEngine{25, 15}
	fmt.Println(myEngine.mpg, myEngine.gallons)
	/*
		Other way to define values
		myEngine.mpg = 20
		myEngine.gallons = 15

	*/
	/*
		//ANONIMOUS STRUCTS -> The main difference is that these aren´t reusable

		var myEngine2 = struct {
			mpg     uint8
			gallons uint8
		}{25, 15}
	*/

	//Method call
	fmt.Printf("The total miles left in tank: %v", myEngine.milesLeft())

	canMakeIt(myEngine2, 50) //canMakeIt(myEngine, 50) -> GAS / ELECTRIC
}
