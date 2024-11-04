package main

import "fmt"

const (
	prefixEnglish = "Hello, "
	prefixSpanish = "Hola, "
	prefixFrance  = "Bonjour, "
	prefixBrasil  = "Olá, "
	spanish       = "Spanish"
	france        = "France"
	brasil        = "Brasil"
)

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = prefixSpanish
	case france:
		prefix = prefixFrance
	case brasil:
		prefix = prefixBrasil
	default:
		prefix = prefixEnglish
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
