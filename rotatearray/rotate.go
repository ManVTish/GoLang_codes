package main

import (
	"fmt"
)

func rotate_elements(N int, arr *[]int) []int {

	min_element := func(a []int) int {
		min := a[0]

		for i := 1; i < len(a); i++ {
			if a[i] < min {
				min = a[i]
			}
		}
		return min
	}

	arr_value := *arr

	for arr_value[0] != min_element(arr_value) {
		arr1 := append(arr_value[1:], arr_value[0])
		arr_value = arr1
	}

	return arr_value
}

func main() {
	N := 5
	arr := []int{3, 2, 1, 5, 4}
	fmt.Println(rotate_elements(N, &arr))
}
