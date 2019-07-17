package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	n, p := build(numCourses, prerequisites)
	return search(n, p)
}

func build(num int, requires [][]int) (next [][]int, pre []int) {
	next = make([][]int, num)
	pre = make([]int, num)

	for _, r := range requires {
		next[r[1]] = append(next[r[1]], r[0])
		pre[r[0]]++
	}

	fmt.Println("Build: ", next, pre)
	return
}

func search(next [][]int, pre []int) []int {
	n := len(pre)
	res := make([]int, n)

	var i, j int
	// 첫 번째 완료 코스의 코드는 j입니다.
	for i = 0; i < n; i++ {
		// 전제 조건 0에서 처음 만난 과정을 수료합니다.
		for j = 0; j < n; j++ {
			if pre[j] == 0 {
				break
			}
		}
		// 각 코스는 이전 코스가 필요합니다.
		// 루프가있다.
		if j == n {
			return nil
		}

		// modify pres[j] to a negative number
		// Avoid rebuilding
		pre[j] = -1

		// After completing the j course
		// The follow-up course for j, the number of pre-requisite courses can be -1
		for _, c := range next[j] {
			pre[c]--
		}

		// 把课程代号放入答案
		res[i] = j

	}

	fmt.Println("res: ", res)
	return res
}

func main() {
	n := 4
	c := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	findOrder(n, c)
}
