package main

import (
	"bots-project/entity"
	"bots-project/interfaces"
)

func main() {
	pb := entity.PortugueseBot{}
	eb := entity.EnglishBot{}

	interfaces.PrintGreeting(pb)
	interfaces.PrintGreeting(eb)
}
