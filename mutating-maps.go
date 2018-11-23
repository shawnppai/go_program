package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The values: ", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The values: ", m["Answer"])

	m["Question"] = 123
	delete(m, "Answer")
	fmt.Println("The values: ", m["Answer"])
	fmt.Println(m)

	v, ok := m["Answer"]
	fmt.Println("The values: ", v, "Present?: ", ok)
}
