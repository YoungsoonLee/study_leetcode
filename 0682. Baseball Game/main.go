package main

import (
	"fmt"
	"strconv"
)

/*
func calPoints(ops []string) int {
	pointStack := make([]int, 0, len(ops))
	fmt.Println(pointStack)

	for i := range ops {
		switch ops[i] {
		case "+":
			r1 := pointStack[len(pointStack)-1]
			r2 := pointStack[len(pointStack)-2]
			pointStack = append(pointStack, r1+r2)
		case "D":
			r1 := pointStack[len(pointStack)-1]
			pointStack = append(pointStack, 2*r1)
		case "C":
			pointStack = pointStack[:len(pointStack)-1]
		default:
			point, _ := strconv.Atoi(ops[i])
			pointStack = append(pointStack, point)
		}
	}

	res := 0
	fmt.Println(pointStack)
	for _, p := range pointStack {
		res += p
	}
	return res
}

func calPoints_my(ops []string) int {
	p := []int{}
	sum := 0

	for i := range ops {
		switch ops[i] {
		case "C":
			tp := p[len(p)-1]
			p = p[:len(p)-1]
			sum -= tp
		case "D":
			tp := p[len(p)-1] * 2
			//p[len(p)-1] = tp
			p = append(p, tp)
			sum += tp
		case "+":
			s1 := p[len(p)-2]
			s2 := p[len(p)-1]
			fmt.Println(s1, s2)
			p = append(p, s1+s2)
			sum += s1 + s2
		default:
			nc, _ := strconv.Atoi(ops[i])
			p = append(p, nc)
			sum += nc
		}
		fmt.Println(p, sum)
	}

	return sum
}
*/

func calPoints(ops []string) int {
	stack := make([]int, 0)
	sum := 0
	pre := 0

	for _, v := range ops {

		if len(stack) > 0 {
			pre = stack[len(stack)-1]
		}

		switch v {
		case "C":
			sum -= pre
			stack = stack[:len(stack)-1]
		case "D":
			sum += pre * 2
			stack = append(stack, pre*2)
		case "+":
			pre_b := stack[len(stack)-2]
			sum += pre + pre_b
			stack = append(stack, pre+pre_b)
		default:
			n, _ := strconv.Atoi(v)
			//fmt.Println(pre, n)
			sum += n
			//fmt.Println(sum)
			stack = append(stack, n)
		}

		fmt.Println(stack, sum)
	}

	//fmt.Println(stack)
	fmt.Println(sum)
	return sum
}

func main() {
	a := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	calPoints(a)
}
