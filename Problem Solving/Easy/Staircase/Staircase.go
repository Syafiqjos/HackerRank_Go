// HackerRank
// syafiqjos
// https://www.hackerrank.com/challenges/staircase/problem

package main

import (
	"fmt"
)

func main() {
	n := 0
	fmt.Scanf("%v%c", &n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i >= n-j-1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
