package main

import "fmt"

func CallFunc(f func()) {
	f()
}

// @desc 関数を引数に取る関数（コールバック関数）
func main() {
	CallFunc(func() {
		fmt.Println("I'm a function")
	})
}
