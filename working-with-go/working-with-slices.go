package main

import "fmt"

func main() {
    var empty []int
    elphas := []string{"abc", "def", "ghi", "jkl"}
    empty = append(empty, 123)
    empty = append(empty, 456)
    fmt.Printf("%v \n", empty)

    elphas = append(elphas, "pqr", "stu")
    fmt.Printf("%v \n", elphas)
    fmt.Println("Length: ", len(elphas))
    fmt.Println(elphas[1])
    if elemExists("def", elphas) {
        fmt.Println("Exists!")
    }
}

func elemExists(s string, a []string) bool {
    for _, v := range a {
        if v == s {
            return true
        }
    }
    return false
}
