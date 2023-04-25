package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
	// X, Y int
}

// 構造型のスライス
type Users []*User

/*
type Users struct {
	Users []+Users
}
*/

// 構造体の埋め込み
func main() {
	m := map[int]User{
		1: {Name: "user1", Age: 10},
		2: {Name: "user2", Age: 20},
	}

	fmt.Println(m)

	m2 := map[User]string{
		{Name: "user1", Age: 10}: "Tokyo",
		{Name: "user2", Age: 20}: "LA",
	}
	fmt.Println(m2)

	m3 := make(map[int]User)
	fmt.Println(m3)
	m3[0] = User{Name: "user3", Age: 30}
	fmt.Println(m3)

	for i, u := range m {
		fmt.Println(i, u)
	}
}
