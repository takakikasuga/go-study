package main

import (
	"context"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	var wg sync.WaitGroup
	wg.Add(2)
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	go func() {
		defer wg.Done()
		defer println("火の穂")
		// defer time.Sleep(1000 * time.Millisecond)
		time.Sleep(500 * time.Millisecond)
		ch1 <- "A"
	}()

	go func() {
		defer wg.Done()
		time.Sleep(800 * time.Millisecond)
		ch2 <- "B"
	}()

loop:
	for ch1 != nil || ch2 != nil {
		select {
		case <-ctx.Done():
			println("timeout")
			break loop
		case msg1 := <-ch1:
			println("received", msg1)
			ch1 = nil
		case msg2 := <-ch2:
			println("received", msg2)
			ch2 = nil
		}
	}
	wg.Wait()
	println("done")
}
