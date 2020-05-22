package main

import "fmt"

const ENGLISH_GREETING_PREFIX = "Hello, "

func Hello(name string) string {
	return ENGLISH_GREETING_PREFIX + name
}

func main() {
	fmt.Println(Hello("Pablo"))
}
