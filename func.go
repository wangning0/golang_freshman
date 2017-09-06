package main

import (
	"fmt"
)

// func main() {
// 	var r int
// 	r = max(1,2)
// 	fmt.Println(r)

// 	a, b := swap("hello", "world")
// 	fmt.Println(a, b)
// }

// // max函数可以放在调用者的前后
// func max(num1, num2 int) int {
// 	var result int

// 	if(num1 > num2) {
// 		result = num1
// 	} else {
// 		result = num2
// 	}

// 	return result
// }

// // 返回多个值的函数
// func swap(x, y string) (string, string) {
// 	return y, x
// }


/*
	函数值传递
*/

// func main() {
// 	// 定义局部变量
// 	var a int = 100
// 	var b int = 200
// 	fmt.Printf("交换前的值\n")
// 	fmt.Printf("a:%d, b:%d\n", a, b)
// 	swap(a, b)
// 	fmt.Printf("交换后的值\n")
// 	fmt.Printf("a:%d, b:%d\n", a, b)
// }

// func swap(x, y int) int {
// 	var temp int
// 	temp = x
// 	x = y
// 	y = temp

// 	return temp
// }

/*
	引用传递
*/

// func main() {
// 	// 定义局部变量
// 	var a int = 100
// 	var b int = 200
// 	fmt.Printf("交换前的值\n")
// 	fmt.Printf("a:%d, b:%d\n", a, b)
// 	swap(&a, &b)
// 	fmt.Printf("交换后的值\n")
// 	fmt.Printf("a:%d, b:%d\n", a, b)
// }

// func swap(x *int, y *int) {
// 	var temp int
// 	temp = *x
// 	*x = *y
// 	*y = temp
// }

/*
	Go语言函数作为值
*/

// func main() {
// 	getSquareRoot := func(x float64) float64 {
// 		return math.Sqrt(x)
// 	}

// 	fmt.Println(getSquareRoot(9))
// }

/*
	Go语言函数闭包
	Go支持匿名函数，可作为闭包，匿名函数是一个“内联”语句或表达式
	匿名函数的优越性在于可以直接使用函数内的变量，不必申明
*/

// func getSequence() func() int {
// 	i := 0
// 	return func() int {
// 		i += 1
// 		return i
// 	}
// }

// func main() {
// 	nextNumber := getSequence()
// 	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
// 	fmt.Println(nextNumber())
// 	fmt.Println(nextNumber())
// 	fmt.Println(nextNumber())
	
// 	/* 创建新的函数 nextNumber1，并查看结果 */
// 	nextNumber1 := getSequence()  
// 	fmt.Println(nextNumber1())
// 	fmt.Println(nextNumber1())
// }

/*
	Go语言函数方法

	Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针
*/

type Circle struct {
	radius float64
}

func main() {
	var c Circle
	c.radius = 10.00
	fmt.Println("Area of Circle(c) = ", c.getArea())
}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}