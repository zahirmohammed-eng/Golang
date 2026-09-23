package employee

import (
	"fmt"
	"employeeapp/utils"
)

func Employeedata(){
	utils.About()

	name := "vimal"
	age:=23
	salary:=50000.00

	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("salary:", salary)
	fmt.Println("hike:", utils.Hike(salary))


}