package main

import "fmt"

func main() {
	a := i()
	for i := 0; i < 10; i++ {
		fmt.Println(a())
	}
}

func i() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
