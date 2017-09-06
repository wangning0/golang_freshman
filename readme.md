# GO

## Go 语言数据类型

* 布尔型

布尔型的值只可以是常量true/false，`var b bool = true`

* 数字类型

整型int和浮点型float,Go语言支持整型和浮点型数组，并且原生支持复数，位的运算采用补码

* 字符串类型

字符串就是一串固定长度的字符连接起来的字符序列，GO的字符串是单个字节连接起来的，采用UTF-8编码标识Unicode文本

* 派生类型

    * 指针类型 Pointer
    * 数组类型
    * 结构化类型 struct
    * Channel类型
    * 函数类型
    * 切片类型
    * 接口类型 interface
    * Map类型

## Go 语言变量

Go语言变量名由字母、数字、下划线组成，其中首个字母不能为数字

`var identifier type`

**变量声明**

* 指定变量类型，声明后若不赋值，使用默认值

```
var v_name v_type
v_name = value
```

* 根据值自行判定变量类型

```
var v_name = value
```

* 省略var 

```
v_name := value // := 左侧的便利那个不应该是已经声明过的，否则会导致编译错误
// 例如
var a int = 10
var b = 10
c := 10
```

**多变量声明**

```
// 类型相同多个变量，非全局变量
var vname1, vname2, vname3 type
vname1, vname2, vname3 = v1, v2, v3

var vname1, vname2, vname3 = v1, v2, v3

vname1, vname2, vname3 := v1, v2, v3

// 这种因式分解关键字的写法一般用于声明全局变量
var (
    vname1 v_type1
    vname2 v_type2
)
```

## Go语言常量

常量是一个简单值的标识符，在程序运行时，不会被修改的量

常量中的数据类型只可以是布尔、数字和字符串型

```
const identifier [type] = value
```

## Go语言条件语句

* if语句

* if...else语句

* if嵌套语句

* switch语句

* select语句

    select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。

## Go语言函数

Go语言最少有个main()函数

```
func function_name([paramter list]) [return_types] {
    // 函数体
}
```

```
func max(num1, num2 int) int {
	var result int

	if(num1 > num2) {
		result = num1
	} else {
		result = num2
	}

	return result
}
```

* 函数返回多个值

```
func swap(x, y string) (string, string) {
	return y, x
}
```

* 函数参数

函数如果使用参数，该变量可称为函数的形参
形参就像定义在函数体内的局部变量
调用函数，可以通过两种方式来传递参数

    * 值传递
    * 引用传递

默认情况下，Go语言使用的是值传递，在调用过程中不会影响到实际参数


* Go语言函数方法

Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针

```
func (variable_name variable_data_type) func_name() [return_type] {
    // 函数体
}
```

## Go语言变量作用域

Go语言中变量可以在三个地方声明

* 函数内定义的变量成为局部变量
* 函数外定义的变量称为全局变量
* 函数定义中的变量称为形式参数

## Go语言数组

Go语言提供了数组类型的数据结构
数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整形、字符串或自定义类型

```
// 数组的声明方式

var variable_name [SIZE] variable_type

// 实例
var balance [10] float32
```


**初始化数组**

`var balance = [5]float32{100.0, 2.0, 3.4, 7.0m 50.0}`

初始化数组中{}中的元素个数不能大于[]中的数字
如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：
`var balance = [...]float32{100.0, 2.0, 3.4, 7.0m 50.0}`

**访问数组**

```
func main() {
	var n [10]int
	var i, j int

	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}
```
**多维数组**

`var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type`

```
func main() {
	var a = [5][2]int { {0, 0}, {1, 2}, {3, 6}, {4, 8}}
	var i, j int

	for i = 0 ; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
		}
	}
}
```

**Go语言向函数传递数组**

如果你想向函数传递数组参数，需要在函数定义时，声明形参为数组，两种方式来声明：

* 方式一 形参设定数组大小

```
void myFunction(param [10]int) {
    // ...
}
```
* 方式二 形参未设定数组大小

```
void myFunction(param []int) {
    // ...
}
```

## Go语言指针

Go语言中使用指针可以更简单的执行一些任务

变量是一种使用方便的占位符，用于引用计算机内存地址

Go语言的取地址符是& 放到一个变量前使用就会返回相应变量的内存地址

```
func main() {
	var a int = 10
	fmt.Printf("变量的地址为：%x", &a)
}
```

**什么是指针**

一个指针变量可以指向任何一个值的内存地址，它指向那个值的内存地址

`var var_name *var_type`

`var ip *int`

`var fp *float32`

**如何使用指针**

* 定义指针变量
* 为指针变量赋值
* 访问指针变量中指向地址的值

**Go空指针**

当一个指针被定义后没有分配任何变量时，它的值为nil

nil指针也被称为空指针， nil都指代零值或空值

## Go语言结构体

Go语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型

**定义结构体**

结构体定义需要使用type和struct语句，struct语句定义了一个新的数据类型，结构体中有一个或多个成员

```
type struct_variable_type struct {
    memeber definition;
    memeber definition;
    ...
    memeber definition;
}
```

一旦定义了结构体类型，他就能用于变量的声明，语法格式如下

`variable_name := structire_variable_type {value1, value2...valuen}`

```
type Books struct {
	title string
	author string
	subject string
	book_id string
}
```

**结构体作为函数参数**

```
func printBook(book Books) {
	fmt.Printf( "Book title : %s\n", book.title);
   	fmt.Printf( "Book author : %s\n", book.author);
   	fmt.Printf( "Book subject : %s\n", book.subject);
   	fmt.Printf( "Book book_id : %d\n", book.book_id);
}
```

**结构体指针**

可以定义指向结构体的指针类似于其他指针变量

`var struct_pointer *Books`

`struct_pointer = &Book1`

## Go语言切片(Slice)

Go语言切片是对数组的抽象

Go数组的长度不可改变，在特定场景中这样的集合就不太实用，Go中提供了一种灵活，功能强悍的内置类型切片（"动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大

`var identifier []type`

切片不需要说明长度，或者使用`make()`函数来创建切片

```
var slice1 []type = make([]type, len)

// or

slice1 := make([]type, len)

// 也可以指定容量
make([]T, length, capacity)
```

**len()和cap()函数**

切片是可索引的，并且可以由 len() 方法获取长度。
切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。

```
package main

import "fmt"

func main() {
   var numbers = make([]int,3,5)

   printSlice(numbers)
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```

**nil切片**

```
func main() {
	var numbers []int
	printSlice(numbers)

	if(numbers == nil) {
		fmt.Printf("切片是空的")
	}
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```

**切片截取**

可以通过设置下限及上限来设置截取切片

```
func main() {
	numbers := [] int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)
	fmt.Println("numbers==", numbers)
	fmt.Println("numbers[1:4]==", numbers[1:4])
	fmt.Println("numbers[1:3]==", numbers[:3])
	fmt.Println("numbers[3:]==", numbers[3:])
	numbers_1 := numbers[1:4]
	printSlice(numbers_1)
}
```

**append()和copy()函数**

如果想增加切片的容量，我们必须创建一个新的更大的切片把原分片的内容都拷贝过来


```
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
```

## Go语言范围(Range)