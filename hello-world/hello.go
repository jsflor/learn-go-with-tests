package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"
	german  = "German"

	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
	germanPrefix  = "Hallo, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(language) + name
}

func getPrefix(language string) string {
	var prefix string

	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	case german:
		prefix = germanPrefix
	default:
		prefix = englishPrefix
	}

	return prefix
}

func main() {
	fmt.Println(Hello("", ""))
}
