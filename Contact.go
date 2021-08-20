package main

import "fmt"

var db []Contact

type Contact struct {
	id int
	firstName, lastName, phone, email, position string
}

func contactInput(contact Contact) Contact {
	var str string
	fmt.Print("Enter first name: ")
	fmt.Scanln(&str)
	contact.firstName = str
	fmt.Print("Enter last name: ")
	fmt.Scanln(&str)
	contact.lastName = str
	fmt.Print("Enter phone: ")
	fmt.Scanln(&str)
	contact.phone = str
	fmt.Print("Enter email: ")
	fmt.Scanln(&str)
	contact.email = str
	fmt.Print("Enter position: ")
	fmt.Scanln(&str)
	contact.position = str
	return contact
}

func contactOutput(contact Contact) {
	fmt.Printf("ID: %d\n" +
		"First name: %s\n" +
		"Last name: %s\n" +
		"Phone: %s\n" +
		"Email: %s\n" +
		"Position: %s\n\n",
		contact.id, contact.firstName, contact.lastName, contact.email, contact.phone, contact.position)
}

func create() {
	l := len(db)
	if l == 0 {
		l = 1
	} else {
		l = db[len(db) - 1].id + 1
	}
	contact := contactInput(Contact{id: l})
	db = append(db, contact)
	fmt.Printf("\nContact %s %s was successfully created!\n\n", contact.firstName, contact.lastName)
}

func exists(id int) int {
	for i := range db {
		if db[i].id == id {
			return i
		}
	}
	fmt.Print("\nRequested contact doesn't exists!\n\n")
	return -1
}

func update(id int) {
	ex := exists(id)
	if ex != -1 {
		contact := contactInput(db[ex])
		db[ex] = contact
		fmt.Printf("\nContact %s %s was successfully updated!\n\n", contact.firstName, contact.lastName)
	}
}

func delete(id int) {
	ex := exists(id)
	if ex != -1 {
		db = append(db[:ex], db[ex+1:]...)
		fmt.Print("\nContact was successfully deleted!\n\n")
	}
}

func getOne(id int) {
	ex := exists(id)
	if ex != -1 {
		contactOutput(db[ex])
	}
}

func getAll() {
	if len(db) == 0 {
		fmt.Print("\nNo contacts yet\n\n")
	} else {
		for i := range db {
			contactOutput(db[i])
		}
	}
}

func getID(str string) int {
	fmt.Printf("Enter id of a contact you want to %s: ", str)
	fmt.Scanln(&id)
	return id
}