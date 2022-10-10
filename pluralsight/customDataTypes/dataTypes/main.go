package main

import (
	"dataTypes/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson("Christian", "Téllez")
	err := p.SetTwitterHandler("@ctellez")
	if err != nil {
		fmt.Printf("An error ocurred setting twitter handler: %s\n", err.Error())
	}
	println(p.ID())
	println(p.FullName())
	println(p.TwitterHandler())
}
