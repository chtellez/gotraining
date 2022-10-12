package main

import (
	"dataTypes/organization"
	"fmt"
)

func main() {

	eur := organization.NewEuropeanUnionIdentifier("987.65-4321", "Germany")

	p := organization.NewPerson("Carlos", "Gonzales", eur)
	err := p.SetTwitterHandler("@cgonza")
	fmt.Printf("%T\n", organization.TwitterHandler("Prueba"))
	fmt.Printf("%T\n", string("Prueba"))
	if err != nil {
		fmt.Printf("An error ocurred setting twitter handler: %s\n", err.Error())
	}

	println(p.ID())
	println(p.Country())
	println(p.FullName())
	println(p.TwitterHandler())
	println(p.TwitterHandler().RedirectUrl())
}
