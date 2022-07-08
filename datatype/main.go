package main

// Learning about Types, Methods and Interfaces
import (
	"fmt"
	"go-basic/datatype/organization"
)

func main() {
	aboutInterfaces()
	aboutTypes()
	p := organization.NewPerson("Yine", "Martinez", organization.NewEuropeanUnionIdentifier("123-45-6788", "Venezuela"))
	err := p.SetTwitterHandler("@neithblue")
	fmt.Printf("%T\n", organization.TwitterHandler("Test"))
	if err != nil {
		fmt.Printf("An error ocurred setting twitter handler: %s", err.Error())
	}
	/*println(p.FullName())
	println(p.TwitterHandler().RedirectUrl())
	println(p.TwitterHandler())
	println(p.ID())
	println(p.Country())*/

	/*name1 := Name{First: "Yine", Last: "Martinez"}
	name2 := Name{First: "Yine", Last: "Martinez"}

	if name1 == name2 {
		println("We match")
	}

	//Interfaces
	ssn := organization.NewSocialSecurityNumber("123-456-789")
	eu := organization.NewEuropeanUnionIdentifier("123-456-789", "Venezuela")

	if ssn == eu {
		println("Interface - We match")
	}

	//Key maps
	portfolio := map[Name][]organization.Person{}
	portfolio[name1] = []organization.Person{p}*/
}

type Name struct {
	First string
	Last  string
}
