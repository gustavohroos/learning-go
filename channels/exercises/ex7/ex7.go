package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				c <- j
			}
		}()
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
