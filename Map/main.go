package main

import "fmt"

type Student struct {
	Name string
	Id   int
	Mark int
}

func main() {
	m := make(map[string]Student)

	m["student1"] = Student{
		Name: "ragini",
		Id:   23,
		Mark: 87,
	}
	fmt.Println(m)
	fmt.Println(m["student1"])
}
