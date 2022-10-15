package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Point2 struct {
	X int
	Y int `mapstructure:"yy"` //tags for structure to serch
}

func (p Point2) method() {
	fmt.Println("call point method")
}

func main() {
	inter()
	fmt.Print("\n")
	inter2()
	fmt.Print("\n")
	inter3()
	fmt.Print("\n")
	maps()
	fmt.Print("\n")
	maps2()
	fmt.Print("\n")
	maps3()
	fmt.Print("\n")
	maps4()
	fmt.Print("\n")
}

func inter() {
	arr := []string{"a", "b", "c"}
	for i, l := range arr {
		fmt.Println(i, l)
		//fmt.Println(l)
	}
}

func inter2() {
	arr := []string{"a", "b", "c"} //short interface
	for _, l := range arr {
		fmt.Println(l)
	}
}

func inter3() {
	arr := []string{"a", "b", "c"} //short interface
	for i := range arr {
		fmt.Println(i)
	}
}

func maps() {
	pointMap := map[string]Point2{} //argumet structe
	// otherMap := map[string]int{}  //argument int
	otherPointsMap := make(map[int]Point2) //same add witout {}
	pointMap["a"] = Point2{                //hard argument, not need to control
		X: 1,
		Y: 2,
	}
	fmt.Println(pointMap)
	fmt.Println(pointMap["a"])       //show onlu struct with key "a"
	otherPointsMap[0] = Point2{1, 2} //short struct, need control arguments

	fmt.Println(otherPointsMap)
	fmt.Println(otherPointsMap[0])

	pointMap = map[string]Point2{ //argumet structe
		"b": { //hard argument, not need to control
			Y: 13,
			X: 15,
		},
	}

	fmt.Println(pointMap)

}

func maps2() {
	var oneMorePointsMap map[string]Point2
	if oneMorePointsMap == nil {
		oneMorePointsMap = map[string]Point2{
			"a": {11, 12},
		}

	}
	fmt.Println(oneMorePointsMap)
	fmt.Println(oneMorePointsMap["a"])
	oneMorePointsMap["a"].method()
	fmt.Print("\n")

	key := "b" //serch in map argument by key
	value, ok := oneMorePointsMap[key]
	if ok {
		fmt.Printf("key=%s exist in map\n", key)
		fmt.Println(value)
	} else {
		fmt.Printf("key=%s dosn't exist in map\n", key)
		fmt.Println(value)
		fmt.Println(Point2{})
	}
}

func maps3() {
	pointsMap := map[string]int{

		"x":  123, //name of index== key of structure
		"yy": 345, //can't found without tags
	}
	pointsMap["yy"] = 789

	p1 := Point2{}
	mapstructure.Decode(pointsMap, &p1)
	fmt.Println(p1)

}

func maps4() {
	pointsMap := map[string]int{

		"x": 123,
		"y": 345,
	}
	for k, v := range pointsMap {
		fmt.Println(fmt.Sprintf("[index]= %s", k), fmt.Sprintf("  argument= %d", v))

	}
}
