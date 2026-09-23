package main


import "fmt"

func main() {

	ids := [3]int{101, 102, 103}
	names := [3]string{"Alice Smith", "Bob Jones", "Charlie Brown"}
	positions := [3]string{"Backend Engineer", "Product Manager", "DevOps Specialist"}
	depts := [3]string{"Engineering", "Product", "Infrastructure"}

	var searchID int

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&searchID)

	found := false

	for i := 0; i < 3; i++ {

		if ids[i] == searchID {

			fmt.Println("=== Employee Found ===")
			fmt.Println("ID:", ids[i])
			fmt.Println("Name:", names[i])
			fmt.Println("Position:", positions[i])
			fmt.Println("Dept:", depts[i])

			found = true
			break
		}
	}

	if found == false {
		fmt.Println("No employee found")
	}
}