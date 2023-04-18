package main

import (
	"fmt"
)

// @desc 論理演算子
func main() {
	fmt.Println(true && false == true)
	fmt.Println(true && true == true)
	fmt.Println(true && false == false)

	fmt.Println(true || false == true)
	fmt.Println(false || false == true)

	fmt.Println(!true)
	fmt.Println(!false)
}
