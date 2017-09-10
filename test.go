package main

import "fmt"

func main() {
	people := map[string]int{
		"wn":  21,
		"csr": 20,
	}

	for key, value := range people {
		fmt.Printf("name: %s, age: %d\n", key, value)
	}
	removePeople(people, "wn")
	for key, value := range people {
		fmt.Printf("name: %s, age: %d\n", key, value)
	}
}

func removePeople(people map[string]int, key string) {
	delete(people, key)
}
