package main

import "fmt"

func compress(chars []byte) int {
	// in-space
	// already sorted? or not.
	// assume sorted.
	c := 1
	for i := 1; i < len(chars); i++ {
		if chars[i-1] == chars[i] {
			c++
		} else {
			// replace
			chars[i-1] = byte(c)
			// init c
			c = 1
		}
	}

	//chars[len(chars)-1] = byte(c)

	fmt.Println(chars[:len(chars)-1])
	return len(chars)
}

/*
func compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}

	cu := 0
	fa := 1
	ret := make([]byte, 0) // not in-memory
	count := 1

	for i := 0; i < len(chars); i++ {
		if fa < len(chars) && chars[cu] == chars[fa] {
			count++
		} else {
			ret = append(ret, chars[cu])
			if count > 10 {
				ret = append(ret, byte(count/10), byte(count%10))
			} else {
				ret = append(ret, byte(count))
			}
			count = 1
		}
		cu++
		fa++
	}

	fmt.Println(ret, len(ret))
	return len(ret)
}

func compress2(chars []byte) int {
	end, count := 0, 1

	for i, char := range chars {
		if i+1 < len(chars) && char == chars[i+1] {
			count++
		} else {
			chars[end] = char
			end++

			if count > 1 {
				for _, num := range strconv.Itoa(count) {
					chars[end] = byte(num)
					end++
				}
			}

			count = 1
		}

	}
	fmt.Println(chars, end)
	return end
}
*/

func main() {
	a := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	compress(a)
}
