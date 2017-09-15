package main

import (
	"fmt"
	"./entities"
)

func main() {
	u := entities.User{
		Name: "Bill",
		// charter5/listing71/listing71.go:11:8: unknown field 'email' i
// n struct literal of type entities.User
		email: "bill@email.com",
	}

	fmt.Printf("%v", u)
}