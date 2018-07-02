package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	q := make([]int, 0)
	m := make(map[int]int, 0)
	re := make([]int, 0)

	for {
		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			log.Printf("Failed to read: %v", scanner.Err())
			return
		}
		line := scanner.Bytes()
		s := strings.Split(string(line), " ")

		if s[0] == "evict" {
			evict(&q, m)
		}

		if s[0] == "add" {
			key, _ := strconv.Atoi(s[1])
			value, _ := strconv.Atoi(s[2])
			addQ(key, &q)
			add(key, value, m)
		}

		if s[0] == "get" {
			key, _ := strconv.Atoi(s[1])
			addQ(key, &q)
			get(key, m, &re)
		}

		if s[0] == "remove" {
			key, _ := strconv.Atoi(s[1])
			remove(key, m, &re)
		}

		if s[0] == "exit" {
			break
		}
	}

	for _, r := range re {
		fmt.Println(r)
	}
}

func addQ(key int, q *[]int) {
	for i := 0; i < len(*q); i++ {
		if (*q)[i] == key {
			*q = (*q)[i+1:]
		}
	}
	*q = append(*q, key)
}

func evict(q *[]int, m map[int]int) {
	delete(m, (*q)[0])
	*q = (*q)[1:]
}

func add(key, value int, m map[int]int) {
	m[key] = value
}

func get(key int, m map[int]int, re *[]int) {
	if v, ok := m[key]; ok {
		*re = append(*re, v)
	} else {
		*re = append(*re, -1)
	}
}

func remove(key int, m map[int]int, re *[]int) {
	if v, ok := m[key]; ok {
		*re = append(*re, v)
		delete(m, key)
	} else {
		log.Fatal("not exists key")
	}
}
