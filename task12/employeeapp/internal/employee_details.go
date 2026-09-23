package internal

import (
	"employeeapp/pkg"
	"fmt"
)

// purpose: The internal/ directory holds private application and business logic that cannot be imported by external packages/projects.

func Employeedata() {
	pkg.About()

	name := "vimal"
	age := 23
	salary := 50000.00

	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("salary:", salary)
	fmt.Println("hike:", pkg.Hike(salary))

}
