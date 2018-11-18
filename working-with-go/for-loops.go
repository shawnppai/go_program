package main

import "fmt"

func main() {
	alphas := []string{"abc", "def", "ghi"}
	for i := 0; i < len(alphas); i++ {
		fmt.Printf("%d: %s \n", i, alphas[i])
	}

	for i, val := range alphas {
		fmt.Printf("%d: %s \n", i, val)
	}

    //不使用val
	for i := range alphas {
		fmt.Println(i)
	}

	for _, val := range alphas {
		fmt.Println(val)
	}

	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}
}