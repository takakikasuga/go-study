package main

import (
	"fmt"
)

// @desc スライス（for）
func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println("sl", sl)

	for i, v := range sl {
		fmt.Println(i, v)
	}

	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}
}
