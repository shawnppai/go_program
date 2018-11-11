package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    fmt.Println(v)
    v.X = 21
    v.Y = 1112
    fmt.Println(v)
}
