package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Person
	Email string
}

func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

func PrintPerson(p Person) {
	fmt.Printf("%s (%d)\n", p.FullName(), p.Age)
}

func main() {
	john := Person{FirstName: "John", LastName: "Doe", Age: 30}
	PrintPerson(john)

	jane := struct {
		FirstName string
		LastName  string
		Age       int
	}{FirstName: "Jane", LastName: "Doe", Age: 25}
	var person = Contact{
		Person: Person{
			FirstName: "Ming",
			LastName:  "Jia",
			Age:       0,
		},
		Email: "abc@msn.com",
	}
	fmt.Printf("fullname: %v\n", person.FullName())
	PrintPerson(jane)
}
