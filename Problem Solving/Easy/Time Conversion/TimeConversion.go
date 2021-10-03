// HackerRank
// syafiqjos
// https://www.hackerrank.com/challenges/time-conversion/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversionNew(s string) string {
	t, _ := time.Parse("03:04:05PM", s)

	return t.Format("15:04:05")
}

func timeConversion(s string) string {
	// Write your code here
	str := s[0 : len(s)-2]
	split := strings.Split(str, ":")

	ampm := s[len(s)-2:]

	hh, _ := strconv.ParseInt(split[0], 10, 64)
	mm, _ := strconv.ParseInt(split[1], 10, 64)
	ss, _ := strconv.ParseInt(split[2], 10, 64)

	if strings.ToLower(ampm) == "am" {
		if hh == 12 {
			hh = 0
		} else {

		}
	} else if strings.ToLower(ampm) == "pm" {
		if hh == 12 {
			hh = 12
		} else {
			hh += 12
		}
	}

	var output string = fmt.Sprintf("%02v:%02v:%02v", hh, mm, ss)

	return output
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	// result := timeConversion(s)
	result := timeConversionNew(s)

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
