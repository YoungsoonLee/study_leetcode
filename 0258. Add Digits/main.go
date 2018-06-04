package main

import "fmt"

// mysolution
func addDigits(num int) int {
	if 0 < num && num < 10 {
		return num
	}
	res := 0
	for num >= 10 {
		r, m := int(num/10), int(num%10)
		//fmt.Println(r, m)
		res = r + m
		num = res
	}
	fmt.Println(res, num)
	return res
}

func addDigits2(num int) int {
	return (num-1)%9 + 1
}

func main() {
	a := 65536
	addDigits(a)
}
