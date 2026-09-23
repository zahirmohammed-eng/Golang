// 3.cli calculater

package main

import (
	"fmt"
	"os"
)

func main() {
	var n1, n2 float32
	var op string
	fmt.Println("enter n1: ")

	_, err := fmt.Scanln(&n1)
	if err != nil {
		fmt.Println("invalid character")
		return
	}
	fmt.Println("enter n2: ")

	_, err1 := fmt.Scanln(&n2)
	if err1 != nil {
		fmt.Println("invalid character")
		return
	}
	fmt.Println("enter operator: ")
	fmt.Scanln(&op)
	switch op {
	case "+":
		fmt.Println(n1 + n2)
	case "-":
		fmt.Println(n1 - n2)
	case "*":
		fmt.Println(n1 * n2)
	case "/":
		if n2 == 0 {
			fmt.Println("error")
			os.Exit(1)
		}
		fmt.Println(n1 / n2)
	default:
		fmt.Println("enter +,-,*,/ only")
		os.Exit(1)
	}

}
