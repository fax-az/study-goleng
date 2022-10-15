package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
	S string
}

func main() {
	adress()
	structs()

	p1 := Point{
		X: 1,
		Y: 2,
		S: "ok",
	}
	p3 := &p1
	p3.method()

	array()
	slices()
	slices2()
	slices3()
	snil()

}

func adress() {
	a := "hello world"

	fmt.Println("name=", a, "   adress=", &a)

	p := &a
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(&p)

	*p = "oh my"
	fmt.Println(a)  //parents argument
	fmt.Println(p)  //argument
	fmt.Println(*p) //parents argument
	fmt.Println(&p) //adress

	b := 42
	fmt.Println(b)
	g := &b
	*g = *g / 2
	fmt.Println(b)

}

func structs() {
	p1 := Point{
		X: 1,
		Y: 2,
		S: "ok",
	}

	fmt.Println(p1)
	fmt.Println(p1.X)
	p1.X = 123
	fmt.Println(p1)
	p2 := Point{
		X: 345,
	}
	fmt.Println(p2)

	h := &p1
	fmt.Println(h)
	fmt.Println(*h)
	fmt.Println(&h)

}

func (p *Point) method() {
	fmt.Println(p.X)
	fmt.Println(p.Y)
	fmt.Println(p.S)
}

func array() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	// a[2]="ok" - will be error array index
	fmt.Println(a)

	numbers := [...]int{1, 2, 3} //auto calc index
	fmt.Println(numbers)
}

func slices() {
	letters := []string{"a", "b", "c"}
	letters[0] = "i"
	letters = append(letters, "d", "e", "f") // add more elements
	fmt.Println(letters)

	createSlice := make([]string, 3)
	fmt.Println("\n", fmt.Sprintf("len: %d", len(createSlice)))
	fmt.Println(fmt.Sprintf("cap: %d", cap(createSlice)))
	createSlice[0] = "1"
	createSlice[1] = "2"
	createSlice[2] = "3"
	fmt.Println(createSlice)

	createSlice = append(createSlice, "4")
	fmt.Println("\n", fmt.Sprintf("len: %d", len(createSlice)))
	fmt.Println(fmt.Sprintf("cap: %d", cap(createSlice))) // when cap full, will be x2 then before
	fmt.Println(createSlice)

	createSlice = append(createSlice, "5", "6", "7", "6", "7", "6", "7")
	fmt.Println("\n", fmt.Sprintf("len: %d", len(createSlice)))
	fmt.Println(fmt.Sprintf("cap: %d", cap(createSlice))) // when cap full, will be x2 then before
	fmt.Println(createSlice)

	createSlice = append(createSlice, "5", "6", "7", "6", "7", "6", "7", "6", "7", "6", "7")
	fmt.Println("\n", fmt.Sprintf("len: %d", len(createSlice)))
	fmt.Println(fmt.Sprintf("cap: %d", cap(createSlice))) // when cap full, will be x2 then before
	fmt.Println(createSlice)                              // (for first add argument, after cap will be +argument)

}
func slices2() {
	animalsArr := [4]string{ //array
		"dog",
		"cat",
		"giraffe",
		"elephant",
	}
	fmt.Println(animalsArr)
	var a []string = animalsArr[0:2]
	fmt.Println(a)
	var b []string = animalsArr[1:3]
	fmt.Println(b)
	b[0] = "123"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(animalsArr)
}

func slices3() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := s[0:5]
	tt := s[:5] //same
	fmt.Println(t)
	fmt.Println(tt)

	ttt := s[5:10]
	tttt := s[5:] //same
	fmt.Println(ttt)
	fmt.Println(tttt)

	ttttt := s[:] //from 0 till end
	fmt.Println(ttttt)
}
func snil() {
	var snil []int
	fmt.Println(snil, fmt.Sprintf("\nlen:%d", len(snil)), fmt.Sprintf("\ncap:%d", cap(snil)))
	if snil == nil {
		fmt.Println("slice is nil")
	}
}
