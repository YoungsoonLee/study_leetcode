package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var node int
	var cycle int
	count := 0
	ss := make([][]string, 0)
	fmt.Scanf("%d", &node)

	for i := 0; i < node; i++ {
		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			log.Printf("Failed to read: %v", scanner.Err())
			return
		}
		line := scanner.Bytes()
		s := strings.Split(string(line), " ")
		ss = append(ss, s)
	}
	fmt.Scanf("%d", &cycle)

	/*
		fmt.Println(node)
		fmt.Println(ss)
		fmt.Println(cycle)
	*/

	findCycle(ss, &count)

	if count == cycle {
		fmt.Println(count)
	} else {
		fmt.Println(0)
	}
}

func findCycle(gr [][]string, c *int) {
	for i := 0; i < len(gr[0]); i++ {
		for j := 0; j < len(gr[i][0]); j++ {
			if gr[i][j] == "1" {
				*c++
				deepFindCycle(gr, i, j, c)
			} else {
				break
			}
		}
	}
}

func deepFindCycle(gr [][]string, i, j int, c *int) {
	for k := 0; k < len(gr[j][0]); k++ {
		if gr[j][k] == "1" {
			*c++
			deepFindCycle(gr, i, j, c)
		} else {
			break
		}
	}
}
