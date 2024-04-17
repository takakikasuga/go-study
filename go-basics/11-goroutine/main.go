package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func main() {
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("goroutine invoked")
	// }()
	// wg.Wait()
	// fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())
	// fmt.Printf("main func finished\n")

	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln("Error: ", err)
		}
	}()
	if err := trace.Start(f); err != nil {
		log.Fatalln("Error: ", err)
	}
	defer trace.Stop()
	ctx, t := trace.NewTask(context.Background(), "main")
	defer t.End()
	fmt.Println("The number of CPUs: ", runtime.NumCPU())

	// task(ctx, "task1")
	// task(ctx, "task2")
	// task(ctx, "task3")

	var wg sync.WaitGroup
	wg.Add(3)
	go cTask(ctx, &wg, "task1")
	go cTask(ctx, &wg, "task2")
	go cTask(ctx, &wg, "task3")
	wg.Wait()
	fmt.Println("main func finished")
}

func task(ctx context.Context, name string) {
	defer trace.StartRegion(ctx, name).End()
	time.Sleep(time.Second)
	fmt.Println(name)
}

func cTask(ctx context.Context, wg *sync.WaitGroup, name string) {
	defer trace.StartRegion(ctx, name).End()
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println(name)
}
