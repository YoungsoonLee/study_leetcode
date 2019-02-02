package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func maxStep(n int32, k int32) int32 {
	if isTriangular(k) {
		return n*(n+1)/2 - 1
	} else {
		return n * (n + 1) / 2
	}
}

func isTriangular(k int32) bool {
	for i := 0; i <= int(k); i++ {
		if i*(i+1)/2 == int(k) {
			return true
		}
	}
	return false
}

/*
func maxStep(n int32, k int32) int32 {
	// Write your code here

	onStep := 0

	var resultS1 int32
	var resultS2 int32

	resultS1, resultS2 = 0, 0

	// sinario 1
	onStep = onStep + 1 // 1 step start
	resultS1 = getStep(int32(n), int32(k), int32(onStep))

	if k > 2 {

		// sinario 2
		onStep = 0
		onStep = onStep + 2 // 2 step start
		resultS2 = getStep(int32(n), int32(k), int32(onStep))

	}
	fmt.Println(resultS1, resultS2)

	if resultS1 > resultS2 {
		return resultS1
	}

	return resultS2
}

func getStep(n, k, onStep int32) int32 {
	action := 1

	var result int32

	for {
		// fmt.Println(n, k, onStep, action)

		result = onStep
		onStep = onStep + 1 // add 1 step

		// do something
		if onStep == k {
			continue // skip
		}

		if action == int(n) {
			break
		}

		action++
	}

	return result
}
*/

func main() {
	n, k := 3, 3
	fmt.Println(maxStep(int32(n), int32(k)))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := maxStep(n, k)

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
