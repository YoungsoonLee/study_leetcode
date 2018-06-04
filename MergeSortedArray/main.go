package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	//update list
	nums1 = nums1[:m]
	nums2 = nums2[:n]
	//fmt.Println(nums1, nums2)

	//expand nums1
	nums1 = append(nums1, make([]int, len(nums2))...)
	fmt.Println(nums1, nums2)

	// compare

}

func main() {
	a := []int{1, 3, 4, 5, 6}
	b := []int{2, 7, 8, 9, 10}
	merge(a, 5, b, 5)
}
