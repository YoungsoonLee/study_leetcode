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
 * Complete the aVeryBigSum function below.
 */
func aVeryBigSum(n int32, ar []int64) int64 {
	/*
	 * Write your code here.
	 */

	var ret int64
	ret = 0

	fmt.Println(n)
	fmt.Println(ar)

	return ret

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	//outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer outputFile.Close()

	//writer := bufio.NewWriterSize(outputFile, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int64

	for arItr := 0; arItr < int(n); arItr++ {
		arItem, err := strconv.ParseInt(arTemp[arItr], 10, 64)
		checkError(err)
		ar = append(ar, arItem)
	}

	result := aVeryBigSum(n, ar)
	fmt.Println(result)

	//fmt.Fprintf(writer, "%d\n", result)

	//writer.Flush()
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
