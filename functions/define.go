package main

func main() {
	sum := add(1, 2)
	add, subt := addAndSubtract(1, 2)
	println(sum)
	println(add)
	println(subt)
}

// adds two numbers
func add(one int, two int) int {
	return one + two
}

func addAndSubtract(one int, two int) (int, int) {
	return one + two, one - two
}
