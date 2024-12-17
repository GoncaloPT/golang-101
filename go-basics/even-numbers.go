package main

import "fmt"

/*
2nd challenge:
find even-end numbers.
An even-end number is a number that ends with an even number.
*/
func main() {
	number := 1001
	numberString := fmt.Sprintf("%d", number)
	first := string(numberString[0])
	last := string(numberString[len(string(numberString))-1])
	println("First digit: ", first)
	println("Last digit: ", last)
	if first == last {
		println("Even-end number")
	} else {
		println("Not an even-end number")
	}

}
