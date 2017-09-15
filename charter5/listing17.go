package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user emial to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user
	level string
}

func (a *admin) notify() {
	fmt.Printf("admin %s<%s>\n", a.name, a.email)
}

func main() {
	ad := admin{
		user: user{
			name: "lisa",
			email: "lisa@emial.com",
		},
		level: "super",
	}

	// 给admin用户发送一个通知
	// 用于实现接口的内部类型的方法 被提升到了外部类型
	sendNotification(&ad)

	ad.user.notify()

	ad.notify()
}

func sendNotification(n notifier) {
	n.notify()
}