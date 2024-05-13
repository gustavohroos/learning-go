package main

import (
	"fmt"
	"testing"
)

type test struct {
	data     []int
	expected int
}

func TestSumTable(t *testing.T) {
	tests := []test{
		{[]int{1, 2, 3, 4}, 10},
		{[]int{10, 20, 30, 40}, 100},
		{[]int{1, 2, 3, 4, 5}, 15},
	}
	for _, v := range tests {
		test := Sum(v.data...)
		if test != v.expected {
			t.Error("Expected", v.expected, "got", test)
		}
	}
}

func TestSum(t *testing.T) {
	test := Sum(1, 2, 3, 4)
	expected := 10
	if test != expected {
		t.Error("Expected", expected, "got", test)
	}
}

func TestEuclidean(t *testing.T) {
	a := point{1, 2}
	b := point{1, 6}
	test := EuclideanDistance(a, b)
	expected := 4.
	if test != expected {
		t.Error("Expected", expected, "got", test)
	}
}

func TestManhattam(t *testing.T) {
	a := point{1, 2}
	b := point{5, 12}
	test := ManhattamDistance(a, b)
	expected := 14
	if test != expected {
		t.Error("Expected", expected, "got", test)
	}
	test = ManhattamDistance(b, a)
	expected = 14
	if test != expected {
		t.Error("Expected", expected, "got", test)
	}
}

func ExampleSum() {
	fmt.Println(Sum(1, 2, 3, 4))
	fmt.Println(Sum(10, 20, 30, 40))
	fmt.Println(Sum(1, 2, 3, 4, 5))
	// Output: 10
	// 100
	// 15
}

func ExampleEuclideanDistance() {
	a := point{1, 2}
	b := point{1, 6}
	fmt.Println(EuclideanDistance(a, b))
	// Output: 4
}

func ExampleManhattamDistance() {
	a := point{1, 2}
	b := point{5, 12}
	fmt.Println(ManhattamDistance(a, b))
	// Output: 14
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2, 3, 4)
	}
}

func BenchmarkEuclideanDistance(b *testing.B) {
	d := point{1, 2}
	e := point{1, 6}
	for i := 0; i < b.N; i++ {
		EuclideanDistance(d, e)
	}
}

func BenchmarkManhattamDistance(b *testing.B) {
	d := point{1, 2}
	e := point{5, 12}
	for i := 0; i < b.N; i++ {
		ManhattamDistance(d, e)
	}
}
