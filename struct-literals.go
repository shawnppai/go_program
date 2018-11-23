package main

import "fmt"

type V struct {
	X, Y int
}

var (
	v1 = V{1, 2}
	v2 = V{X: 1}
	v3 = V{}
	p  = &V{1, 2}
)

func main() {
	fmt.Println(v1, v2, v3, *p)
}
