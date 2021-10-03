// HackerRank
// syafiqjos
// https://www.hackerrank.com/challenges/magic-square-forming/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'formingMagicSquare' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY s as parameter.
 */

func getMagicSquareRotation(square [][]int32) [][]int32 {
	newM := make([][]int32, 3)
	for i := 0; i < 3; i++ {
		newM[i] = make([]int32, 3)
	}

	newM[0][0], newM[0][1], newM[0][2], newM[1][0], newM[1][1], newM[1][2], newM[2][0], newM[2][1], newM[2][2] = square[2][0], square[1][0], square[0][0], square[2][1], square[1][1], square[0][1], square[2][2], square[1][2], square[0][2]

	return newM
}

func getMatrixDifference(square, other [][]int32) int32 {
	val := int32(0)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			val += int32(math.Abs(float64(square[i][j] - other[i][j])))
		}
	}

	return val
}

func formingMagicSquare(s [][]int32) int32 {
	// Write your code here

	magicSquareTemplate := [][]int32{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}}
	magicSquareTemplateInverse := [][]int32{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}}

	diffs := [8]int32{}

	diffs[0] = getMatrixDifference(s, magicSquareTemplate)
	diffs[1] = getMatrixDifference(s, getMagicSquareRotation(magicSquareTemplate))
	diffs[2] = getMatrixDifference(s, getMagicSquareRotation(getMagicSquareRotation(magicSquareTemplate)))
	diffs[3] = getMatrixDifference(s, getMagicSquareRotation(getMagicSquareRotation(getMagicSquareRotation(magicSquareTemplate))))

	diffs[4] = getMatrixDifference(s, magicSquareTemplateInverse)
	diffs[5] = getMatrixDifference(s, getMagicSquareRotation(magicSquareTemplateInverse))
	diffs[6] = getMatrixDifference(s, getMagicSquareRotation(getMagicSquareRotation(magicSquareTemplateInverse)))
	diffs[7] = getMatrixDifference(s, getMagicSquareRotation(getMagicSquareRotation(getMagicSquareRotation(magicSquareTemplateInverse))))

	min := int32(0)

	for k, v := range diffs {
		if k == 0 || v < min {
			min = v
		}
	}

	return min
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	var s [][]int32
	for i := 0; i < 3; i++ {
		sRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var sRow []int32
		for _, sRowItem := range sRowTemp {
			sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
			checkError(err)
			sItem := int32(sItemTemp)
			sRow = append(sRow, sItem)
		}

		if len(sRow) != 3 {
			panic("Bad input")
		}

		s = append(s, sRow)
	}

	result := formingMagicSquare(s)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
