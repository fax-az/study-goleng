package main

import "fmt"

type Point4 struct {
	X, Y int
}

func main() {
	sumCallback := func(n1, n2 int) int {
		return n1 + n2
	}
	result := doSomething(sumCallback, "plus")
	fmt.Println(result)

	multipleCallback := func(n1, n2 int) int {
		return n1 * n2
	}

	result = doSomething(multipleCallback, "multiple")
	fmt.Println(result)

	orderPrice := totalPrice(1) //sum=1
	fmt.Print("\n")
	fmt.Println(orderPrice(1))  //sum=1+1
	fmt.Println(orderPrice(1))  //sum=2+1
	fmt.Println(orderPrice(10)) //sum=3+10

	fmt.Print("\n")
	p := Point4{1, 2}
	fmt.Println(p)

	//fmt.Print("\n")
	// fmt.Println(movePoint(p, 1, 1)) //add copy of original p
	// fmt.Println(p)

	// fmt.Print("\n")
	// fmt.Println(p.movePoint2(1, 1))
	// fmt.Println(p)

	// fmt.Print("\n")
	// p.movePointPtr2(1, 1) //move original p by adress
	// fmt.Println(p)

	fmt.Print("\n")
	v := &p //v it's link to p
	v.movePoint2(1, 1)
	fmt.Println("move p, by v", p)
	v.movePointPtr2(1, 1)
	fmt.Println("move original p, by v", p)
	fmt.Print(v)

	// fmt.Print("\n")
	// p.movePointPtr2(1, 1)
	// movePointPtr(&p, 1, 1) //move original p by adress
	// fmt.Println(p)
}

func doSomething(callback func(int, int) int, s string) int {
	resulte := callback(5, 1)
	fmt.Println(s)
	return resulte

}

func totalPrice(initPrice int) func(int) int {
	sum := initPrice
	return func(x int) int {
		sum += x
		return sum
	}
}

// func movePoint(p Point4, x, y int) Point4 { //add new copy of point
// 	p.X += x
// 	p.Y += y
// 	return p
// }

// func movePointPtr(p *Point4, x, y int) {
// 	p.X += x
// 	p.Y += y
// }

func (p Point4) movePoint2(x, y int) Point4 { //add new copy of point
	p.X += x
	p.Y += y
	return p
}

func (p *Point4) movePointPtr2(x, y int) { //same like movePointPtr. Short ver
	p.X += x
	p.Y += y
}
