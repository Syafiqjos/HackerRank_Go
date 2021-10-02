// HackerRank
// syafiqjos
// https://www.hackerrank.com/challenges/plus-minus/problem

package main

import (
	"fmt"
)

func main() {
	n := 0
	fmt.Scanf("%v%c", &n)

	arr := make([]int, n)
	for i := range arr {
		fmt.Scanf("%v%c", &arr[i])
	}

	positive := 0
	negative := 0
	zero := 0

	for _, v := range arr {
		if v == 0 {
			zero++
		} else if v < 0 {
			negative++
		} else if v > 0 {
			positive++
		}
	}

	ratio_p := float64(positive) / float64(n)
	ratio_n := float64(negative) / float64(n)
	ratio_z := float64(zero) / float64(n)

	fmt.Println(ratio_p)
	fmt.Println(ratio_n)
	fmt.Println(ratio_z)
}
