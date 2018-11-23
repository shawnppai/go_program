package main

import "fmt"

func add_num(x int, y int) int {
	return x + y
}

func main() {
	c := add_num(2, 3)
	fmt.Printf("%d + %d = %d\n", 2, 3, c)
}
