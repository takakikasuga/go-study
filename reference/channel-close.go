package main

import (
	"fmt"
	"time"
)

func receiver(name string, c chan int) {
	for {
		i, ok := <-c
		if !ok { // チャネルの中身が空でありcloseされた場合
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

// @desc チャネル-クローズ
func main() {
	ch1 := make(chan int, 2)

	/*
	   ch1 <- 1
	   close(ch1)
	   // fmt.Println(<-ch1)
	   // ch1 <- 1 // panic: send on closed channel

	   i, ok := <-ch1 // ok = channelがcloseしているかどうか、ただ、データを保持している場合はok = trueになる。
	   fmt.Println(i, ok)
	   i2, ok := <-ch1
	   fmt.Println(i2, ok)
	*/

	go receiver("1.goroutin", ch1)
	go receiver("2.goroutin", ch1)
	go receiver("3.goroutin", ch1)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}

	close(ch1)
	time.Sleep(100 * time.Second)
}
