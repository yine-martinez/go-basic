package main

import "fmt"

// Las interfaces se heredan implicitamente
//A concrete type does not declare that it implements an interface.

type LogicProvider struct{}

/* Method declarations look just like function declarations,
with one addition: the receiver specification (They can be pointer
receivers (the type is a pointer) or value receivers (the type is a value type).).*/
func (lp LogicProvider) Process(data string) string {
	// business logic
	fmt.Println("Procesando data en el método")
	return data
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	// get data from somewhere
	c.L.Process("data")
}

func aboutInterfaces() {
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}
