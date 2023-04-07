package main

import (
	"fmt"
	"sort"
)

func custom_sort(arr []string, N int) []string {
	new_arr := []string{}
	len_str_map := make(map[int][]string)
	// make a mulit-valued map, as in this case keys can be same
	for i := range arr {
		len_str_map[len(arr[i])] = append(len_str_map[len(arr[i])], arr[i])
	}

	// always important that we mention size of arr/slice
	// mostly use slice
	len_arr := make([]int, len(len_str_map))
	idx := 0
	for k, v := range len_str_map {
		len_arr[idx] = k
		idx += 1
		if len(v) > 1 {
			sort.Strings(v)
		}
	}

	sort.Ints(len_arr)

	for i := range len_arr {
		new_arr = append(new_arr, len_str_map[len_arr[i]]...)
	}
	new_arr = append(new_arr[:N], new_arr[N:]...)
	return new_arr
}

func main() {
	arr := []string{"abcd", "abc", "def", "ab"}
	N := len(arr)
	fmt.Println(custom_sort(arr, N))
}
