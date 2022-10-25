package main

import "fmt"

func main() {
	a, b, v := vals()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(v)

	_, c, _ := vals()
	fmt.Println(c)
}

func vals() (int, int, int) {
	return 3, 7, 9
}
