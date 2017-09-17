package main

import (
	"fmt"
	"./entities"
)

func main() {
	u := entities.User{
		Name: "Bill",
		// email是非公开的字段，会报错
		email: "bill@email.com",
	}

	fmt.Printf("%v", u)
}