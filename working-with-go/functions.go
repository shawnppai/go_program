package main

import (
   "fmt"
   "math"
)

func Echo(s string) {
	fmt.Println(s)
}

func Say(s string) string {
	phrase := "Hello " + s
	return phrase
}

func Say2(s string) (phrase string) {
	phrase = "Hello " + s
	return
}

func Divide(x, y float64) (float64, float64) {
	q := math.Trunc(x / y)
	r := math.Mod(x, y)
	return q, r
}

func Divide2(x, y float64) (q, r float64) {
	q = math.Trunc(x / y)
	r = math.Mod(x, y)
	return
}

func main() {
	Echo("Bonjour tout le monde")
	fmt.Println(Say("Gopher"))

	q, r := Divide(11, 3)
	fmt.Printf("Quotient: %v, Remainder %v \n", q, r)
}