// HackerRank
// syafiqjos
// https://www.hackerrank.com/challenges/encryption/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

/*
 * Complete the 'encryption' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func encryption(s string) string {
	root := math.Sqrt(float64(len(s)))

	n := int32(math.Floor(root))
	m := int32(math.Ceil(root))

	if n*m < int32(len(s)) {
		n++
	}

	fmt.Println(n)
	fmt.Println(m)

	input := ""

	for i := int32(0); i < n; i++ {
		input += s[i*m : int32(math.Min(float64(i*m+m), float64(len(s))))]

		for int32(len(input)) < i+(i+1)*m {
			input += " "
		}

		input += "\n"
	}

	fmt.Println(input)

	output := ""

	for i := int32(0); i < m; i++ {
		for j := int32(0); j < n; j++ {
			k := i + (m+1)*j
			if input[k] != ' ' || (output[len(output)-1] == ' ') {
				output += string(input[k])
			}
		}
		output += " "
	}

	return output
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := encryption(s)

	fmt.Fprintf(writer, "%s\n", result)

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
