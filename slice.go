package main

import "fmt"

/*
	切片是可索引的，并且可以用len()方法获取长度
	切片提供了计算容量的方法cap()可以测量切片最长可以达到多少
*/
// func main() {
// 	var numbers = make([]int, 3, 5)
// 	printSlice(numbers)
// }

// func printSlice(x []int) {
// 	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
// }

// 空(nil)切片

// func main() {
// 	var numbers []int
// 	printSlice(numbers)

// 	if(numbers == nil) {
// 		fmt.Printf("切片是空的")
// 	}
// }

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

// 切片截取

// func main() {
// 	numbers := [] int{0, 1, 2, 3, 4, 5, 6, 7, 8}
// 	printSlice(numbers)
// 	fmt.Println("numbers==", numbers)
// 	fmt.Println("numbers[1:4]==", numbers[1:4])
// 	fmt.Println("numbers[1:3]==", numbers[:3])
// 	fmt.Println("numbers[3:]==", numbers[3:])
// 	numbers_1 := numbers[1:4]
// 	printSlice(numbers_1)
// }

// append copy函数

func main() {
	var numbers []int
	printSlice(numbers)

	numbers = append(numbers, 0)
	printSlice(numbers)

	numbers = append(numbers, 1)
	printSlice(numbers)

	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	numbers_1 := make([] int, len(numbers), cap(numbers) * 2)

	copy(numbers_1, numbers)
	printSlice(numbers_1)
}


