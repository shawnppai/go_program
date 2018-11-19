package main

import "fmt"

func main() {
	num := 4

	if num > 3 {
		fmt.Println("Many")
	}

	if num == 1 {
		fmt.Println("one")
	} else if num == 2 {
		fmt.Println("two")
	} else {
		fmt.Println("Many")
	}

	switch {
	case num == 1:
		fmt.Println("one")
	case num == 2:
		fmt.Println("two")
	case num > 2:
		fmt.Println("Many")
	default:
		fmt.Println("Thrown over boat")
	}
}