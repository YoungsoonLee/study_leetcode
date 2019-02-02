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
 * Complete the 'numberOfPaths' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY a as parameter.
 */

func numberOfPaths(a [][]int32) int32 {
	// Write your code here
	m := len(a)
	if m == 0 {
		return 0
	}

	n := len(a[0])
	if n == 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	if a[0][0] == 1 {
		dp[0][0] = 1
	}

	for i := 1; i < m; i++ {
		if a[i][0] == 1 {
			dp[i][0] = dp[i-1][0]
		}
	}

	for j := 1; j < n; j++ {
		if a[0][j] == 1 {
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if a[i][j] == 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	fmt.Println(dp)
	fmt.Println(dp[m-1][n-1])
	return int32(dp[m-1][n-1])
}

func main() {
	n := [][]int32{{1, 1, 0, 1}, {1, 1, 1, 1}}
	numberOfPaths(n)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	aRows, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	aColumns, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var a [][]int32
	for i := 0; i < int(aRows); i++ {
		aRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var aRow []int32
		for _, aRowItem := range aRowTemp {
			aItemTemp, err := strconv.ParseInt(aRowItem, 10, 64)
			checkError(err)
			aItem := int32(aItemTemp)
			aRow = append(aRow, aItem)
		}

		if len(aRow) != int(aColumns) {
			panic("Bad input")
		}

		a = append(a, aRow)
	}

	result := numberOfPaths(a)

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
