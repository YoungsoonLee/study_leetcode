package main

import "fmt"

func main() {
	//fmt.Printf("Hello")
	a := []string{"donut", "juice", "apple", "banana"}
	LRU(3, a) // 20

	b := []string{"donut", "juice", "donut"}
	LRU(3, b) // 11

	c := []string{"donut", "juice", "apple", "donut"}
	LRU(3, c) // 16

	d := []string{"donut", "juice", "apple", "donut"}
	LRU(0, d) // 20

	e := []string{"donut", "juice", "apple", "donut", "apple", "coffee"}
	LRU(5, e) // 22

	f := []string{"donut", "juice", "juice", "apple", "donut"}
	LRU(2, f) // 21
}

func LRU(initSize int, arr []string) int {
	hitDelay := 1
	missDelay := 5
	curSize := 0
	result := 0

	//cache := make(map[string]bool, initSize)
	cache := make([]string, 0)
	//memoized
	m := make(map[string]bool, 0)

	for _, a := range arr {
		if initSize == 0 {
			result += missDelay
		} else if curSize < initSize {
			if _, exists := m[a]; exists {
				// add hitdelay to result
				result += hitDelay
			} else {
				cache = append(cache, a)
				curSize++
				m[a] = true
				result += missDelay
			}

		} else {
			// over the size
			if _, exists := m[a]; exists {
				result += hitDelay
			} else {
				//delete map
				delete(m, cache[0])
				//delete cache
				cache = cache[1:]
				//insert cache
				cache = append(cache, a)

				m[a] = true
				result += missDelay
			}
		}
	}

	fmt.Println(cache)
	fmt.Println(m)
	fmt.Println(result)
	return result
}

/*
# // LRU(Least Recently Used) Cache
#
# // Cache Hit Delay: 1
# // Cache Miss Delay: 5
#
# // Cache Size: 3
# // Data Strings: ['donut', 'juice', 'apple', 'banana']
# // Output Value: 20
#
# // Cache Size: 3
# // Data Strings: ['donut', 'juice', 'donut']
# // Output Value: 11
#
# // Cache Size: 3
# // Data Strings: ['donut', 'juice', 'apple', 'donut']
# // Output Value: 16
#
# // Cache Size: 0
# // Data Strings: ['donut', 'juice', 'apple', 'donut']
# // Output Value: 20
#
# // Cache Size: 5
# // Data Strings: ['donut', 'juice', 'apple', 'donut', 'apple', 'coffee']
# // Output Value: 22
#
# // Cache Size: 2
# // Data Strings: ["donut", "juice", "juice", "apple", "donut"]
# // Output: 21

func lruCache(initial_cache_size, array_data_strings):
  # TODO
  pass
*/
