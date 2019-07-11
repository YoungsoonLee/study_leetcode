package main

import (
	"fmt"
)

func isHappy(n int) bool {
	slow, fast := n, trans(n)
	//fmt.Println(slow, fast)

	for slow != fast {
		fmt.Println("slow: ", slow)
		slow = trans(slow)
		fmt.Println("fast: ", fast)
		fast = trans(trans(fast)) // 2x faster
	}

	//fmt.Println(slow, fast)

	if slow == 1 {
		return true
	}
	return false
}

func trans(n int) int {
	res := 0
	for n != 0 {
		res += (n % 10) * (n % 10) // !!!
		n /= 10                    // !!!
	}

	//fmt.Println("res: ", res)
	return res
}

func main() {
	n := 1999
	fmt.Println((n % 10) * (n % 10))
	fmt.Println(isHappy(n))
}
