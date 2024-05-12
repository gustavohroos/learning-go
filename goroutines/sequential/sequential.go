package main

import (
	"fmt"
	"time"
)

var size = 1000000

func main() {
	start := time.Now()

	func1()
	func2()

	sequentially := time.Since(start)
	fmt.Println("Time running", size, "println sequentially: ", sequentially)
}

func func1() {
	for i := 0; i < size; i++ {
		println("func1:", i)
	}
}

func func2() {
	for i := 0; i < size; i++ {
		println("func2:", i)
	}
}
