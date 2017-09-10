package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	for k, v := range slice {
		fmt.Printf("index: %d, value: %d\n", k, v)
	}
}