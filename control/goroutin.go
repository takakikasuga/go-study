package main

import (
	"fmt"
	"time"
)

// @desc 並行処理(goroutin)

func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Microsecond)
	}
}

func main() {
	go sub()
	go sub()

	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Microsecond)
	}
}
