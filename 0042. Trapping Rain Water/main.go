package main

import "fmt"

// brute-force
func trap1(height []int) int {
	ans := 0
	size := len(height) - 1

	for i := 1; i < size; i++ {
		maxLeft, maxRight := 0, 0
		for j := i; j >= 0; j-- {
			maxLeft = max(maxLeft, height[j])
		}
		for j := i; j <= size; j++ {
			maxRight = max(maxRight, height[j])
		}

		fmt.Printf("max_left: %d, max_right: %d, i: %d, ans: %d", maxLeft, maxRight, i, ans)
		fmt.Println()

		ans += min(maxLeft, maxRight) - height[i]
	}

	fmt.Println("return ans: ", ans)
	return ans
}

// dynamic
func trap2(height []int) int {
	if len(height) == 0 {
		return 0
	}

	ans := 0
	size := len(height)

	maxLeft, maxRight := make([]int, size), make([]int, size)

	maxLeft[0] = height[0]
	for i := 1; i < size; i++ {
		maxLeft[i] = max(height[i], maxLeft[i-1])
	}
	fmt.Println("maxLeft: ", maxLeft)

	maxRight[size-1] = height[size-1]
	for i := size - 2; i >= 0; i-- {
		maxRight[i] = max(height[i], maxRight[i+1])
	}
	fmt.Println("maxRight: ", maxRight)

	for i := 1; i < size-1; i++ {
		ans += min(maxLeft[i], maxRight[i]) - height[i]
	}

	fmt.Println("ans: ", ans)
	return ans
}

//
func trap3(height []int) int {
	length := len(height)
	if length <= 2 {
		return 0
	}

	left, right := make([]int, length), make([]int, length)
	left[0], right[length-1] = height[0], height[length-1]

	for i := 1; i < length; i++ {
		left[i] = max(left[i-1], height[i])
		right[length-1-i] = max(right[length-i], height[length-1-i])
	}

	fmt.Println("left: ", left)
	fmt.Println("right: ", right)

	water := 0
	for i := 0; i < length; i++ {
		water += min(left[i], right[i]) - height[i]
	}
	fmt.Println("water: ", water)
	return water

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	c := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	trap3(c)
}
