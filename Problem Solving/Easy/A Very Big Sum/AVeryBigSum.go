// HackerRank
// syafiqjos
// https://www.hackerrank.com/challenges/a-very-big-sum/problem

package main

import (
	"fmt"
)

func main() {
	n := 0
	sum := int64(0)

	fmt.Scanf("%v%c", &n)

	for i := 0; i < n; i++ {
		x := 0
		fmt.Scanf("%v%c", &x)

		sum += int64(x)
	}

	fmt.Println(sum)
}
