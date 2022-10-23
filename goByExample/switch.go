package main

import (
	"fmt"
	"time"
)

func main() {
	//one()
	//two()
	//three()
	four()

}

func one() {
	var i int = 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
}

func two() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")

	}
}

func three() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")

	}
}

func four() {
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Printf("I'm a bool %v \n", t)
		case float64:
			fmt.Printf("I'm a % +6.2f \n", t)
			fmt.Printf("I'm a % 2.2f \n", t)
		default:
			fmt.Printf("Don't know type %T \n", t)
		}
	}
	whatAmI(true)
	whatAmI(1111.777)
	whatAmI("hey")

}
