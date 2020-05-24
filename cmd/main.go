package main

import (
	"fmt"

	hello "github.com/pablosilvab/testing-golang/01-hello-world"
	integers "github.com/pablosilvab/testing-golang/02-integers"
)

func main() {
	fmt.Println(hello.Hello("Pablo", "Spanish"))
	fmt.Println(hello.Hello("Pablo", "French"))
	fmt.Printf("Result Add: %d", integers.Add(4, 5))
}
