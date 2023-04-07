package main

import (
	"fmt"
)

// bot b1 and b2 starts at t=0 which will complete the most tasks

func bot_challenge(N int, arr []int) string {
	b1 := 0
	b2 := 0
	idx_b1 := 0
	idx_b2 := len(arr) - 1

	if len(arr) == 1 {
		return "b1"
	}

	for (idx_b2 - idx_b1) > 1 {
		if arr[idx_b1] < arr[idx_b2] {
			idx_b1++
			b1++
			arr[idx_b2] -= arr[idx_b1]
		} else if arr[idx_b1] == arr[idx_b2] {
			idx_b1++
			b1++
			idx_b2--
			b2++
		} else {
			idx_b2--
			b2++
			arr[idx_b1] -= arr[idx_b2]
		}
	}

	if b1 > b2 {
		return "b1"
	} else if b1 == b2 {
		return "-1"
	} else {
		return "b2"
	}
}

func main() {
	N := 7
	// each element represents the time taken to complete a task
	arr := []int{3, 7, 2, 1, 5, 8, 4}
	fmt.Println(bot_challenge(N, arr))
}
