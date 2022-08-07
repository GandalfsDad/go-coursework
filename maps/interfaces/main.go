package main

import "fmt"

type languageBot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	sb := spanishBot{}
	eb := englishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(lb languageBot) {
	fmt.Println(lb.getGreeting())
}

func (englishBot) getGreeting() string {
	//Some voodoo for English greeting
	return "Hi"
}

func (spanishBot) getGreeting() string {
	//Some voodoo for spanish greeting
	return "Hola"
}
