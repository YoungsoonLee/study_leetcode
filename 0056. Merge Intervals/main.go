package main

import "fmt"

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	n := len(intervals)
	fmt.Println(n)

	var res []Interval

	if n == 1 {
		return intervals
	}

	for i := 1; i < n; i++ {
		if intervals[i-1].End >= intervals[i].Start {
			temp := Interval{intervals[i-1].Start, intervals[i].End}
			res = append(res, temp)
		} else {
			temp := Interval{intervals[i].Start, intervals[i].End}
			res = append(res, temp)
		}
	}

	fmt.Println(res)
	return res
}

func main() {
	a := []Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merge(a)
}
