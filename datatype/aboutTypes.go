package main

import "fmt"

/* Whether or not a struct is comparable depends on the struct’s fields.
Structs that are entirely composed of comparable types are comparable;*/
func aboutTypes() {

	type firstPerson struct {
		name string
		age  int
	}
	type secondPerson struct {
		secondName string
		age        int
	}

	type Manager struct {
		firstPerson //type embebido
		ID          string
	}

	//person1:=firstPerson{"Julia",2}
	person2 := secondPerson{"Julia", 2}
	person3 := secondPerson{"Julia", 2}

	fmt.Println(person3 == person2)
	//No se puede comparar diferentes types
	//fmt.Println(person1 == person2)

	//Enumerations
	type MailCategory int

	const (
		Uncategorized MailCategory = iota
		Personal
		Spam
		Social
		Advertisements
	)
	fmt.Printf("ENUMERATIONS [ %d ] --->\n", Spam)
	fmt.Printf("ENUMERATIONS [ %d ] --->\n", Advertisements)
}
