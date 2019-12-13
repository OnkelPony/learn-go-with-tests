package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const ivrit = "Hebrew"
const spanishHelloPrefix = "Hola, "
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const ivritHelloPrefix = "Shalom, "

// Hello returns a greeting
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case ivrit:
		prefix = ivritHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("world", ""))
}
