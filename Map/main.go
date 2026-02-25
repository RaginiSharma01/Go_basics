package main

import "fmt"

type Student struct {
	Name string
	Id   int
	Mark int
}

func main() {
	m := make(map[string]Student)

	m["s1"] = Student{"ragini", 1, 90}
	m["s2"] = Student{"anmol", 2, 88}
	s := m["s1"]

	fmt.Println(s)

	s.Mark = 99
	m["s1"] = s
	fmt.Println(s)

	delete(m, "s2")
	fmt.Println(m)

	v, ok := m["s1"]
	fmt.Println(v, "present?", ok)

}
