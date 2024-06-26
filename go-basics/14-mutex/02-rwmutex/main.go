package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var rwMu sync.RWMutex
	var c int
	wg.Add(4)
	go write(&rwMu, &wg, &c)
	go read(&rwMu, &wg, &c)
	go read(&rwMu, &wg, &c)
	go read(&rwMu, &wg, &c)

	wg.Wait()
	fmt.Println("Done")
}

func read(mu *sync.RWMutex, wg *sync.WaitGroup, c *int) {
	defer wg.Done()
	time.Sleep(10 * time.Millisecond)
	mu.RLock()
	defer mu.RUnlock()
	fmt.Println("Read lock")
	fmt.Println(*c)
	time.Sleep(1 * time.Second)
	fmt.Println("Read unlock")
}

func write(mu *sync.RWMutex, wg *sync.WaitGroup, c *int) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("Write lock")
	*c += 1
	time.Sleep(1 * time.Second)
	fmt.Println("Write unlock")
}
