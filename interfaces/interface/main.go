package main

type genericBot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func pringGreeting(gb genericBot) {
	println(gb.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	pringGreeting(eb)
	pringGreeting(sb)
}
