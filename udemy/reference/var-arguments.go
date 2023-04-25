package main

import "fmt"

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

// @desc スライス（可変長引数）
func main() {
	fmt.Println(Sum(1, 2, 3, 4, 5))
	fmt.Println(Sum())

	sl := []int{1, 2, 3}
	fmt.Println(Sum(sl...))
}
