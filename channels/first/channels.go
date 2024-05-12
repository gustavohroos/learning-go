package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		fmt.Println("This is a goroutine!")
		channel <- 42
	}()

	fmt.Println(<-channel)
}
