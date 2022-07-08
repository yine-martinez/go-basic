package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Identifiable interface {
	ID() string
}

type Citizen interface {
	Identifiable
	Country() string
}

type socialSecurityNumber string

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "United States of America"
}
func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

type europeanUnionIdentifier struct {
	id      string
	country string
}

func (eui europeanUnionIdentifier) ID() string {
	return eui.id
}
func (eui europeanUnionIdentifier) Country() string {
	return eui.country
}
func NewEuropeanUnionIdentifier(id string, country string) Citizen {
	return europeanUnionIdentifier{
		id:      id,
		country: country,
	}
}

//Alias ¿Cual es la diferencia entre una alias una definición de tipos?
//type TwitterHandler = string
//Type declaration
type TwitterHandler string

func (th TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

type Name struct {
	first string
	last  string
}

/* Method declarations look just like function declarations,
with one addition: the receiver specification (They can be pointer
receivers (the type is a pointer) or value receivers (the type is a value type).).*/
func (n *Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

type Employee struct {
	Name
}

type Person struct {
	firstName string
	lastName  string // campos no editables
	Name
	twitterHandler TwitterHandler
	Citizen
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}

func NewPerson(firstName string, lastName string, citizen Citizen) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
		Name:      Name{first: firstName, last: lastName},
		Citizen:   citizen}
}
func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func (p *Person) ID() string {
	return fmt.Sprintf("Person's identifier: %s", p.Citizen.ID())
}
func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handler must start with and @ symbol")
	}
	p.twitterHandler = handler
	return nil
}
