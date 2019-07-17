package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 다익스트라 최단 거리 알고리즘. heap을 이용
// N 노드 개수, K 시작 포인트
func networkDelayTime(times [][]int, N int, K int) int {
	h := &IntHeap{K}
	heap.Init(h)

	timeMap := make(map[int]map[int]int)

	for _, t := range times {
		if timeMap[t[0]] == nil {
			timeMap[t[0]] = make(map[int]int)
		}
		timeMap[t[0]][t[1]] = t[2]
	}

	fmt.Println("init map: ", timeMap)

	minTime := make([]int, N+1)
	for i := 0; i < N+1; i++ {
		minTime[i] = -1 // 초기화
	}
	minTime[K] = 0 // 출발지는 0

	fmt.Println("minTime: ", minTime)

	for h.Len() > 0 {
		top := heap.Pop(h).(int)
		fmt.Println("top: ", top)

		for d, t := range timeMap[top] {
			fmt.Printf("d: %d, t: %d\n", d, t)
			cost := minTime[top] + t
			if minTime[d] == -1 || minTime[d] > cost {
				minTime[d] = cost
				heap.Push(h, d)
			}
		}
	}

	mx := 0
	for i := 1; i < N+1; i++ {
		if minTime[i] == -1 {
			return -1
		}
		if mx < minTime[i] {
			mx = minTime[i]
		}
	}

	return mx
}

func main() {
	times := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	N := 4
	K := 2
	networkDelayTime(times, N, K)
}
