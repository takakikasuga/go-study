package main

import (
	"fmt"
)

const Pi = 3.14

const (
	URL      = "http://www.google.com"
	SiteName = "Google"
)

const List1, List2 = 3.14, 3.00

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

// var Big int = 9223372036854775807 + 1

const Big = 9223372036854775807 + 1

const (
	c0 = iota // c0 == 0
	c1 = iota // c1 == 1
	c2 = iota // c2 == 2
)

// @desc 定数
func main() {
	fmt.Println(Pi)

	// Pi = 3
	// fmt.Println(Pi)

	fmt.Println(URL, SiteName)
	fmt.Println(A, B, C, D, E, F)
	fmt.Println(List1, List2)
	fmt.Println(Big - 1)
	fmt.Println(c0, c1, c2)
}
