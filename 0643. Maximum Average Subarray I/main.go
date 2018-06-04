package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	temp := 0
	// 첫 번째 임시 변수 가져 오기
	for i := 0; i < k; i++ {
		temp += nums[i]
	}

	max := temp

	for i := k; i < len(nums); i++ {
		// i를 기반으로 임시 직원을 얻습니다.
		// fmt.Println(i, nums[i-k], nums[i], temp)
		temp = temp - nums[i-k] + nums[i] // !!!!!

		if max < temp {
			max = temp
		}
	}

	//fmt.Println(temp, max)
	//fmt.Println(float64(max) / float64(k))
	return float64(max) / float64(k)
}

func findMaxAverage_my(nums []int, k int) float64 {
	maxAvg := 0.0
	l := len(nums)
	s := (l - k) + 1
	//c := 1
	sum := 0

	fmt.Println(l, s)
	temp := make([]int, k)

	for i := 0; i < s; i++ {
		//fmt.Println(nums[l-k-i : l-i])
		temp = nums[l-k-i : l-i]
		fmt.Println(temp)

		for j := 0; j < len(temp); j++ {
			sum += temp[j]
		}

		if maxAvg < float64(sum/k) {
			maxAvg = float64(sum) / float64(k)
		}

		fmt.Println("sum: ", sum, float64(sum)/float64(k), maxAvg)
		sum = 0
	}
	//fmt.Println(maxAvg)
	return maxAvg
}

func main() {
	a := []int{1, 12, -5, -6, 50, 3}
	k := 4
	findMaxAverage(a, k)
}
