package main

import "fmt"

func main() {

	arr := make([]int, 0, 3)

	for i := 0; i < 256; i++ {
		arr = append(arr, i)
		fmt.Println(len(arr), cap(arr))
	}
}
