package main

import "fmt"

func Example1() {
	fmt.Println("------ Example1 ------")
	x := 20
	a := make(chan int)
	b := make(chan int)

	go func(x int) {
		for i := 0; i < x; i++ {
			a <- i
		}
	}(x / 2)

	go func(x int) {
		for i := 0; i < x; i++ {
			b <- i
		}
	}(x / 2)

	for i := 0; i < x; i++ {
		select {
		case v := <-a:
			fmt.Println("From a:", v)
		case v := <-b:
			fmt.Println("From b:", v)
		}
	}

}
