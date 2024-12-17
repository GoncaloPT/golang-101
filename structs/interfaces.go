package main

import (
	"fmt"
	"io"
	"os"
)

type Capper struct {
	writer io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		if p[i] >= 'a' && p[i] <= 'z' {
			p[i] -= 'a' - 'A'
		}
	}
	return c.writer.Write(p)
}

func main() {
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello here!")
}
