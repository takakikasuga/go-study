package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg.Add(2)
	go func() {
		defer wg.Done()
		v, err := criticalTask(ctx)
		if err != nil {
			fmt.Printf("criticalTask: %v\n", err)
			cancel()
			return
		}
		fmt.Printf("success: %v\n", v)
	}()

	go func() {
		defer wg.Done()
		v, err := normalTask(ctx)
		if err != nil {
			fmt.Printf("normalTask: %v\n", err)
			return
		}
		fmt.Printf("success: %v\n", v)
	}()
	wg.Wait()

}

func criticalTask(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 1200*time.Millisecond)
	defer cancel()

	t := time.NewTicker(1000 * time.Millisecond)

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-t.C:
		t.Stop()
	}
	return "A", nil
}

func normalTask(ctx context.Context) (string, error) {
	t := time.NewTicker(3000 * time.Millisecond)

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-t.C:
		t.Stop()
	}
	return "B", nil
}
