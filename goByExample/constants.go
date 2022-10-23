package main

import (
	"fmt"
	"math"
)

const s string = "constant" //can be used in any func
const p = math.Pi / 2

func main() {
	fmt.Println(s)
	one()
	two()

}

func one() {
	const n = 500000000000 // can be used only inside func
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int32(d))

}

func two() {
	fmt.Printf("sin p=%.4f", math.Sin(p))
}
