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
 * Complete the 'usernameDisparity' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts STRING_ARRAY inputs as parameter.
 */

func usernameDisparity(inputs []string) []int32 {
	//fmt.Println(inputs[0])
	// Write your code here
	result := make([]int32, 0)
	// set init value

	for _, in := range inputs {
		sum := 0
		sum = sum + len(in)

		for i := 1; i < len(in); i++ {
			s := in[i:]
			count := 0

			if s[0] == in[0] {
				for j := 0; j < len(s); j++ {
					if s[j] == in[j] {
						count++
					}
				}
			}

			fmt.Println("count: ", count)
			sum = sum + count
		}

		result = append(result, int32(sum))
	}

	fmt.Println(result)
	return result
}

func main() {
	s := make([]string, 0)
	s = append(s, "ababaa")
	usernameDisparity(s)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	inputsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var inputs []string

	for i := 0; i < int(inputsCount); i++ {
		inputsItem := readLine(reader)
		inputs = append(inputs, inputsItem)
	}

	result := usernameDisparity(inputs)

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
