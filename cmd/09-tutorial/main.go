package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
CHANNELS
	- Hold data
	- Thread safe
	- Listen for data
*/

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 12

func main() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)

	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)

	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}

	}
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofuPrice = rand.Float32() * 20
		if tofuPrice <= MAX_TOFU_PRICE {
			tofuChannel <- website
			break
		}

	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("\nFound a deal on chicken at %s ", website)
	case website := <-tofuChannel:
		fmt.Printf("\nFound a deal on tofu at %s ", website)
	}
}

/*
func main() {
	var c = make(chan int) //They only can hold one int value
	go process(c)

	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second * 1) //Waits one second before starting again
	}

	//EXAMPLE OF DEADLOCK
	//		c <- 1      //This is the way to assign values to a channel
	//		var i = <-c //The value in the channel is extracted and stored into a variable
	//		fmt.Println(i)

}


func process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}

}

*/
