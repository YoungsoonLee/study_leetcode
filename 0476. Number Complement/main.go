package main

import "fmt"

/*
func findComplement(num int) int {
	temp := num
	res := 0
	for temp > 0 {
		temp >>= 1
		//fmt.Println("1: ", temp)
		res <<= 1
		//fmt.Println("2: ", res)
		res++
		//fmt.Println("3: ", res)
	}
	//fmt.Println("4: ", res)
	return res ^ num
}
*/

func findComplement(num int) int {
	res := 0
	temp := num

	for temp > 0 {
		temp >>= 1 // 나누기 2
		fmt.Println("1: ", temp)

		res <<= 1 // 곱하기 2
		fmt.Println("2: ", res)

		res++
		fmt.Println("3: ", res)
	}

	fmt.Println("4: ", res)
	return res ^ num
}

func main() {
	a := 5
	fmt.Println(findComplement(a))
}
