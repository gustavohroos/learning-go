package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func Example2() {

	fmt.Println("------ Example2 ------")

	channel := make(chan int)
	quit := make(chan int)
	wg.Add(2)
	go receiveQuit(channel, quit)
	go sendToChannel(channel, quit)
	wg.Wait()
}

func receiveQuit(channel chan int, quit chan int) {
	for i := 0; i < 20; i++ {
		fmt.Println("Received:", <-channel)
	}
	quit <- 0
	wg.Done()
}

func sendToChannel(channel chan int, quit chan int) {
	anything := 1
	for {
		select {
		case channel <- anything:
			anything++

		case <-quit:
			wg.Done()
			return
		}
	}

}
