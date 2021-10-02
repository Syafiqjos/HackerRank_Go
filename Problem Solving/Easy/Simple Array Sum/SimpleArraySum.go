// HackerRank
// syafiqjos
// https://www.hackerrank.com/challenges/simple-array-sum/problem

package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scanf("%v%c", &n)

	var arr []int = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v%c", &arr[i])
	}

	sum := 0
	for _, v := range arr {
		sum += v
	}

	fmt.Println(sum)
}
