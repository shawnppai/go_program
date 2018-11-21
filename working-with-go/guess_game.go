package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var guess int
	var count int

	num := rand.Intn(100) + 1
	fmt.Println(">> I'm thinking of a number between 1-100 <<")
	for {
		fmt.Print("Guess：\n")
		_, err := fmt.Scanf("%d", &guess)
		if err == nil {
			count += 1
			if guess > num {
				fmt.Println(">> Too High")
			} else if guess < num {
				fmt.Println(">> Too Low")
			} else {
				fmt.Printf("Correct ! You Guess The %d Times Guess!\n", count)
				break
			}

		} else {
			fmt.Println(">> Please Input a number")
		}
	}
}
