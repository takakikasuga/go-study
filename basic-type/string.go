package main

import "fmt"

// @desc 文字列型
func main() {
	var s string = "Hello Golang"

	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var s1 string = "300"
	fmt.Println(s1)
	fmt.Printf("%T\n", s1)

	fmt.Println(`test
	test
	   test`)

	fmt.Println("\"")
	fmt.Println(`"`)

	fmt.Println(s[0])
	fmt.Println(string(s[0]))
}
