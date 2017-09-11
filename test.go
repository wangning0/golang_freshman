package main

import "fmt"

func main() {
	type Duration int64
	
	var dur Duration
	dur = int64(100)
	fmt.Println(dur)
}
