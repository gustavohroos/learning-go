package main

import (
	"fmt"
	"sync"
	"time"
)

func Example2() {
	fmt.Println("------ Example2 ------")

	channel1, channel2 := make(chan int), make(chan int)
	functions := 5

	go send3(100, channel1)
	go send4(functions, channel1, channel2)

	for v := range channel2 {
		fmt.Println(v)
	}

}

func send3(n int, channel chan int) {
	for i := 0; i < n; i++ {
		channel <- i
	}
	close(channel)
}

func send4(functions int, channel1, channel2 chan int) {
	var wg sync.WaitGroup

	for i := 0; i < functions; i++ {
		wg.Add(1)
		go func() {
			for v := range channel1 {
				channel2 <- work2(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(channel2)
}

func work2(n int) int {
	time.Sleep(time.Millisecond * 1000)
	return n
}
