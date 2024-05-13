package main

import (
	"math"
)

func Sum(a ...int) int {
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}

type point struct {
	x, y int
}

func EuclideanDistance(a, b point) float64 {
	// d=√((x2 – x1)² + (y2 – y1)²)
	return math.Sqrt(math.Pow(float64(b.x-a.x), 2) + math.Pow(float64(b.y-a.y), 2))
}

func ManhattamDistance(a, b point) int {
	// d=|x2 – x1| + |y2 – y1|
	return int(math.Abs(float64(b.x-a.x)) + math.Abs(float64(b.y-a.y)))
}
