package main

import "container/heap"

// entry. priorityQueue
type entry struct {
	key      string
	value    int
	maxIndex int
	minIndex int
}

type minPQ []*entry

func (pq minPQ) Len() int { return len(pq) }
func (pq minPQ) Less(i, j int) bool {
	return pq[i].value < pq[j].value
}

func (pq minPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].minIndex = i
	pq[j].minIndex = j
}

func (pq *minPQ) Push(x interface{}) {
	n := len(*pq)
	entry := x.(*entry)
	entry.minIndex = n
	*pq = append(*pq, entry)
}

func (pq *minPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	entry := old[n-1]
	entry.minIndex = -1
	*pq = old[0 : n-1]
	return entry
}

type maxPQ []*entry

func (pq maxPQ) Len() int { return len(pq) }

func (pq maxPQ) Less(i, j int) bool {
	return pq[i].value > pq[j].value
}

func (pq maxPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].maxIndex = i
	pq[j].maxIndex = j
}

// Push 往 pq 中放 entry
func (pq *maxPQ) Push(x interface{}) {
	n := len(*pq)
	entry := x.(*entry)
	entry.maxIndex = n
	*pq = append(*pq, entry)
}

// Pop 从 pq 中取出最优先的 entry
func (pq *maxPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	entry := old[n-1]
	entry.maxIndex = -1 // for safety
	*pq = old[0 : n-1]
	return entry
}

type AllOne struct {
	m map[string]*entry

	max maxPQ
	min minPQ
}

func Constructor() AllOne {
	return AllOne{
		m:   make(map[string]*entry),
		max: maxPQ{},
		min: minPQ{},
	}
}

func (a *AllOne) Inc(key string) {
	if ep, ok := a.m[key]; ok {
		ep.value++
		heap.Fix(&a.max, ep.maxIndex)
		heap.Fix(&a.min, ep.maxIndex)
	} else {
		ep = &entry{key: key, value: 1}
		a.m[key] = ep
		heap.Push(&a.max, ep)
		heap.Push(&a.min, ep)
	}
}

func (a *AllOne) Dec(key string) {
	if ep, ok := a.m[key]; !ok {
		return
	} else if ep.value == 1 {
		heap.Remove(&a.max, ep.maxIndex)
		heap.Remove(&a.min, ep.minIndex)
	} else {
		ep.value--
		heap.Fix(&a.max, ep.maxIndex)
		heap.Fix(&a.min, ep.minIndex)
	}
}

func (a *AllOne) GetMaxKey() string {
	if len(a.max) == 0 {
		return ""
	}
	return a.max[0].key
}

func (a *AllOne) GetMinKey() string {
	if len(a.min) == 0 {
		return ""
	}
	return a.min[0].key
}

func main() {

}
