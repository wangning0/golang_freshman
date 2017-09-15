package main

import (
	"fmt"
)

// notifier 是一个定义了通知类行为的接口
type notifier interface {
	notify()
}

// user在程序里定义一个用户类型
type user struct {
	email string
	name string
}

// notify 是使用指针接收者实现的方法
func (u *user) notify() {
	fmt.Printf("Sending user emial to %s<%s>\n", u.name, u.email)
}

func main() {
	u := user{"Bill", "bill@email.com"}
	sendNotification(&u)
}

func sendNotification(n notifier) {
	n.notify()
}