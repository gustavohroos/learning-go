package main

import "fmt"

func main() {
	for x := 0; x < 10; x++ {
		fmt.Printf("Factorial of %v: %v\n", x, factorial(x))
	}
}

func factorial(x int) int {
	if x <= 1 {
		return x
	}
	return x * factorial(x-1)
}
