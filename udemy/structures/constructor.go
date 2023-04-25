package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
	// X, Y int
}

// constructor関数
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

// 構造体の埋め込み
func main() {
	user1 := NewUser("user1", 10)
	fmt.Println(user1)
	fmt.Println(*user1)
}
