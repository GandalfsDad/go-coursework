package main

import "fmt"

type englishBot struct{}

//type spanishBot struct{}

func main() {
	//sb := spanishBot{}
	eb := englishBot{}

	printGreeting(eb)
	//printGreeting(sb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (englishBot) getGreeting() string {
	//Some voodoo for English greeting
	return "Hi"
}

func (spanishBot) getGreeting() string {
	//Some voodoo for spanish greeting
	return "Hola"
}
