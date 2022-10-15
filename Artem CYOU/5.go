package main

import "fmt"

// func other(a, b int) (int, error) {
// 	return a * b

// }

func summcBack(item, cBack float64) (float64, error) {
	if item < 0 {
		return 0, fmt.Errorf("item %.2f, can't be less then 0", item)
	}

	if cBack < 0 {
		return 0, fmt.Errorf("item %.2f, can't be less then 0", cBack)
	}
	return item * cBack / 100, nil
	// fmt.Printf("cBack: %.2f cent \n", summ)

}

func main() {

	// a := other(2, 4)
	// fmt.Println(a)
	var total, summ float64
	var err error

	summ, err = summcBack(-1000, 5.7)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("cBack: %.2f cent \n", summ)
	}

	total += summ

	summ, err = summcBack(786, 5.7)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("cBack: %.2f cent \n", summ)
	}
	total += summ

	summ, err = summcBack(1999, 5.7)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("cBack: %.2f cent \n", summ)
	}
	total += summ

	fmt.Printf("Total cBack: %.2f cent \n", total)

	// item := 1000
	// cBack := 5.7
	// summ := item * cBack / 100

	// item = 786
	// cBack = 5.7
	// summ2 := float64(item) * cBack / 100

	// fmt.Printf("cBack 1 item: %.2f cent \ncBack 2 item: %.2f cent", summ, summ2)

	//fmt.Println(summ, summ2)
}
