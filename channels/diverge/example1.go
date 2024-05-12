package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Example1() {
	fmt.Println("------ Example1 ------")
	channel1, channel2 := make(chan int), make(chan int)

	go send1(10, channel1)
	go send2(channel1, channel2)

	for v := range channel2 {
		fmt.Println(v)
	}

}

func send1(n int, channel chan int) {
	for i := 0; i < n; i++ {
		channel <- i
	}
	close(channel)
}

func send2(channel1, channel2 chan int) {
	var wg sync.WaitGroup
	for v := range channel1 {
		wg.Add(1)
		go func(x int) {
			channel2 <- Work(x)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(channel2)
}

func Work(x int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	return x
}
