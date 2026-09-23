package main


import "fmt"

func main() {

	var ids [100]int
	var names [100]string
	var positions [100]string
	var depts [100]string

	count := 0

	for {

		fmt.Println("\n===== Employee Management =====")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee")
		fmt.Println("3. Display Employees")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			// Add Employee

			fmt.Print("Enter Employee ID: ")
			fmt.Scan(&ids[count])

			fmt.Print("Enter Employee Name: ")
			fmt.Scan(&names[count])

			fmt.Print("Enter Position: ")
			fmt.Scan(&positions[count])

			fmt.Print("Enter Department: ")
			fmt.Scan(&depts[count])

			count++

			fmt.Println("Employee added successfully!")

		case 2:
			// Search Employee

			var searchID int
			found := false

			fmt.Print("Enter Employee ID: ")
			fmt.Scan(&searchID)

			for i := 0; i < count; i++ {

				if ids[i] == searchID {

					fmt.Println("\n=== Employee Found ===")
					fmt.Println("ID:", ids[i])
					fmt.Println("Name:", names[i])
					fmt.Println("Position:", positions[i])
					fmt.Println("Department:", depts[i])

					found = true
					break
				}
			}

			if !found {
				fmt.Println("Employee not found.")
			}

		case 3:
			// Display Employees

			if count == 0 {
				fmt.Println("No employees available.")
			} else {

				fmt.Println("\n=== Employee List ===")

				for i := 0; i < count; i++ {

					fmt.Println("--------------------")
					fmt.Println("ID:", ids[i])
					fmt.Println("Name:", names[i])
					fmt.Println("Position:", positions[i])
					fmt.Println("Department:", depts[i])
				}
			}

		case 4:
			// Delete Employee

			var deleteID int
			found := false

			fmt.Print("Enter Employee ID to delete: ")
			fmt.Scan(&deleteID)

			for i := 0; i < count; i++ {

				if ids[i] == deleteID {

					for j := i; j < count-1; j++ {
						ids[j] = ids[j+1]
						names[j] = names[j+1]
						positions[j] = positions[j+1]
						depts[j] = depts[j+1]
					}

					count--
					found = true

					fmt.Println("Employee deleted successfully!")
					break
				}
			}

			if !found {
				fmt.Println("Employee not found.")
			}

		case 5:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}