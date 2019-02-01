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
 * Complete the function below.
 */
import "sort"

func triangleOrNot(a []int32, b []int32, c []int32) []string {
	// TODO: not three input, more
	fmt.Println(a, b, c)
	r := make([]string, 0)
	n := len(a)
	fmt.Println("n; ", n)

	for i := 0; i < n; i++ {
		temp := make([]int32, 0)
		temp = append(temp, a[i], b[i], c[i])

		//sort.Ints(temp)
		sort.Slice(temp, func(i, j int) bool { return temp[i] < temp[j] })
		fmt.Println(temp)

		var sum int32
		sum = 0
		m := len(temp)

		for j := 0; j < m-1; j++ {
			fmt.Println("inner sum: ", sum, " , temp[j]: ", temp[j], "j: ", j)
			sum = sum + temp[j]
		}

		fmt.Println("outer sum: ", sum, " , temp[m-1]: ", temp[m-1])
		if sum > temp[m-1] {
			r = append(r, "Yes")
		} else {
			r = append(r, "No")
		}
	}
	fmt.Println(r)
	return r
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	aSize, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	var a []int32
	for i := 0; i < int(aSize); i++ {
		aItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)

		a = append(a, aItem)
	}

	bSize, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	var b []int32
	for i := 0; i < int(bSize); i++ {
		bItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		bItem := int32(bItemTemp)

		b = append(b, bItem)
	}

	cSize, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	var c []int32
	for i := 0; i < int(cSize); i++ {
		cItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		cItem := int32(cItemTemp)

		c = append(c, cItem)
	}

	res := triangleOrNot(a, b, c)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, resItem := range res {
		fmt.Fprintf(writer, "%s\n", resItem)
	}

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
