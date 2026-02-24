package main

import "fmt"

func main() {
	//Arrays
	num := [...]int{1, 2, 3, 4, 5}
	fmt.Println(num)
	car := [...]string{"volvo", "BMW", "ford"}
	fmt.Println(car)
	//slices
	//len() --> length of the slice
	//cap() --> capacity of the slice
	//syntax --> slice_name := [] dataype{values}

	albums := []string{"taylor", "sabrina", "blackpink", "jennei"}
	fmt.Println(len(albums))
	fmt.Println(cap(albums))

}
