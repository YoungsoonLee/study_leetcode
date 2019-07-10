package main

import "fmt"

func majorityElement(nums []int) int {
	// 제목에 따라 len [nums]> 0이고 n / 2보다 큰 발생 수가 존재합니다.
	x, t := nums[0], 1

	for i := 1; i < len(nums); i++ {
		switch {
		case x == nums[i]:
			t++
		case t > 0:
			t--
		default:
			// 이제 x! = nums [i] 및 t == 0
			// 내가 짝이 있어야한다는 것을 알아라.
			// nums에서 가장 눈에 띄는 요소가 z이고 해당 항목이 zn> n / 2라고 가정합니다.
			// nums [: i]에서 z <= i / 2의 발생 횟수
			// 그런 다음, nums [i :]에서 z> = zn - i / 2> n / 2 - i / 2 = (n-i) / 2의 발생 횟수
			// nums [i :]의 z 발생 횟수가 여전히 len (nums [i :]) / 2를 초과합니다.
			x = nums[i]
			t = 1
		}
	}

	return x

	/* my solution */
	/*
		m := make(map[int]int)
		for _, v := range nums {
			if _, ok := m[v]; ok {
				m[v]++
			} else {
				m[v] = 1
			}
		}

		p := len(nums) / 2

		for k, v := range m {
			if v > p {
				return k
			}
		}

		return 0
	*/
}

func main() {
	//fmt.Println(7 / 2)
	//fmt.Println(3 / 2)

	n := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(n))
}
