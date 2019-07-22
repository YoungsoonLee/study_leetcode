package main

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	res := []string{}
	l := len(nums)

	if l == 0 {
		return res
	}

	begin := nums[0]
	str := ""

	for i := 0; i < l; i++ {
		if i == l-1 || nums[i]+1 != nums[i+1] {
			if nums[i] == begin {
				str = fmt.Sprintf("%d", begin)
			} else {
				str = fmt.Sprintf("%d->%d", begin, nums[i])
			}

			if i+1 < l {
				begin = nums[i+1]
			}

			res = append(res, str)
		}
	}

	fmt.Println(res)
	return res
}

/*
func summaryRanges(nums []int) []string {
	res := make([]string, 0)

	if len(nums) == 0 {
		return []string{""}
	}

	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}

	count := 0
	m := make(map[int]int, 0)

	var j int
	for i := 0; i < len(nums)-1; i++ {
		j = i + 1
		if nums[j]-nums[i] == 1 {
			if j == len(nums)-1 {
				m[nums[j]] = nums[j-1]
			}
			count++
		} else {
			m[nums[i]-count] = nums[i]
			count = 0
			//fmt.Println(j, len(nums))
			if j == len(nums)-1 {
				m[nums[j]] = nums[j]
			}
		}
	}

	for k, v := range m {
		if k == v {
			res = append(res, strconv.Itoa(v))
		} else {
			res = append(res, strconv.Itoa(k)+"->"+strconv.Itoa(v))
		}
	}
	fmt.Println(res)
	return res
}
*/

func main() {
	n := []int{0, 1, 2, 4, 5, 7}
	summaryRanges(n)
}
