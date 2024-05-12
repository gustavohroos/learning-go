package main

import (
	"fmt"
	"sync"
)

func Example1() {
	fmt.Println("------ Example1 ------")
	odd, even, converge := make(chan int), make(chan int), make(chan int)

	go send(odd, even)
	go receive(odd, even, converge)

	for v := range converge {
		fmt.Println("From converge:", v)
	}

}

func send(odd, even chan int) {
	x := 100

	for i := 0; i < x; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(odd)
	close(even)
}

func receive(odd, even, converge chan int) {
	defer close(converge)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range odd {
			converge <- v
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range even {
			converge <- v
		}
	}()
	wg.Wait()

}
