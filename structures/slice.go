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
	user1 := User{"user1", 10}
	user2 := User{"user2", 20}
	user3 := User{"user3", 30}
	user4 := User{"user4", 40}

	users := Users{}
	users = append(users, &user1)
	users = append(users, &user2)
	users = append(users, &user3, &user4)
	fmt.Println(users)

	for i, u := range users {
		fmt.Println(i, *u)
	}

	users2 := make([]*User, 0)
	fmt.Println(users2)
	users2 = append(users2, &user1, &user2)
	for i, u := range users2 {
		fmt.Println(i, *u)
	}
}
