package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	doublePre := 0
	pre := 1
	return func() int {
		doublePre, pre = pre, doublePre+pre
		return pre - doublePre
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Println(f())
	}
}
