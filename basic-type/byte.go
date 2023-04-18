package main

import "fmt"

// @desc byte(unit8)型
func main() {
	byteA := []byte{72, 73}

	fmt.Println(byteA)
	fmt.Printf("%T\n", byteA)

	fmt.Println(string(byteA))

	c := []byte("AI")
	fmt.Println(c)
	fmt.Println(string(c))
}
