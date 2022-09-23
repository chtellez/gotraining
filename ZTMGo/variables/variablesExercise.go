package main

import (
	"fmt"
)

func main() {
	var favoriteColor = "red"
	birthYear, age := 1986, 36
	var (
		firstInitial = "C"
		lastInitial  = "T"
	)
	var ageInDays int
	ageInDays = age * 365

	fmt.Println("Favorite color :", favoriteColor)
	fmt.Println("Birth year :", birthYear)
	fmt.Println("Age :", age)
	fmt.Println("Age in days:", ageInDays)
	fmt.Println("First and last initials :", firstInitial, lastInitial)

}
