package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	var name, age = "Azat", 30
	var cc = fmt.Sprintf("My name is %s. And i'm %d years old", name, age)
	fmt.Println(cc)
}
