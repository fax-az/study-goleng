package main

import "fmt"

func main() {
	//one()
	two()
}

func one() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	m["k3"] = 12
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)
	fmt.Println("len:", len(m))
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}

	commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}
	fmt.Println(commits)

	fmt.Println("map:", m)
	m = map[string]int{}
	fmt.Println("map:", m) //initialize an empty map
}

func two() {
	type Node struct {
		Next  *Node
		Value interface{}
	}
	var first *Node

	visited := make(map[*Node]bool)
	for n := first; n != nil; n = n.Next {
		if visited[n] {
			fmt.Println("cycle detected")
			break
		}
		visited[n] = true
		fmt.Println(n.Value)
	}

}
