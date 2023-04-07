package main

import (
	"fmt"
)

func identicalcount(A [][]int, B [][]int, C [][]int, N int, M int) int {
	arr_A := make([]int, N*M)
	arr_B := make([]int, N*M)
	arr_C := make([]int, N*M)

	for i := 0; i < N; i++ {
		arr_A = append(arr_A, A[i]...)
		arr_B = append(arr_B, B[i]...)
		arr_C = append(arr_C, C[i]...)
	}
	arr_A = arr_A[N*M:]
	arr_B = arr_B[N*M:]
	arr_C = arr_C[N*M:]

	count := 0
	for _, va := range arr_A {
		for _, vb := range arr_B {
			for _, vc := range arr_C {
				if va == vb && vb == vc {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	A := [][]int{{2, 2}, {3, 4}}
	B := [][]int{{2, 5}, {6, 7}}
	C := [][]int{{8, 2}, {0, 9}}
	N := 2
	M := 2
	fmt.Println(identicalcount(A, B, C, N, M))
}
