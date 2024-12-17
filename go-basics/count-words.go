package main

import "strings"

/*
- Create a program that counts how many times a word appears in a text.
*/
func main() {
	phrase := "We are learning Go! Go is a great language to learn. Go is fun!"
	split := strings.Fields(strings.ToLower(phrase))
	println("split	", split)
}
