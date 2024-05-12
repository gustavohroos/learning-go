package main

import (
	"fmt"
	"sync"
)

var wg_example3 = sync.WaitGroup{}

func Example3() {
	fmt.Println("------ Example3 ------")
	odd, even, quit := make(chan int), make(chan int), make(chan bool)
	wg_example3.Add(2)
	go send(odd, even, quit)
	go func() {
		defer wg_example3.Done()
		for {
			select {
			case v, ok := <-odd:
				if ok {
					println("Odd: ", v)
				}
			case v, ok := <-even:
				if ok {
					println("Even:", v)
				}
			case <-quit:
				println("Quit")
				return
			}
		}
	}()
	wg_example3.Wait()
}

func send(odd, even chan int, quit chan bool) {
	defer wg_example3.Done()
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(odd)
	close(even)
	quit <- true
	close(quit)
}
