package main

import (
	"fmt"
	"math/rand"
)

func main() {
	number := rand.Intn(1000)
	number = 712
	var guess int
	for {
		fmt.Print("Your guess: ")
		fmt.Scanf("%d", &guess)
		if guess > number {
			fmt.Println("Too high")
		} else if guess < number {
			fmt.Println("Too low")
		} else {
			fmt.Println("You got it!")
			break
		}
	}
}
