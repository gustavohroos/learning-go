package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var size = 1000000

func main() {

	start := time.Now()
	go func1()
	go func2()
	wg.Add(2)
	wg.Wait()
	concurrently := time.Since(start)

	fmt.Println("Time running", size, "println concurrently: ", concurrently)
}

func func1() {
	for i := 0; i < size; i++ {
		fmt.Println("func1:", i)
	}
	wg.Done()
}

func func2() {
	for i := 0; i < size; i++ {
		fmt.Println("func2:", i)
	}
	wg.Done()
}
