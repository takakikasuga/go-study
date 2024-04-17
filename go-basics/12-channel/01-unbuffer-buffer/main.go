package main

import (
	"fmt"
	"runtime"
)

func main() {
	// ch := make(chan int)
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	ch <- 10
	// 	time.Sleep(500 * time.Millisecond)
	// }()
	// fmt.Println(<-ch)
	// wg.Wait()

	ch := make(chan int)
	go func() {
		fmt.Println(<-ch)
	}()
	ch <- 10
	fmt.Println("num of working goroutines: ", runtime.NumGoroutine())

	ch2 := make(chan int, 2)
	ch2 <- 2
	ch2 <- 3
	fmt.Println(<-ch2)
}
