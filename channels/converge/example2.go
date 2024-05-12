package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Example2() {
	fmt.Println("------ Example2 ------")
	c := converge(work("First"), work("Secnd"))

	for i := 0; i < 30; i++ {
		fmt.Println(<-c)
	}
}

func work(s string) chan string {
	c := make(chan string)
	go func(s string, c chan string) {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Func %v says: %v", s, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}(s, c)
	return c
}

func converge(x, y chan string) chan string {
	n := make(chan string)
	go func() {
		for {
			n <- <-x
		}
	}()
	go func() {
		for {
			n <- <-y
		}
	}()
	return n
}
