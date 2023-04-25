package main

import (
	"fmt"
)

func init() {
	fmt.Println("init")
}

func init() {
	fmt.Println("init2")
}

// @desc init
func main() {
	fmt.Println("Main")
}
