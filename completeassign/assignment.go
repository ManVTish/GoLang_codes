package main

import (
	"fmt"
)

type Assignment struct {
	str string
	day int
}

func complete_assignment(N int, arr []int) Assignment {
	assign := Assignment{"", 0}

	sum := 0
	f := 0
	count := 0
	for sum < N && count <= 60 {
		d := 0
		for i := range arr {
			sum += arr[i]
			d++
			if sum >= N {
				f++
				break
			}
		}
		count += d
	}

	if f == 0 {
		assign.str = "NO"
	} else {
		assign.str = "YES"
	}
	assign.day = count

	return assign
}

func main() {
	N := 100
	arr := []int{15, 50, 20, 30, 10, 15, 30}

	res := complete_assignment(N, arr)
	if res.str == "YES" {
		fmt.Println(res.str)
		fmt.Println(res.day)
	} else {
		fmt.Println(res.str)
	}
}
