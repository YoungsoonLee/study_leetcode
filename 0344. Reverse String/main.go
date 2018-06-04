package main

import "fmt"

func reverseString(s string) string {
	bytes := []byte(s)
	fmt.Println(bytes)

	i, j := 0, len(s)-1

	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	fmt.Println(string(bytes))
	return string(bytes)
}

func main() {
	a := "hello"
	reverseString(a)
}
