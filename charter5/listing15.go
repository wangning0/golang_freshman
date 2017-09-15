package main

import "fmt"

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

func main() {
	ad := admin{
		user: user{
			name: "lisa",
			email: "lisa@emial.com",
		},
		level: "super",
	}

	ad.user.notify()

	// 内部类型的方法也被提升到了外部类型
	ad.notify()
}