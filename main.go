package main

import (
	"ddr/maths"
	"fmt"
)

func main() {
	sum := maths.Add(10, 5)
	diff := maths.Sub(10, 5)
	mod := maths.Mod(10, 2)

	fmt.Println("Addition:", sum)
	fmt.Println("Subtraction:", diff)
	fmt.Println("Modulo:", mod)

}
