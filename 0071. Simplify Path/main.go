package main

import (
	"fmt"
	"strings"
)

type Stack struct {
	value []string
}

func simplifyPath(path string) string {
	lp := len(path)
	stack := make([]string, 0, lp/2)
	dir := make([]byte, 0, lp)

	for i := 0; i < lp; i++ {
		// 사용하기 전에 dir을 지운다.
		// 이 방법은 dir = [] byte {}보다 거의 8 배 빠릅니다.
		dir = dir[:0]
		fmt.Println("dir1: ", dir)

		// dir을 얻는다.
		for i < lp && path[i] != '/' {
			dir = append(dir, path[i])
			i++
		}

		fmt.Println("dir: ", dir)

		s := string(dir)
		fmt.Println("s: ", s)

		switch s {
		case ".", "":
			//do nothing
		case "..":
			if len(stack) > 0 {
				// pop
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, s)
		}
	}

	fmt.Println("stack: ", stack)
	return "/" + strings.Join(stack, "/")
}

/*
func simplifyPath(path string) string {
	// using stack
	stack := new(Stack)

	// init stack adding root
	// stack.value = append(stack.value, "/")
	ts := ""

	for i := 1; i < len(path)-1; i++ {
		cu := string(path[i])
		ne := string(path[i+1])

		if cu == "/" {
			//fmt.Println("/", ts)
			add(stack, ts)
			ts = ""
			//continue
		} else if cu == "." {
			if ne != "." {
				continue
			} else if ne == "." {
				//remove
				remove(stack)
			}
		} else {
			if ne != "." || ne != "/" {
				ts = ts + cu
			} else {
				continue
			}
		}
		fmt.Println(ts)
		--*
			cu := string(path[i])
			ne := string(path[i+1])

			if cu == "/" {
				continue
			} else if cu == "." && ne != "." {
				continue
			} else if cu == "." && ne == "." {
				//remove
				remove(stack)
			} else {
				// if character
				// add to stack
				add(stack, cu)
			}
		--/
	}

	// chack last
	add(stack, ts)

	// return stack
	fmt.Println(stack.value, len(stack.value))
	result := "/" + strings.Join(stack.value, "/")
	return result
}

func remove(s *Stack) {
	if len(s.value) != 0 {
		s.value = s.value[0 : len(s.value)-1]
	}

}

func add(s *Stack, c string) {
	fmt.Println("add: ", c)
	s.value = append(s.value, c)
}
*/

func main() {
	s := "/home//foo/"
	fmt.Println(simplifyPath(s))
}
