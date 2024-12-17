package main

import (
	"fmt"
	"log"
)

type Square struct {
	x      int
	y      int
	length int
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if x < 0 || y < 0 || length < 0 {
		return nil, fmt.Errorf("x, y, and length must be greater than 0")
	}
	s := Square{
		x:      x,
		y:      y,
		length: length,
	}
	return &s, nil
}

func (s *Square) Move(dx int, dy int) {
	s.x += dx
	s.y += dy

}
func (s *Square) Area() int {
	return s.length * s.length
}

func main() {
	s, err := NewSquare(1, 2, 10)
	if err != nil {
		log.Fatalf("ERROR: %s\n", err)
	}

	s.Move(2, 2)
	fmt.Printf("%+v\n", s)

}
