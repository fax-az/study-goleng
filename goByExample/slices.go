package main

import "fmt"

func main() {
	//one()
	two()
}

func one() {
	s := make([]string, 3)
	fmt.Println("emp:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("get:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	c = append(c, "d")
	c = append(c, "e", "f")
	fmt.Println("apd:", c)

	l := c[2:5]
	fmt.Println("sl1:", l)

	l = c[:5]
	fmt.Println("sl2:", l)

	l = c[2:]
	fmt.Println("sl3:", l)

}

func two() {
	t := []string{"1", "2", "3"}
	fmt.Println("dls:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
