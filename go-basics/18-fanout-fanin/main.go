package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	cores := runtime.NumCPU()
	fmt.Println("cores:", cores)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nums := []int{1, 2, 3, 4, 5, 7, 8}

	outChs := make([]<-chan string, cores)
	genCh := generate(ctx, nums...)
	fmt.Println("genCh:", genCh)
	for i := 0; i < cores; i++ {
		fmt.Printf("i:%v\n", i)
		outChs[i] = fanOut(ctx, genCh, i+1)
	}
	fmt.Println("outChs:", outChs)

	var i int
	flag := true
	for v := range fanIn(ctx, outChs...) {
		if i == 3 {
			cancel()
			flag = false
		}

		if flag {
			fmt.Printf("✌️%v\n", v)
		}
		i++
	}

	fmt.Println("Done")
}

func generate(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- n:
				fmt.Println("generate", n)
			}
		}
	}()
	return out
}
func fanOut(ctx context.Context, in <-chan int, id int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		heavyWork := func(i int, id int) string {
			time.Sleep(5 * time.Second)
			return fmt.Sprintf("result:%v（id:%v）", i*i, id)
		}
		for v := range in {
			select {
			case <-ctx.Done():
				return
			case out <- heavyWork(v, id):
				fmt.Println("fanOut")
			}
		}
	}()
	return out
}

func fanIn(ctx context.Context, chs ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)
	multiplex := func(ch <-chan string) {
		defer wg.Done()
		for text := range ch {
			select {
			case <-ctx.Done():
				fmt.Println("ctx.Done")
				return
			case out <- text:
				fmt.Println("fanIn")
			}
		}
	}

	wg.Add(len(chs))

	for _, ch := range chs {
		go multiplex(ch)
	}

	go func() {
		fmt.Println("🔥 Before")
		wg.Wait()
		fmt.Println("🔥 After")
		close(out)
	}()

	fmt.Println("fanIn / End")
	return out
}
