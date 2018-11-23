package main

import "fmt"

type Person struct {
	Name string
	age  int
}

//实现String()方法的重写，重新定义了fmt.Println()行为
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.age)
}

func main() {
	a := Person{"Sunny", 26}
	b := Person{"Jack", 22}
	fmt.Println(a, b)
}
