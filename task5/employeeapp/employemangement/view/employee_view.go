package view

import (
	"employeeapp/employemangement/service"
	"fmt"
)

func Seemessage() {
	m := service.GetMessage()
	fmt.Println(m)

}
