package main

import (
	"fmt"
)

// 構造体 = 値型
type T struct {
	// User User
	User // 省略可能
}

type User struct {
	Name string
	Age  int
	// X, Y int
}

func (u *User) SetName() {
	u.Name = "A"
}

// 構造体の埋め込み
func main() {
	t := T{User: User{Name: "user1", Age: 10}}
	fmt.Println(t)
	fmt.Println(t.User)
	fmt.Println(t.User.Name)
	fmt.Println(t.Name) // 構造体の省略記法を使わないとこうなる。 ⇨ t.Name undefined (type T has no field or method Name)

	t.SetName()
	fmt.Println(t)
}
