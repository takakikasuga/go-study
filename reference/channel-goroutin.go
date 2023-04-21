package main

import (
	"fmt"
	"time"
)

func receiver(c chan int) {
	fmt.Println("receiver")
	for {
		fmt.Println("Start")
		i := <-c
		fmt.Println(i)
		fmt.Println("End")
	}
}

// @desc チャネル-ゴルーチン(複数のゴルーチン間でのデータの受け渡しをするために設計されたデータ構造)
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// fmt.Println(<-ch1)

	go receiver(ch1)
	go receiver(ch2)

	// i := 0
	// for i < 100 {
	// 	ch1 <- i
	// 	ch2 <- i
	// 	time.Sleep(50 * time.Microsecond)
	// 	i++
	// }
	time.Sleep(8000 * time.Microsecond)
	// defer fmt.Println("defer")
}
