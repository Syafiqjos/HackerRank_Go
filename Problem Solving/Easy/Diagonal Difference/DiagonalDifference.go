// HackerRank
// syafiqjos
// https://www.hackerrank.com/challenges/diagonal-difference/problem

package main

import (
	"fmt"
	"math"
)

func main() {
	n := 0
	fmt.Scanf("%v%c", &n)

	matrix := make([][]int32, n)

	for i := range matrix {
		matrix[i] = make([]int32, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scanf("%v%c", &matrix[i][j])
		}
	}

	ltr := int32(0)
	rtl := int32(0)

	for i := 0; i < n; i++ {
		ltr += matrix[i][i]
	}

	for i := 0; i < n; i++ {
		rtl += matrix[i][n-i-1]
	}

	result := int(math.Abs(float64(ltr) - float64(rtl)))

	fmt.Println(result)
}
