package main

import (
	"fmt"
	"sort"
	"strings"
)

func order_str(arr []string, str string, N int) string {
	rank_str := make(map[int]string)

	for _, value := range arr {
		letter := rune(value[0])
		index := strings.IndexRune(str, letter)
		rank_str[index] = value
	}

	key_arr := make([]int, N)
	i := 0
	for key := range rank_str {
		key_arr[i] = key
		i += 1
	}

	sort.Ints(key_arr)

	arr_str := make([]string, N)

	j := 0
	for _, v := range key_arr {
		// arr_byte := []byte(strings.ToUpper(rank_str[v]))
		arr_byte := strings.ToUpper(rank_str[v])
		sub_str := arr_byte[0:1]
		sub_str = strings.ToLower(sub_str)
		sub_str += arr_byte[1:]

		arr_str[j] = sub_str
		j++
	}

	new_str := strings.Join(arr_str, " ")
	return new_str
}

func main() {
	var str string = "nfwdovjeasicmzugbxtlhrkpqy"

	N := 5

	case_str := []string{"hello", "world", "this", "is", "go"}

	fmt.Println(order_str(case_str, str, N))
}
