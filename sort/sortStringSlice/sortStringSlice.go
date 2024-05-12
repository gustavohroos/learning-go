package main

import (
	"fmt"
	"sort"
)

func main() {
	ss := []string{"banana", "apple", "orange", "grape", "watermelon", "pineapple", "cherry", "strawberry"}
	fmt.Println(ss)
	sort.Strings(ss)
	fmt.Println(ss)
}
