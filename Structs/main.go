package main

import (
	"fmt"
)

//embedding

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	id     int
	Salary int
}

func main() {
	e := Employee{
		Person: Person{
			Name: "ragini",
			Age:  21,
		},

		id:     23,
		Salary: 60000,
	}

	fmt.Println(e.Name)
	fmt.Println(e.id)
	fmt.Println(e.Age)
	fmt.Println(e.Salary)

}

// type User struct {
// 	Name  string
// 	Age   int
// 	grade int
// }

// // this func return the struct
// func test() (User, User) {
// 	u1 := User{"ragini", 21, 87}
// 	u2 := User{"asha", 66, 55}
// 	return u1, u2
// }

// // methods
// func (u User) greet() {
// 	fmt.Println("hello", u.Name)
// }

// // pointer -- used to modify the memory not just values
// func updateAge(u *User) {
// 	u.Age = 30
// }

// // slices with structs
// // global
// var users = []User{
// 	{"ragini", 21, 87},
// 	{"aman", 22, 90},
// }

// // json

// func main() {
// 	u1, u2 := test()
// 	fmt.Println(u1.Name)
// 	fmt.Println(u1.Age)
// 	u1.Age = 23
// 	fmt.Println(u1.Age)

// 	u1.greet()

// 	updateAge(&u1)
// 	fmt.Println(u1.Age)

// 	fmt.Println(u1 == u2)

// 	fmt.Println(users)
// 	fmt.Println("first user name :", users[0].Name)

// 	data, err := json.Marshal(users) //"_" --> ignore the erros
// 	if err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}
// 	fmt.Println(string(data))

// }
