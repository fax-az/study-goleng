package main

import "fmt"

func main() {
	//one()
	//two()
	three()
}

func one() {
	var a [5]int
	fmt.Println("emp:", a)
	if a[2] == 0 {
		fmt.Println("array have nil")
	}
}

func two() {
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	fmt.Println("len array b:", len(b))
	fmt.Println("cap array b:", cap(b))

	c := []string{"hello", "a", "3", "world"}
	fmt.Println("dcl:", c)
	fmt.Println("len array c:", len(c))
	fmt.Println("cap array c:", cap(c))

	c = append(c, "end")
	fmt.Println("dcl:", c)
	fmt.Println("len array c:", len(c))
	fmt.Println("cap array c:", cap(c)) //double cap

	c = append(c, "6", "7", "8", "9")
	fmt.Println("dcl:", c)
	fmt.Println("len array c:", len(c))
	fmt.Println("cap array c:", cap(c))

}

func three() {
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
	fmt.Println(twoD[1][1])
}
