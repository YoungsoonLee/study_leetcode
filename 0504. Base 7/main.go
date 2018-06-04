package main

import "fmt"

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	minus := ""
	if num < 0 {
		minus = "-"
		num *= -1
	}

	s := ""

	for num > 0 {
		s = fmt.Sprintf("%d", num%7) + s
		num /= 7
	}

	fmt.Println(s)
	return minus + s
}

func main() {
	a := 100
	convertToBase7(a)
}
