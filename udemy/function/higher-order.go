package main

import "fmt"

func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

// @desc 関数を返す関数（高階関数）
func main() {
	f := ReturnFunc()
	f()
}
