package main

import (
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	sort.Ints(nums)

	max, temp := 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[i]-1 == nums[i-1] {
			temp++
		} else if nums[i] != nums[i-1] {
			// [0,1,1,2]　仍然可以视为连续
			temp = 1
		}
		//　更新　max
		if max < temp {
			max = temp
		}
	}

	return max
}

/** mine **/
/*
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// sort
	a := unique(nums)
	sort.Ints(a)
	fmt.Println(a)

	s, c := 0, 0
	for i := 0; i < len(a)-1; i++ {
		next := i + 1
		if a[i]+1 == a[next] {
			c = max(c, c+1)
		} else {
			s = max(s, c)
			c = 0
		}

	}

	fmt.Println(s+1, c+1)
	return max(s+1, c+1)
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
*/

func main() {
	n := []int{1, 2, 0, 1}
	//n := []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}
	longestConsecutive(n)
}
