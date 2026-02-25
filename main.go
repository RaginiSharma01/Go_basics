package main

import (
	"fmt"
	"log"

	methods "ddr/Methods"
	"ddr/maths"
)

func main() {
	sum := maths.Add(10, 5)
	diff := maths.Sub(10, 5)
	mod := maths.Mod(10, 2)

	fmt.Println("Addition:", sum)
	fmt.Println("Subtraction:", diff)
	fmt.Println("Modulo:", mod)

	emp1 := methods.Empolyee{
		Name:     "Sam Adolf",
		Salary:   500000,
		Currency: "$",
	}

	log.Fatal(emp1)
}
