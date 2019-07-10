package main

import (
	"fmt"
	"sort"
	"strconv"
)

type bytes [][]byte

func largestNumber(nums []int) string {
	size := len(nums)
	b := make(bytes, size)
	resSize := 0

	for i := range b {
		b[i] = []byte(strconv.Itoa(nums[i]))
		resSize += len(b[i])
	}

	fmt.Println(b)
	sort.Sort(b)
	fmt.Println(b)

	return ""
}

func main() {
	// fmt.Println([]byte(strconv.Itoa(2))) //
	// fmt.Println(byte('2')) 왜 스트링 2를 바이트 배열로 변한 하면 50 이지?? 바이트 변환을 알아야 되.
	n := []int{10, 2}
	largestNumber(n)
}
