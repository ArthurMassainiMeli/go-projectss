package main

import (
	"fmt"
	"geometry_project/entity"
	"geometry_project/interfaces"
)

func main() {
	square := entity.Square{
		Side: 5.2,
	}

	triangle := entity.Triangle{
		Height: 3.5,
		Base:   2.8,
	}

	fmt.Println("Square area is: ")
	interfaces.PrintArea(square)
	fmt.Println("Triangle area is: ")
	interfaces.PrintArea(triangle)
}
