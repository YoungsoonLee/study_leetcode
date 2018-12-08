package main

import (
	"fmt"
)

type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	// already sorted.

	lenA := len(intervals)
	if lenA == 0 {
		return []Interval{newInterval}
	}

	if newInterval.End < intervals[0].Start {
		return append([]Interval{newInterval}, intervals...)
	}

	if intervals[lenA-1].End < newInterval.Start {
		return append(intervals, newInterval)
	}

	res := make([]Interval, 0, lenA)
	for i := range intervals {
		if isOverlap(intervals[i], newInterval) {
			newInterval = merge(intervals[i], newInterval)

			if i == lenA-1 {
				res = append(res, newInterval)
			}

			continue
		}

		if intervals[i].End < newInterval.Start {
			res = append(res, intervals[i])
			continue
		}

		if newInterval.End < intervals[i].Start {
			fmt.Println("res 0: ", res)
			res = append(res, newInterval)
			fmt.Println("res 1: ", res)
			res = append(res, intervals[i:]...)
			fmt.Println("res 2: ", res)
			break
		}
	}

	return res
}

func isOverlap(a, b Interval) bool {
	return (a.Start <= b.Start && b.Start <= a.End) || (a.Start <= b.End && b.End <= a.End) || (b.Start <= a.Start && a.Start <= b.End)
}

func merge(a, b Interval) Interval {
	return Interval{
		Start: min(a.Start, b.Start),
		End:   max(a.End, b.End),
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	a := []Interval{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	b := Interval{4, 8}
	insert(a, b)
}
