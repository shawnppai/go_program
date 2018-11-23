package main

import "fmt"

func main() {
	i, j := 42, 270000
	p := &i
	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println(i, j)
	*p = 21
	fmt.Println(i)
}
