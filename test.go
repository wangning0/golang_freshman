package main

import (
	"fmt"
	// "bytes"
)

func main() {
	readBytes := make([]byte, 2)
	bs := []byte{1, 2, 3}
	fmt.Println(readBytes, bs)
}