package main

import "fmt"

// @desc 無名関数
func main() {
	f := func(x, y int) int {
		return x + y
	}
	i := f(1, 2)
	fmt.Println(i)

	i2 := func(x, y int) int {
		return x + y
	}(2, 3)
	fmt.Println(i2)
}
