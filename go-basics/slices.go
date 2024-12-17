package main

import "fmt"

func main() {
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)

	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}
	loons = append(loons, "tweety")
	fmt.Println(loons)
}
