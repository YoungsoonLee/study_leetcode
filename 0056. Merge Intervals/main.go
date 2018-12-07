package main

import (
	"fmt"
	"sort"
)

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
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	fmt.Println(intervals)

	res := make([]Interval, 0, len(intervals))
	temp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start <= temp.End {
			temp.End = max(temp.End, intervals[i].End)
		} else {
			res = append(res, temp)
			temp = intervals[i]
		}
	}

	res = append(res, temp)
	return res

	/*
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
	*/

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	a := []Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merge(a)
}
