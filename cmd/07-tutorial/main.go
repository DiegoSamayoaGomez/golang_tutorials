package main

import "fmt"

// POINTERS
//Variables wich stores memory locations rather than values

func main() {
	// "= new(int32) it´s used to asign a memory space.
	//If we don´t use it, we´ll get a null pointer exception"
	var p *int32 = new(int32)
	var i int32

	fmt.Printf("The value of p points to %v\n", *p)
	fmt.Printf("The value of v is %v\n", i)

	//We can create a pointer from the address of another variable by using &
	p = &i //Now p and i share the same memory addres, hence if we change the value of p or i both of them will share the same value

	*p = 10 //By using * p now points to a value of 10

	fmt.Printf("The value of p points to %v\n", *p)
	fmt.Printf("The value of v is %v\n", i)

	//Pointers and functions
	fmt.Println("USING POINTERS WITH FUNCTIONS")
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("The memory location of the thing1 array is: %p\n", &thing1) //We change %v to %p so it can format a pointer
	var result [5]float64 = square(thing1)                                  //We send thing1 as argument to the function and save the return value into result expecting an array of 5
	//with pointers
	/*
		var result [5]float64 = square(&thing1)
	*/
	fmt.Printf("The result is: %v\n", result)

	//It creates 2 different arrays, so we can modify thing1 without affecting thing2
	fmt.Printf("The array of thing 1 is: %v ", thing1)
}

func square(thing2 [5]float64) [5]float64 { //array de 5 valores
	fmt.Printf("The memory location of the thing1 array is: %p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i] //its a multiplication, not a pointer
	}
	return thing2
}

//SAME FUNCTIONS WITH POINTERS

/*
func square(thing2 *[5]float64) [5]float64 { //array de 5 valores
	fmt.Printf("The memory location of the thing1 array is: %p\n", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i] //its a multiplication, not a pointer
	}
	return *thing2
}
*/
