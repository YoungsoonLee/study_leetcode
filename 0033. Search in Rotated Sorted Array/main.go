package main

import "fmt"

func search(nums []int, target int) int {
	var index, indexOfMax int
	length := len(nums)

	if length == 0 {
		return -1
	}

	// 색인 번호 변환을 수행하기 위해 최대 값의 색인 번호를 가져옵니다.
	for indexOfMax+1 < length && nums[indexOfMax] < nums[indexOfMax+1] {
		indexOfMax++
	}

	low, high, median := 0, length-1, 0
	for low <= high {
		median = (low + high) / 2
		// 인덱스 번호를 바꾼다.
		index = median + indexOfMax + 1
		if index >= length {
			index -= length
		}

		// num이 이전 슬라이스의 오름차순에서 변환된다고 가정합니다.
		// 그런 다음 old [median] == nums [index]

		// 전통적인 이진 검색 방법의 비교 판단
		// 원래 [median], 필요 nums [index]
		switch {
		case nums[index] > target:
			high = median - 1
		case nums[index] < target:
			low = median + 1
		default:
			return index
		}
	}

	return -1
}

func main() {
	n := []int{4, 5, 6, 7, 0, 1, 2}
	t := 0
	fmt.Println(search(n, t))
}
