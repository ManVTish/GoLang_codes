package main

import "fmt"

func subcount(A []int, B []int, N int) int {
	count := 0
	for {
		for i := 0; i < N; i++ {
			if A[i] < B[i] {
				return count
			}
		}
		count++
		for i := 0; i < N; i++ {
			A[i] -= B[i]
		}
	}
}

func main() {
	A := []int{7, 12, 100, 6, 200}
	B := []int{1, 1, 1, 1, 1}
	N := len(A)
	fmt.Println(subcount(A, B, N))
}
