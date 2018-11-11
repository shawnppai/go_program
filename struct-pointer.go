package main

import "fmt"

type Vertex struct {
    X string
    Y int
}

func main() {
    var v Vertex
    v = Vertex{"hello", 123}
    fmt.Println(v)
    v.X = "World"
    p := &v
    v.Y = 123456789
    fmt.Println(v)
    fmt.Println(p)
    fmt.Println(*p)
    fmt.Println(p.X, p.Y)
    fmt.Println((*p).X, (*p).Y)
}
