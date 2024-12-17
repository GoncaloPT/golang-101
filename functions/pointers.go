package main

import "fmt"

func main() {
	value := 2
	double(value)
	fmt.Print(value)
	doublePtr(&value)
	fmt.Print(value)
}

// this does nothing since the value and not the reference is passed
func double(value int) {
	value *= 2
}

func doublePtr(value *int) {
	*value *= 2
}
