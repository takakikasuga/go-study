package main

import (
	"fmt"
)

// @desc チャネル-セレクト
func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)
	ch1 <- 2
	ch2 <- "A"
	/*
		v1 := <-ch1
		v2 := <-ch2
		fmt.Println("v1 === ", v1)
		fmt.Println("v2 === ", v2)
	*/

	// 複数の条件に当てはまる場合は、ランダムに実行される。
	select {
	case v1 := <-ch1:
		fmt.Println(v1 * 1000)
	case v2 := <-ch2:
		fmt.Println(v2 + "!?")
	default: // 条件に当てはまらない場合、実行される。
		fmt.Println("どちらでもない。")
	}

	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	// レシーバー
	go func() {
		for {
			i := <-ch3
			ch4 <- i * 2
		}
	}()

	go func() {
		for {
			i2 := <-ch4
			ch5 <- i2 - 1
		}
	}()

	n := 0
Label:
	for {
		select {
		case ch3 <- n:
			n++
		case i3 := <-ch5:
			fmt.Println("received", i3)
		default:
			if n > 100 {
				break Label
			}
		}
		/*
			if n > 100 {
				break
			}
		*/
	}

}
