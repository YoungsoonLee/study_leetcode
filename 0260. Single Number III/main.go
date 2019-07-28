package main

func singleNumber(nums []int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			m[nums[i]]++
		} else {
			m[nums[i]] = 1
		}
	}

	res := []int{}

	for k, v := range m {
		if v == 1 {
			res = append(res, k)
		}
	}

	return res
}

func main() {

}
