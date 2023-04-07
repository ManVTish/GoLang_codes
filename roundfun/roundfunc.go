package main

import (
	"fmt"
	"math"
)

func round_func(arrA []float64, arrB []float64, n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		decimal := arrA[i] / arrB[i]
		arr[i] = int(math.Round(decimal))
	}
	return arr
}

func main() {
	arrA := []float64{1, 3, 5, 12, 3}
	arrB := []float64{1, 4, 4, 5, 3}
	N := len(arrA)
	fmt.Println(round_func(arrA, arrB, N))
}
