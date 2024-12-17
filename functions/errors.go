package main

import (
	"fmt"
	"math"
)

func main() {

	s1, err := sqrt(2.0)
	if err != nil {
		fmt.Println("ERROR IS %s§n", err)
	} else {
		fmt.Println(s1)
	}

}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0.0, fmt.Errorf("cannot sqrt negative number (%f)", n)
	}
	return math.Sqrt(n), nil

}
