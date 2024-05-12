package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0
	totalGoroutines := 10

	var wg sync.WaitGroup
	wg.Add(totalGoroutines)

	var mu sync.Mutex

	for i := 0; i < totalGoroutines; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
	fmt.Println("Goroutines:", runtime.NumGoroutine())

}
