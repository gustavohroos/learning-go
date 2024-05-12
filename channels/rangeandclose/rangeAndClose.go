package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	c := make(chan int)
	t := 10
	wg.Add(2)
	go loop(t, c)
	go prints(c)
	wg.Wait()
}

func loop(t int, cs chan<- int) {
	for i := 0; i < t; i++ {
		cs <- i
	}
	close(cs)
	wg.Done()
}

func prints(cr <-chan int) {
	for v := range cr {
		fmt.Println("Received from channel:", v)
	}
	wg.Done()
}
