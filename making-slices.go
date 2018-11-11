package main

import "fmt"

func main() {
    a := make([]int, 5)
    printSlice("a", a)
    b := make([]int, 0, 5)
    printSlice("b", b)
    c := b[:2]
    printSlice("c", c)
    //没有指定后面的index，使用的是slice
    d := c[2:]
    printSlice("d", d)
    //指定了后面的index的时候，使用的重新分配的slice
    e := c[2:5]
    printSlice("e", e)
}

func printSlice(s string, sl []int) {
    fmt.Printf("%s's len = %d, cap = %d, %v\n", s, len(sl), cap(sl), sl)
}
