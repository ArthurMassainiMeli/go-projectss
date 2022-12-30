package interfaces

import "fmt"

type shape interface {
	GetArea() float64
}

func PrintArea(s shape) {
	area := fmt.Sprintf("%.2f", s.GetArea())
	fmt.Println(area)
}
