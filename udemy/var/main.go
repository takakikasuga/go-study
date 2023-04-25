package main

import "fmt"

// i5 := 500 syntax error: non-declaration statement outside function body
var i5 int = 500 // 明示的に定義する場合は、関数の外で定義できる。

func outer() {
	var s3 string = "outer duplicate"
	var s4 string = "outer"
	fmt.Println(s3)
	fmt.Println(s4)
}

func main() {
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Go Lang"
	)
	fmt.Println(i2, s2)

	var i3 int    // 初期値 = 0
	var s3 string // 初期値 = ""（空文字）
	fmt.Println(i3, s3)
	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	// 暗黙的な定義 ⇨ 変数名 := 値
	i4 := 400
	fmt.Println(i4)
	i4 = 450
	fmt.Println(i4)
	// i4 = "error" cannot use "error" (untyped string constant) as int value in assignment
	// fmt.Println(i4)

	fmt.Println(i5)

	outer()

	var s5 string = "Not Use"
	fmt.Println(s5)
}
