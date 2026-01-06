package main

import (
	"log"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
	BirthDate time.Time
}

func main() {
	createUser()
}

func createUser() {
	user := User{
		FirstName: "Ashutosh",
		LastName:  "Kumar",
		Age:       27,
	}
	log.Println(user.FirstName, user.LastName, user.Age, user.BirthDate)
}
