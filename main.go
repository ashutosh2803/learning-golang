package main

import (
	"fmt"
	"log"
)

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

	var color string
	color = "Green"
	log.Println("color is set to", color)
	changeUsingPointer(&color)
	log.Println("after func call, color is set to", color)
}

func saySomething() (string, string) {
	return "Something", "else"
}

func changeUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}
