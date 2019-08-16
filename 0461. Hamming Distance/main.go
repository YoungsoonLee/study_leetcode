package main

import (
	"fmt"
)

/*
func hammingDistance(x int, y int) int {
	x ^= y
	fmt.Println(x, y)

	res := 0
	for x > 0 {
		res += x & 1
		fmt.Println("res: ", res)
		x >>= 1
		fmt.Println("x: ", x)
	}
	return res
}
*/

func hammingDistance(x int, y int) int {
	// XOR
	// XOR은 입력 값이 다르면 1을 출력
	// 더해서 나누기 2한 것과 동일
	// XOR을 두번 하면 자신이 된다.
	x ^= y
	//x ^= y
	fmt.Println(x)

	// 결국 나누어 1카운트 인데....
	res := 0
	for x > 0 {
		res += x & 1
		fmt.Println("res: ", res)
		x >>= 1
		fmt.Println("x: ", x)
	}

	return res
}

func main() {
	a := 1
	b := 4
	//fmt.Println("main: ", 5&1)
	hammingDistance(a, b)
}
