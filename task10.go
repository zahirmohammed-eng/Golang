package main

import (
	"fmt"
)

// func getEnv(key, defaultValue string) string {
// 	value := os.Getenv(key)
// 	if value == "" {
// 		return defaultValue
// 	}
// 	return value
// }

func main() {

	appName := "EmployeeManagementApp"
	appPort := "8080"

	fmt.Printf("Application Name: %s\n", appName)
	fmt.Printf("Running on Port:  %s\n", appPort)
}
