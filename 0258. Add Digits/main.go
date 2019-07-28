package main

import "fmt"

func addDigits(num int) int {

	m := num / 10
	r := num % 10

	if r != 0 {
		num = m + r
		addDigits(num)
	}
	fmt.Println(m + r)
	return m + r

}

/*
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
*/

func main() {
	fmt.Println(65535 % 9)
	a := 65536
	addDigits(a)
}
