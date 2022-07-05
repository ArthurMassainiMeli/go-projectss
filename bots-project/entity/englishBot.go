package entity

type EnglishBot struct {
	chat string
	name string
}

func (eb EnglishBot) GetGreeting() string {
	// custom logic
	return "Welcome to the world"
}
