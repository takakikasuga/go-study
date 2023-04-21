package main

import (
	"fmt"
)

// @desc スライス（copy）
func main() {
	sl := []int{1, 2, 3}
	sl2 := sl
	sl2[0] = 1000
	fmt.Println("sl", sl)
	fmt.Println("sl2", sl2)

	var i int = 10
	i2 := 100
	fmt.Println("i", i)
	fmt.Println("i2", i2)

	sl3 := []int{1, 2, 3, 4, 5}
	sl4 := make([]int, 5, 10)
	fmt.Println("sl4", sl4)
	n := copy(sl4, sl3) // n = copyに成功した数
	fmt.Println("n", n)
	fmt.Println("sl4", sl4)
}
