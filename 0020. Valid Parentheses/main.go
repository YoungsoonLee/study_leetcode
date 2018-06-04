package main

import "fmt"

func isValid(s string) bool {
	stack := make([]string, 0)
	opa := make(map[string]string)
	opa["("] = ")"
	opa["{"] = "}"
	opa["["] = "]"

	cpa := make(map[string]string)
	cpa[")"] = ")"
	cpa["}"] = "}"
	cpa["]"] = "]"

	for _, c := range s {
		ov, ook := opa[string(c)]
		if string(c) != ov && ook { //open
			stack = append(stack, ov)
		}

		cv, cok := cpa[string(c)]
		if cok { //close
			if len(stack) == 0 {
				return false
			}

			if cv != stack[len(stack)-1] {
				return false
			} else {
				stack = stack[0 : len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}

func main() {
	s := "()"
	fmt.Println(isValid(s))
}
