package main

import (
	"fmt"
)

type user struct {
	name string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	// user类型的值可以用来调用
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	// 指向user类型值的指针也可以用来调用
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	bill.changeEmail("bill@newemail.com")
	bill.notify()

	lisa.changeEmail("lisa@newemail.com")
	lisa.notify()
}