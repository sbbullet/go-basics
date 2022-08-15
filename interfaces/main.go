package main

import "fmt"

type Bot interface {
	getGreeting() string
}

type EnglishBot struct{}
type SpanishBot struct{}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}

	// eb.printGreeting()
	// sb.printGreeting()

	printGreeting(eb)
	printGreeting(sb)
}

func (EnglishBot) getGreeting() string {
	return "Hi, there!"
}

func (SpanishBot) getGreeting() string {
	return "Hola"
}

func printGreeting(b Bot) {
	fmt.Println(b.getGreeting())
}

// func (eb EnglishBot) printGreeting() {
// 	fmt.Println(eb.getGreeting())
// }

// func (sb SpanishBot) printGreeting() {
// 	fmt.Println(sb.getGreeting())
// }
