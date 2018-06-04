package main

import "fmt"

func missingNumber(nums []int) int {
	sum := 0
	nsum := 0
	// total sum
	for i := 1; i <= len(nums); i++ {
		sum += i
	}
	fmt.Println(sum)

	// nums sum
	for _, n := range nums {
		nsum += n
	}

	fmt.Println(sum - nsum)
	return sum - nsum

}

func missingNumber2(nums []int) int {
	xor := 0

	for i, n := range nums {
		fmt.Println(i, n, i^n)
		xor ^= i ^ n
		fmt.Println(xor)
	}

	// 所有的 i 再加上 len(nums)，就相当于 0,1,...,n。
	// 利用 相同的数异或为0，及其交换律
	// xor 最后的值，就是那个缺失的数

	return xor ^ len(nums)
}

func main() {
	a := []int{3, 0, 1}
	missingNumber2(a)
}
