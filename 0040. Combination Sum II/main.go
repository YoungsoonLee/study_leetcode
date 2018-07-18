package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	fmt.Println(candidates)

	res := [][]int{}
	solution := []int{}

	cs2(candidates, solution, target, &res)

	fmt.Println(res)
	return res
}

func cs2(candidates []int, solution []int, target int, res *[][]int) {
	if target == 0 {
		*res = append(*res, solution)
	}

	if len(candidates) == 0 || target < candidates[0] {
		return
	}

	//이 처리의 목적은 길이를 같게하는 슬라이스의 크기를 나중에 추가 할 때 새로운 기본 배열을 할당하게하는 것이다.
	// 기본 배열에 여러 번 수정하면 동시에 잘못된 결과가 발생합니다.
	// 다음 문을 주석으로 처리하고, 단위 테스트를 실행하고, 오류가 발생했는지 확인할 수 있습니다.
	solution = solution[:len(solution):len(solution)]
	fmt.Println("solution: ", solution)

	cs2(candidates[1:], append(solution, candidates[0]), target-candidates[0], res)

	// 후보자 [0]을 사용하지 않으면 후보자 [0]와 같은 모든 요소를 제거합니다.
	cs2(next(candidates), solution, target, res)
}

func next(candidates []int) []int {
	i := 0
	for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
		i++
	}
	//fmt.Println(candidates[i+1:])
	return candidates[i+1:]
}

func main() {
	c := []int{10, 1, 2, 7, 6, 1, 5}
	t := 8
	combinationSum2(c, t)
}
