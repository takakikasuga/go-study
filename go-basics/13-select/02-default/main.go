package main

import (
	"sync"
	"time"
)

const bufSize = 3

func main() {
	var wg sync.WaitGroup
	ch := make(chan string, bufSize)
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < bufSize; i++ {
			time.Sleep(1000 * time.Millisecond)
			ch <- "Hello"
		}
	}()

	for i := 0; i < 3; i++ {
		select {
		case msg := <-ch:
			println("received", msg)
		default:
			println("no message received")
		}
		time.Sleep(1500 * time.Millisecond)
	}
	println(<-ch)
}
