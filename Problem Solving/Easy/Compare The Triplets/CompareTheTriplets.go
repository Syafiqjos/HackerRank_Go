// HackerRank
// syafiqjos
// https://www.hackerrank.com/challenges/compare-the-triplets/problem

package main

import (
	"fmt"
)

func main() {
	n := 3
	a := make([]int, n)
	b := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%v%c", &a[i])
	}

	for i := 0; i < n; i++ {
		fmt.Scanf("%v%c", &b[i])
	}

	scoreA := 0
	scoreB := 0

	for i := 0; i < n; i++ {
		if a[i] > b[i] {
			scoreA++
		} else if a[i] < b[i] {
			scoreB++
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%v\n", a[i])
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%v\n", b[i])
	}

	fmt.Printf("%v %v\n", scoreA, scoreB)
}
