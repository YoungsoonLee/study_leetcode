package main

import (
	"fmt"
	"strings"
)

func bnf(name string) bool {
	// make graph
	g := make(map[string][]string)
	g["you"] = []string{"alice", "bob", "claire"}
	g["bob"] = []string{"anuj", "peggy"}
	g["alice"] = []string{"peggy"}
	g["anju"] = []string{"thom"}
	g["peggy"] = []string{}
	g["thom"] = []string{}
	g["jonny"] = []string{}

	// queue
	queue := make([]string, 0)
	for _, q := range g[name] {
		queue = append(queue, q)
	}

	//searched := make([]string, 0)
	searched := make(map[string]bool)

	//fmt.Println(queue, len(queue))

	for queue != nil {

		person := strings.Join(queue[0:1], "")
		queue = queue[1:]

		fmt.Println(person, queue)

		if ok := searched[person]; !ok {
			if findSeller(person) {
				fmt.Printf("%s is a mango seller\n", person)
				return true
			} else {
				if g[person] != nil {
					for _, q := range g[person] {
						if q == "" {
							return false
						}
						queue = append(queue, q)
					}
					searched[person] = true
				}

			}
		}
	}

	return false
}

//}

func findSeller(name string) bool {
	fmt.Println("find name: ", name)
	s := []string{name}
	return strings.Join(s[len(s)-1:1], "") == "m"
}

func main() {
	//var n Graph
	fmt.Println(bnf("you"))
}
