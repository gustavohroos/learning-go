package main

import "fmt"

func main() {
	c := make(chan int)
	cs := make(chan<- int)
	cr := make(<-chan int)

	fmt.Println("-------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cs\t%T\n", cs)
	fmt.Printf("cr\t%T\n", cr)

	fmt.Println("-------")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))

	fmt.Println("-------")
	go send(c)
	receive(c)

}

func send(s chan<- int) {
	s <- 42
}

func receive(r <-chan int) {
	fmt.Println(<-r)
}
