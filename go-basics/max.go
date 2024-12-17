package main

/*
Main function, the challenge is to print the maximum value from a list of numbers
*/
func main() {

	numbers := []int{4, 2, 44, 446, 212, 445}

	// declare a variable called foundMax with no initial value
	var foundMax int
	for _, number := range numbers {
		if number > foundMax {
			foundMax = number
		}

	}
	println("Max number is: ", foundMax)
}
