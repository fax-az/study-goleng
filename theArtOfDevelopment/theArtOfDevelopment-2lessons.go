package main

import (
	"fmt"
)

func main() {
	var s6 int
	var err error
	s, s2, s3 := test()
	s4 := test1()
	s5 := cycle1()
	s6, err = cycle2()

	defer fmt.Println("first") // will be last in func main. stack LIFO
	defer fmt.Println("second")
	defer fmt.Println("third")

	if s7 := cycle1(); s7 == 20 { //short type, can't use s7 out of "if"
		fmt.Println("two")

	}

	switch s8 := cycle1(); s8 { //short case to don't use alot else, can't use s8 out of "switch"
	case 1000:
		fmt.Println("case1")
	case 20:
		fmt.Println("case2")
	default:
		fmt.Println("default")
	}

	fmt.Println(s, s2, s3, s4, s5, s6, err)
}

func test() (a, b, c int) {
	a = 1
	b = 2
	c = 3
	return a, c, b
}

func test1() (a string) {
	a = "ok"
	return a
}

func cycle1() (sum int) {
	sum = 0
	for i := 0; i < 10; i = i + 2 {
		sum += i
	}
	return sum
}

func cycle2() (int, error) {
	var sum2 int
	sum2 = 0

	for sum2 < 1000 { //shor until
		sum2 += 10
	}

	return sum2, nil //argument and nil
}
