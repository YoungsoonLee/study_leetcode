package main

import (
	"fmt"
	"sort"
	"strconv"
)

// athlete는 sort.Interface 인터페이스를 구현합니다.
type athlete struct {
	sorce, index int
}
type athletes []athlete

//sort.Interface
func (a athletes) Len() int           { return len(a) }
func (a athletes) Less(i, j int) bool { return a[i].sorce > a[j].sorce }
func (a athletes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func findRelativeRanks(nums []int) []string {
	res := make([]string, len(nums))
	as := make(athletes, len(nums))

	for i := range nums {
		as[i] = athlete{sorce: nums[i], index: i}
	}

	fmt.Println(as)

	sort.Sort(as)
	fmt.Println(as)

	for i, a := range as {
		switch i {
		case 0:
			res[a.index] = "Gold Medal"
		case 1:
			res[a.index] = "Silver Medal"
		case 2:
			res[a.index] = "Bronze Medal"
		default:
			res[a.index] = strconv.Itoa(i + 1)
		}
	}

	fmt.Println(res)
	return res
}

func main() {
	a := []int{10, 3, 8, 9, 4}
	findRelativeRanks(a)
}
