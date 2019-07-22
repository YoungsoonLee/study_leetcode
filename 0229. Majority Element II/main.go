package main

func majorityElement(nums []int) []int {
	s := len(nums)
	t := s/3 + 1

	// map을 쓰면 메모리 O(n)
	// first, use map
	m := make(map[int]int)

	for i := 0; i < s; i++ {
		if _, ok := m[nums[i]]; ok {
			m[nums[i]]++
		} else {
			m[nums[i]] = 1
		}
	}

	res := []int{}
	for _, v := range m {
		if v > t {
			res = append(res, v)
		}
	}

	return res
}

func main() {

}
