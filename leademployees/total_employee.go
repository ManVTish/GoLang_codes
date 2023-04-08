package main

import (
	"fmt"
	"sort"
)

// ambiguous due to not enough test cases
func total_employees(mat [][]int, N int, K []int) int {
	count := 0

	for i := range mat {
		a := mat[i][0]
		b := mat[i][1]
		// for 1st element check if it's present in K
		// for 2nd element check if it's not present in K
		if sort.SearchInts(K, a) != len(K) && sort.SearchInts(K, b) == len(K) {
			count++
		}
	}
	return count + len(K) // add the number of K elements with count (number of employees under head)
}

func main() {
	N := 7
	K := []int{2, 4}

	a := []int{1, 2}
	b := []int{1, 7}
	c := []int{2, 4}
	d := []int{4, 5}
	e := []int{4, 6}
	mat := [][]int{a, b, c, d, e}
	fmt.Println(total_employees(mat, N, K))
}
