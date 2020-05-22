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
	switch language {
	case SPANISH:
		return SPANISH_GREETING_PREFIX + name
	case FRENCH:
		return FRENCH_GREETING_PREFIX + name
	default:
		return ENGLISH_GREETING_PREFIX + name
	}
}

func main() {
	fmt.Println(Hello("Pablo", "Spanish"))
}
