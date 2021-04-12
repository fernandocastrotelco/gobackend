package main

import (
	"fmt"

	"github.com/fernandocastrotelco/gobackend/datatypes/organization"
)

func main() {
	p := organization.NewPerson("James", "Wilson", organization.NewEuropeanUnionIdentifier("my real id", "Germany"))
	err := p.SetTwitterHandler("@jam_wils")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err.Error())
	}

	println(p.ID())
	println(p.Country())
}