package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := combine(nums1, nums2)
	return medianOf(nums)
}

func combine(mis, njs []int) []int {
	lenMis, i := len(mis), 0
	lenNjs, j := len(njs), 0

	res := make([]int, lenMis+lenNjs)

	for k := 0; k < lenMis+lenNjs; k++ {
		if i == lenMis || (i < lenMis && j < lenNjs && mis[i] > njs[j]) {
			res[k] = njs[j]
			j++
			continue
		}

		if j == lenNjs || (i < lenMis && j < lenNjs && mis[i] <= njs[j]) {
			res[k] = mis[i]
			i++
		}

		fmt.Println(res)
	}
	return res
}

func medianOf(nums []int) float64 {
	l := len(nums)
	fmt.Println(nums, l)

	if l == 0 {
		panic("切片的长度为0，无法求解中位数。")
	}

	if l%2 == 0 {
		return float64(nums[l/2]+nums[l/2-1]) / 2.0
	}

	return float64(nums[l/2])
}

func findMedianSortedArrays_my(nums1 []int, nums2 []int) float64 {
	//sum := 0
	//count := 0

	/*
		for _, n := range nums1 {
			sum += n
			count++
		}

		for _, n := range nums2 {
			sum += n
			count++
		}
	*/
	m1 := 0.0
	m2 := 0.0

	if len(nums1) == 0 {
		m1 = 0.0
	} else if len(nums1) == 1 {
		m1 = float64(nums1[0])
	} else {
		m1 = float64(nums1[0]+nums1[len(nums1)-1]) / float64(2)
	}

	if len(nums2) == 0 {
		m2 = 0.0
	} else if len(nums2) == 1 {
		m2 = float64(nums2[0])
	} else {
		m2 = float64(nums2[0]+nums2[len(nums2)-1]) / float64(2)
	}

	fmt.Println(m1, m2)
	return float64(m1+m2) / float64(2.0)
}

func main() {
	n1 := []int{1, 3}
	n2 := []int{2}
	fmt.Println(findMedianSortedArrays(n1, n2))
}
