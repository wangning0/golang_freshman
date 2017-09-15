package main

import (
	"fmt"
	"./entities"
)

func main() {
	a := entities.Admin{
		Rights: 10,
	}

	a.Name = "wingerwang"
	a.Email = "wingerwang@qq.com"

	fmt.Printf("User: %v", a)
}