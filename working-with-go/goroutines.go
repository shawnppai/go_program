package main

import (
	"fmt"
	"time"
)

func printSlowly(s string, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i, s)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	printSlowly("directly functioning", 3)
	go printSlowly("red fish goroutines", 3)
	go printSlowly("bule fish goroutines", 3)

	go func(ss string, nn int) {
		for i := 0; i < nn; i++ {
			fmt.Println(i, ss)
			time.Sleep(150 * time.Millisecond)
		}
	}("annoy fish goroutines", 3)

	var input int
	fmt.Print("Input Num: ")
	fmt.Scanf("%d", &input)
	fmt.Printf("Input num is %d, Program Done!", input)
}
