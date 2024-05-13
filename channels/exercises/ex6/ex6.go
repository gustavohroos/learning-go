package main

func main() {
	c := make(chan int)

	go send(c)

	for v := range c {
		println(v)
	}

}

func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
