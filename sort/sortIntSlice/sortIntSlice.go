package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	si := []int{rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100)}
	fmt.Println(si)
	sort.Ints(si)
	fmt.Println(si)
}
