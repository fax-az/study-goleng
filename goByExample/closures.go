package main

import "fmt"

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}

func intSeq() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}
