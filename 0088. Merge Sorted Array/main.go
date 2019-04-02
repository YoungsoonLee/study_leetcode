package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	// append and sort
	/*
		nums1 = nums1[0:m]
		fmt.Println(nums1)

		nums1 = append(nums1, nums2...)
		sort.Ints(nums1)
		fmt.Println(nums1)
	*/

	// another way.
	temp := make([]int, 0)
	i := 0
	j := 0

	for i < len(nums1)-m {
		for j < n {
			fmt.Println(nums1[i], nums2[j])

			if nums1[i] < nums2[j] {
				temp = append(temp, nums1[i])
				i++
			} else if nums1[i] > nums2[j] {
				temp = append(temp, nums2[j])
				j++
			} else {
				temp = append(temp, nums1[i])
				temp = append(temp, nums2[j])
				i++
				j++
			}

			break
		}
	}
	fmt.Println(i, j, len(nums1)-m, n)
	if i != len(nums1)-m {
		for i < len(nums1)-m {
			temp = append(temp, nums1[i])
			i++
		}
	}

	if j != n {
		for j < n {
			temp = append(temp, nums2[j])
			j++
		}
	}

	fmt.Println(temp)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	n1 := []int{1, 2, 3, 0, 0, 0}
	n2 := []int{2, 5, 6}

	m := 3
	n := 3
	merge(n1, m, n2, n)
}
