package interfaces

import "fmt"

type bot interface {
	GetGreeting() string
}

func PrintGreeting(b bot) {
	fmt.Println(b.GetGreeting())
}
