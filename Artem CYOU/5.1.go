package main

import (
	"fmt"
)

// func double(a int) {
// 	a *= 2
// 	fmt.Println(a)

// }

func cPointer(myPointer *int) {
	fmt.Println(*myPointer)

}

func main() {

	var number int = 65
	cPointer(&number)

	// var number int = 6
	// fmt.Println(number)

	// var a *int = &number
	// *a = 8
	// //double(number)

	// fmt.Println(a)
	// fmt.Println(*a)
	// fmt.Println(number)
	// fmt.Println(reflect.TypeOf(number))
	// //fmt.Println(reflect.TypeOf(a))
	// //	fmt.Println(&number)

}
