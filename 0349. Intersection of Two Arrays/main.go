package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	m1 := getInts(nums1)
	m2 := getInts(nums2)

	fmt.Println("m1: ", m1)
	fmt.Println("m2: ", m2)

	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}

	fmt.Println("m1: ", m1)
	fmt.Println("m2: ", m2)

	// m1은 더 짧은 사전이며 더 빨라질 것입니다.
	for n := range m1 {
		if m2[n] {
			res = append(res, n)
		}
	}

	fmt.Println("res: ", res)
	return res
}

// 중복 값을 정리하고 쿼리하기 쉽습니다.
func getInts(a []int) map[int]bool {
	res := make(map[int]bool, len(a))

	for i := range a {
		res[a[i]] = true
	}

	return res
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	intersection(nums1, nums2)
}
