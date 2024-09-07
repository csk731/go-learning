package ch7

import (
	"fmt"
	"time"
)

type PersonData struct {
	name      string
	age       int
	phone     string
	createdAt time.Time
}

type Admin struct {
	email      string
	password   string
	PersonData // Embedding PersonData in this struct pulls all the variables and methods of PersonData struct.
	// PersonData PersonData // Accessing members of PersonData from Admin obj requires us to use PersonData in between if the declaraton is like this. I flike above one - noe intermediate is required.!
}

// Function that is attached to the struct is method
func (pd PersonData) PrintPersonDataOfStruct() {
	fmt.Println("Address-Struct:", &pd)
	fmt.Println("PrintPersonDataOfStruct:", pd.name, pd.age, pd.phone)
}

func (pd *PersonData) clearUserName() {
	fmt.Println("Address of Name Inside:", &pd.name)
	pd.name = ""
}

func Structs() {
	personData := new(PersonData)
	personData.age = 15
	personData.name = "Chay"
	personData.phone = "1234567890"
	personData.createdAt = time.Now()
	fmt.Println(personData)

	var personData1, pd2 PersonData
	fmt.Println(personData1)
	personData1 = PersonData{
		name:      "C",
		age:       12,
		phone:     "213",
		createdAt: time.Now(),
	}
	pd2 = PersonData{}
	fmt.Println(personData1, pd2)

	fmt.Println("Address:", &personData1)
	PrintPersonData(&personData1)
	personData1.PrintPersonDataOfStruct()

	fmt.Println("Address of Name Outside:", &personData1.name)
	personData1.clearUserName()
	fmt.Println(personData1)

	admin := &Admin{
		email:    "*@gmail.com",
		password: "2dCnm@1",
		PersonData: PersonData{
			name:      "C",
			age:       12,
			phone:     "22",
			createdAt: time.Now(),
		},
	}
	admin.PrintPersonDataOfStruct()
	fmt.Println(admin)
}

func PrintPersonData(personData1 *PersonData) {
	fmt.Println("PrintPersonData:", personData1.name, personData1.age) // Dereferencing for pointers in not required for structs to access members - An exception
}