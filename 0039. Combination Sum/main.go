package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	res := [][]int{}
	solution := []int{}

	cs(candidates, solution, target, &res)

	fmt.Println("res: ", res)
	return res
}

func cs(candidates, solution []int, target int, result *[][]int) {
	if target == 0 {
		*result = append(*result, solution)
	}

	if len(candidates) == 0 || target < candidates[0] {
		return
	}

	// 이 처리의 목적은 길이를 같게하는 슬라이스의 크기를 나중에 추가 할 때 새로운 기본 배열을 할당하게하는 것이다.
	// 기본 배열에 여러 번 수정하면 동시에 잘못된 결과가 발생합니다.
	// 다음 문을 주석으로 처리하고, 단위 테스트를 실행하고, 오류가 발생했는지 확인할 수 있습니다.
	// !!!
	solution = solution[:len(solution):len(solution)]
	fmt.Println("solution: ", solution)

	// !!!
	cs(candidates, append(solution, candidates[0]), target-candidates[0], result)

	cs(candidates[1:], solution, target, result)
}

func main() {
	c := []int{2, 3, 6, 7}
	t := 7
	combinationSum(c, t)
}
