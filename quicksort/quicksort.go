package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func QuickSort(arr []int, left, right int) []int {
	if left < right {
		pi := partition(arr, left, right)
		QuickSort(arr, left, pi-1)
		QuickSort(arr, pi+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]

	i := left - 1

	for j := left; j < right; j++ {
		if arr[j] < pivot {
			i++
			temp := arr[j]
			arr[j] = arr[i]
			arr[i] = temp
		}
	}

	temp := arr[i+1]
	arr[i+1] = arr[right]
	arr[right] = temp

	return i + 1
}

func arrAreEqual(arr1, arr2 []int) (bool, error) {
	if len(arr1) != len(arr2) {
		return false, errors.New("arrays are not of the same length")
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false, nil
		}
	}
	return true, nil
}

func main() {
	var start, end time.Time
	arrSize := 100000
	arr1 := make([]int, arrSize)
	arr2 := make([]int, arrSize)

	start = time.Now()
	for i := 0; i < arrSize; i++ {
		value := rand.Intn(arrSize)
		arr1[i] = value
		arr2[i] = value
	}
	end = time.Now()
	fmt.Println("array generation: ", end.Sub(start).Seconds())

	start = time.Now()
	QuickSort(arr1[:], 0, len(arr1)-1)
	end = time.Now()
	fmt.Println("implemented quicksort: ", end.Sub(start).Seconds())

	start = time.Now()
	sort.Ints(arr2[:])
	end = time.Now()
	fmt.Println("go quicksort", end.Sub(start).Seconds())

	equal, err := arrAreEqual(arr1, arr2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("arrays are equal: ", equal)
	}

}
