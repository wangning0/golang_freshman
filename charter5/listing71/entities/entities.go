package entities

type User struct {
	Name string
	// email 小写 外包不能直接访问
	email string
}