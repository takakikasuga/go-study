package main

import "fmt"

// @desc チャネル(複数のゴルーチン間でのデータの受け渡しをするために設計されたデータ構造)
func main() {
	var ch1 chan int // 双方向のチャネル

	// var ch2 <-chan int // 僕たちがチャネルから受信する専用のチャネル / チャネルは送信のみ行う。
	// var receiveCh <-chan int
	// <-receiveCh
	// <-receiveCh
	// <-receiveCh

	// var ch3 chan<- int // 僕たちがチャネルに送信する専用のチャネル / チャネルは受信のみ受け付ける。
	var sendCh chan<- int
	fmt.Println("cap(sendCh) === ", cap(sendCh))
	// sendCh <- 1
	// sendCh <- 2
	// sendCh <- 3

	ch1 = make(chan int)
	ch2 := make(chan int)

	fmt.Println("cap(ch1) === ", cap(ch1))
	fmt.Println("cap(ch2) === ", cap(ch2))
	ch3 := make(chan int, 5)
	fmt.Println("cap(ch3) === ", cap(ch3))

	ch3 <- 1
	fmt.Println("ch3 === ", ch3)
	fmt.Println("len(ch3) === ", len(ch3))
	ch3 <- 2
	ch3 <- 3
	fmt.Println("ch3 === ", ch3)
	fmt.Println("len(ch3) === ", len(ch3))

	i := <-ch3
	fmt.Println("i === ", i)
	i2 := <-ch3
	fmt.Println("i2 === ", i2)
	fmt.Println("len(ch3) === ", len(ch3))
	fmt.Println("<-ch3 === ", <-ch3)
	fmt.Println("len(ch3) === ", len(ch3))

	ch3 <- 1
	fmt.Println("len(ch3) === ", len(ch3))
	fmt.Println("<-ch3 === ", <-ch3)
	fmt.Println("len(ch3) === ", len(ch3))
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	ch3 <- 6 // fatal error: all goroutines are asleep - deadlock!
}
