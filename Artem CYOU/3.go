package main

import "fmt"

func main() {

	for x := 1; x < 3; x++ {
		fmt.Println("go2")
		break
		fmt.Println("go rules")
	}

	fmt.Println("go2 rules")
}
