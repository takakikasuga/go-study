package main

import (
	"fmt"
)

type MyInt int

func (mi MyInt) Print() {
	fmt.Println(mi)
}

// 構造体の埋め込み
func main() {
	var mi MyInt
	fmt.Println(mi)
	fmt.Printf("%T\n", mi)

	// i := 100
	// fmt.Println(mi + i) // invalid operation: mi + i (mismatched types MyInt and int)

	mi.Print()
}
