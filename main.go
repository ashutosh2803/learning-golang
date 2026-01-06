package main

import "fmt"

// package level variable
var name string = "Ashutosh"

func main() {
	fmt.Println("Hello World!")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, Cruel World!"
	i = 7

	fmt.Println(whatToSay)
	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()
	fmt.Println("The function returned", whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "Something", "else"
}
