package main

import "fmt"

//swap using ptr
//go is always an pass by value

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	//x := 10

	// var p *int
	// p = &x

	// fmt.Println("x value:", x)
	// fmt.Println("x address:", &x)
	// fmt.Println("pointer p:", p)
	// fmt.Println("value using pointer:", *p)

	// //changing the value using pte

	// *p = 20
	// fmt.Println(x)

	x := 40
	y := 30

	swap(&x, &y)
	fmt.Println(x, y)
}
