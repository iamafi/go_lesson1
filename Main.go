package main

import "fmt"

var id int

func main() {
	choices := "Choose the number to continue: \n" +
		"1. Create new contact\n" +
		"2. Update contact\n" +
		"3. Delete contact\n" +
		"4. Show particular contact\n" +
		"5. Show all contacts\n" +
		"0. Exit the program\n\n"
	var choice int = 1
	fmt.Print("Welcome to Contact List!\n\n")
	for choice != 0 {
		fmt.Print(choices)
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			create()
		case 2:
			update(getID("update"))
		case 3:
			delete(getID("delete"))
		case 4:
			getOne(getID("see"))
		case 5:
			getAll()
		default:
			break
		}
	}
}
