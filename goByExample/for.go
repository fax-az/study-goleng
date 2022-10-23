package main

import "fmt"

func main() {

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 { //divisible by two
			continue
		}
		fmt.Println(n)
	}

}
