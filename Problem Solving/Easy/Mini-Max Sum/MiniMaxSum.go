// HackerRank
// syafiqjos
// https://www.hackerrank.com/challenges/mini-max-sum/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func miniMaxSum(arr []int32) (int64, int64) {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	n := len(arr)

	mini := int64(0)
	maxi := int64(0)

	for i := 0; i < n-1; i++ {
		mini += int64(arr[i])
	}

	for i := n - 1; i > 0; i-- {
		maxi += int64(arr[i])
	}

	return mini, maxi
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	mini, maxi := miniMaxSum(arr)
	fmt.Printf("%v %v\n", mini, maxi)
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
