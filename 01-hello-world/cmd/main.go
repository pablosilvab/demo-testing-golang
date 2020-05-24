package main

import "fmt"

const SPANISH = "Spanish"
const FRENCH = "French"
const ENGLISH_GREETING_PREFIX = "Hello, "
const SPANISH_GREETING_PREFIX = "Hola, "
const FRENCH_GREETING_PREFIX = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case SPANISH:
		return SPANISH_GREETING_PREFIX
	case FRENCH:
		return FRENCH_GREETING_PREFIX
	default:
		return ENGLISH_GREETING_PREFIX
	}
}

func main() {
	fmt.Println(Hello("Pablo", "Spanish"))
}
