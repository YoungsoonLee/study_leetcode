package main

import "fmt"

func search(nums []int, target int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}
	k := 1
	for k < len(nums) && nums[k-1] <= nums[k] {
		k++
	}
	fmt.Println(k)

	// binary search
	i, j := 0, length-1
	for i <= j {
		m := (i + j) / 2
		med := (m + k) % length
		fmt.Println("med: ", med)

		switch {
		case nums[med] < target:
			i = m + 1
		case nums[med] > target:
			j = m - 1
		default:
			return true
		}
	}

	return false

	/* mine with hash
	hash := make(map[int]bool, len(nums))

	for _, i := range nums {
		if hash[i] == false {
			hash[i] = true
		}
	}

	fmt.Println(hash)
	if hash[target] {
		return true
	}

	return false
	*/
}

func main() {
	n := []int{2, 5, 6, 0, 0, 1, 2}
	t := 3
	fmt.Println(search(n, t))
}
