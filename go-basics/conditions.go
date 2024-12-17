package main

import (
	"fmt"
)

func main() {
	x := 10
	if x > 5 {
		fmt.Println("x is big")
	} else {
		fmt.Println("x is small")
	}
	
	if frac := x / 10; frac >= 1 {
		fmt.Println("x has at least one whole number")
	} else {
		fmt.Println("x is a fraction")
	}

	switch x {
	case 1:
		fmt.Println("x is 1")
	default:
		fmt.Println("x is not 1")
	}
}
