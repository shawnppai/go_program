package main

import "fmt"
import "math"

func main() {
	sum := 0.0
	for sum < 1000 {
		fmt.Println(math.Sqrt(sum))
		sum += 1
	}
}
