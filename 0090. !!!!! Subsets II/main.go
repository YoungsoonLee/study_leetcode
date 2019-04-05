package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	//fmt.Println(nums)

	var dfs func(int, []int)
	dfs = func(idx int, temp []int) {
		fmt.Println(idx, temp)

		t := make([]int, len(temp))
		copy(t, temp) // copy(dest, src)

		// 위의 두 줄이 없으면 대답이 잘못되었습니다.
		// 재귀 프로세스 동안 temp의 기본 배열이 지속적으로 수정되기 때문에
		// 프로그램이 끝날 때 temp의 기본 배열 값은 모두 num의 최대 값입니다.
		res = append(res, t)

		for i := idx; i < len(nums); i++ {
			if i == idx || nums[i] != nums[i-1] {
				dfs(i+1, append(temp, nums[i]))
			}
		}
	}

	// call dfs
	temp := make([]int, 0, len(nums))
	fmt.Println("first temp: ", temp)
	dfs(0, temp)

	return res
}

func main() {
	n := []int{1, 2, 3}
	fmt.Println(subsetsWithDup(n))
}
