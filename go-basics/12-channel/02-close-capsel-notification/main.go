package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-ch)
	}()
	ch <- 10
	close(ch)
	v, ok := <-ch
	fmt.Printf("v: %v, ok: %v\n", v, ok)
	wg.Wait()

	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	close(ch2)
	v2, ok2 := <-ch2
	fmt.Printf("v: %v, ok: %v\n", v2, ok2)
	v3, ok3 := <-ch2
	fmt.Printf("v: %v, ok: %v\n", v3, ok3)
	v4, ok4 := <-ch2
	fmt.Printf("v: %v, ok: %v\n", v4, ok4)

	ch3 := generateCountStream()
	for v := range ch3 {
		fmt.Println(v)
	}

	nCh := make(chan struct{})
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("goroutine %v\n", i)
			<-nCh
			fmt.Println(i)
		}(i)
	}
	time.Sleep(2 * time.Second)
	close(nCh)
	wg.Wait()
	fmt.Println("done")

	ch4 := generateCountStream2()
	for v := range ch4 {
		fmt.Println(v)
		time.Sleep(2 * time.Second)
	}
}

func generateCountStream() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
	return ch
}

func generateCountStream2() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Println("write to ch")
		}
	}()
	return ch
}
