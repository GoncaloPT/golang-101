package main

import "fmt"

/*
* FizzBuzz is a simple number game where you count,
* but say "Fizz" and/or "Buzz" instead of numbers adhering
* to certain rules.
 */
func main() {
	for i := 1; i <= 20; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz")
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
		}
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
		}
	}
}
