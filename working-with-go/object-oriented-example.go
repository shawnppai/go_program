package main

import (
    "fmt"
)

type Dog struct {
	Name string
	Color string
}

func (d Dog) Call() {
	fmt.Printf("Here comes a %s dog, %s.\n", d.Name, d.Color)
}

func main() {
	Spot := Dog{Name: "Spot", Color: "brown"}
	var Rover Dog
	Rover.Name = "Rover"
	Rover.Color = "light blondish with some dark patches and very scruffy"

	Spot.Call()
	Rover.Call()
}