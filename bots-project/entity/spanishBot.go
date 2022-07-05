package entity

type PortugueseBot struct {
	chat string
	name string
}

func (pb PortugueseBot) GetGreeting() string {
	// custom logic
	return "Bem vindo ao mundo"
}
