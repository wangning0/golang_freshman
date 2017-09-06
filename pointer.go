package main

import "fmt"

// func main() {
// 	var a int = 10
// 	fmt.Printf("变量的地址为：%x", &a)
// }

// 如何使用指针

// func main() {
// 	var a int = 20
// 	var ip *int
// 	ip = &a

// 	fmt.Printf("a 变量的地址是: %x\n", &a  )

//    /* 指针变量的存储地址 */
//    fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

//    /* 使用指针访问值 */
//    fmt.Printf("*ip 变量的值: %d\n", *ip )
// }

// 指针数组

// const MAX int = 3

// func main() {
// 	a := [] int{ 10, 100, 200 }
// 	var i int
// 	var ptr[MAX] *int

// 	for i = 0; i < MAX; i++ {
// 		ptr[i] = &a[i]
// 	}

// 	for i = 0; i < MAX; i++ {
// 		fmt.Printf("a[%d] = %d\n", i,*ptr[i] )
// 	}
// }

// 指向指针的指针

// func main() {
// 	var i int = 10
// 	var ptr *int
// 	var ptr_1 **int
// 	ptr = &i
// 	ptr_1 = &ptr
// 	fmt.Printf("i = %d\n", i)
// 	fmt.Printf("*i = %d\n", *ptr)
// 	fmt.Printf("**i = %d\n", **ptr_1)
// }

// 指针作为函数参数
func main() {
   /* 定义局部变量 */
   var a int = 100
   var b int= 200

   fmt.Printf("交换前 a 的值 : %d\n", a )
   fmt.Printf("交换前 b 的值 : %d\n", b )

   /* 调用函数用于交换值
   * &a 指向 a 变量的地址
   * &b 指向 b 变量的地址
   */
   swap(&a, &b);

   fmt.Printf("交换后 a 的值 : %d\n", a )
   fmt.Printf("交换后 b 的值 : %d\n", b )
}

func swap(x *int, y *int) {
   var temp int
   temp = *x    /* 保存 x 地址的值 */
   *x = *y      /* 将 y 赋值给 x */
   *y = temp    /* 将 temp 赋值给 y */
}