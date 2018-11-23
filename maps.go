package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	//m["Bell Labs"] = Vertex{40.68433, -74.39967}
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967, //这样实例化结构体需要在末尾在逗号
	}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
}
