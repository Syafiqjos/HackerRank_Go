// HackerRank
// syafiqjos
// https://www.hackerrank.com/challenges/grading/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'gradingStudents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY grades as parameter.
 */

func getRoundedGrade(grade, minGrade, multGrade int32) int32 {
	if grade < minGrade {
		return grade
	}

	// Test
	// 85 % 5 == 0
	// 84 % 5 == 4
	// 83 % 5 == 3
	// 82 % 5 == 2
	// 81 % 5 == 1

	var val int32 = grade

	if grade%5 >= multGrade {
		reminder := 5 - grade%5
		val += reminder
	}

	return val
}

func gradingStudents(grades []int32) []int32 {
	// Write your code here

	output := make([]int32, len(grades))

	for k, v := range grades {
		output[k] = getRoundedGrade(v, 38, 3)
	}

	return output
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	gradesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var grades []int32

	for i := 0; i < int(gradesCount); i++ {
		gradesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		gradesItem := int32(gradesItemTemp)
		grades = append(grades, gradesItem)
	}

	result := gradingStudents(grades)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
