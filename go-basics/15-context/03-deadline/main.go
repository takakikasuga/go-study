package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(20*time.Millisecond))
	defer cancel()
	ch := subTask(ctx)
	v, ok := <-ch
	if ok {
		fmt.Printf("success: %v\n", v)
	} else {
		fmt.Println("subTask: Deadline exceeded")
	}
	fmt.Println("Done")
}

func subTask(ctx context.Context) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		deadline, ok := ctx.Deadline()
		fmt.Printf("Deadline %v\n %v\n", deadline, ok)

		if ok {
			if deadline.Sub(time.Now().Add(30*time.Millisecond)) < 0 {
				fmt.Println("Deadline exceeded")
				return
			}
		}
		time.Sleep(30 * time.Microsecond)
		ch <- "Hello"
	}()

	return ch
}
