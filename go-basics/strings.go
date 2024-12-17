package main

import (
	"fmt"
)

func main() {
	phrase := "We are learning Go!"
	fmt.Println("Original string: ", phrase)
	fmt.Println("length: ", len(phrase))

	fmt.Println("First character: ", phrase[0])

	fmt.Println("first 2: ", phrase[2:])
}
