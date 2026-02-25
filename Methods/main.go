package methods

import (
	"fmt"
)

type Empolyee struct {
	Name     string
	Salary   int
	Currency string
}

func (e Empolyee) displaySalary() {
	fmt.Printf("salary of %s is %s%d", e.Name, e.Currency, e.Salary)
}
