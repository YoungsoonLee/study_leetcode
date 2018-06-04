package main

import "fmt"

func findRestaurant(list1 []string, list2 []string) []string {
	if len(list1) > len(list2) {
		list1, list2 = list2, list1
	}

	m2 := make(map[string]int, len(list2))
	for i := range list2 {
		m2[list2[i]] = i
	}

	min := len(list1) + len(list2)
	res := make([]string, 0, 1000)

	for i, r := range list1 {
		if j, ok := m2[r]; ok {
			fmt.Println(min, i, j, i+j, res)
			if min == i+j {
				res = append(res, r)
			}
			if min > i+j {
				min = i + j
				res = append(res[len(res):], r) // !!!!
			}
		}
	}

	fmt.Println(res)
	return res
}

func findRestaurant_my(list1 []string, list2 []string) []string {

	if len(list1) > len(list2) {
		list1, list2 = list2, list1
	}

	l1m := make(map[string]int, len(list1))
	l2m := make(map[string]int, len(list2))

	for i, w := range list1 {
		l1m[w] = i
	}

	for i, w := range list2 {
		l2m[w] = i
	}

	fmt.Println(l1m, l2m)
	sum := len(list1) + len(list2)

	for k, v := range l1m {
		v2, ok := l2m[k]
		if ok {
			sum = min(sum, v+v2)
		}
	}

	fmt.Println(sum, list2[sum])
	return []string{list2[sum]}
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func main() {
	s1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	s2 := []string{"KFC", "Shogun", "Burger King"}

	findRestaurant(s1, s2)
}
