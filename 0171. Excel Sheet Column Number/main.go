package main

import "fmt"

func titleToNumber(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {

		temp := int(s[i] - 'A' + 1)
		fmt.Println(s[i], 'A', temp)
		res = res*26 + temp
	}
	fmt.Println(res)
	return res
}

func main() {
	fmt.Println('Z')
	s := "Z"
	titleToNumber(s)
}
