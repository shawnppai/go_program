package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    sl := make([][]uint8, dy, dy)
    for index, _ := range sl {
    	sl[index] = make([]uint8, dx, dx)
    }
    return sl

}

func main() {
	pic.Show(Pic)
}