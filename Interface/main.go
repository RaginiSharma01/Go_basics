package main

import "fmt"

type Speaker interface {
	Speak()
}

type User struct {
	Name string
}

type Dog struct {
	Breed string
}

func (u User) Speak() {
	fmt.Println("User speaks:", u.Name)
}

func (d Dog) Speak() {
	fmt.Println("Dog barks:", d.Breed)
}

func main() {

	var s Speaker

	s = User{"Ragini"}
	s.Speak()

	s = Dog{"Labrador"}
	s.Speak()
}
