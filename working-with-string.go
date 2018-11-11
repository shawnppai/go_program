package main

import (
    "fmt"
    "strings"
)


func main() {
    s := "Hi, I am Shawnpai, a new go programe"
    lower := strings.ToLower(s)
    fmt.Println(lower)
    if strings.Contains(lower, "shawnpai") {
        fmt.Println("Yes, it contains 'shawnpai'")
    }
    fmt.Println("The Character 5-16 is : " + s[4:17])
    sentence := "I am a sentence made up of words"
    words := strings.Split(sentence, " ")
    for i := 0; i < len(words); i ++ {
        fmt.Println(words[i])
    }
    fields := strings.Fields(sentence)
    fmt.Printf("%v \n", fields)
}
