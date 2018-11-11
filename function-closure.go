package main

import "fmt"

/*func adder() func(a int) int {
    sum := 0
    return func(a int) int {
    	sum += a
    	return sum
    }
}
*/
func adder() func(a int) int {
    sum := 0
    fn := func(a int) int {
    	sum += a
    	return sum
    }
    return fn
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}