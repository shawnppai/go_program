package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "World"
	fmt.Println(a)

	primes := [6]int{2, 3, 4, 5, 6, 7}
	fmt.Println(primes)
	p := &[3]string{"a", "b", "c"}
	fmt.Println(p[1])
	fmt.Println(*p)
	p[0] = "GO GO GO"
	fmt.Println(p)
	fmt.Println(len(*p))
}
