package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 分配一个逻辑处理器给调度器用
	runtime.GOMAXPROCS(1)

	// wg用来等待程序完成，计数加2，表示要等待两个goroutine

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// 声明一个匿名函数，并创建一个goroutine
	go func() {
		// 在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a' + 26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()
	
	go func() {
		// 在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char <'A' + 26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	fmt.Println("Waiting to Finish")
	wg.Wait()

	fmt.Println("程序终止")
}