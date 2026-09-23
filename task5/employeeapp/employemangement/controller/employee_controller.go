package controller

import "employeeapp/employemangement/dao"

func Getmessage() string{
	return dao.GetMessage()
}