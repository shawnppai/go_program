package main

import "fmt"

func main() {
	names := [...]string{
		"John",
		"Sunny",
		"Jack",
		"Simmon",
	}
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	a[1] = "ShawnPai"
	fmt.Println(a, b, names)
}
