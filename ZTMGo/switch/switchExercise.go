// --Summary:
//
//	Create a program to display a classification based on age.
package main

import "fmt"

func main() {

	//--Requirements:
	//* Use a `switch` statement to print the following:
	//  - "newborn" when age is 0
	switch age := 17; {
	case age == 0:
		fmt.Println("Newborn")
	//  - "toddler" when age is 1, 2, or 3
	case age <= 3:
		fmt.Println("Toddler")
	//  - "child" when age is 4 through 12
	case age <= 12:
		fmt.Println("Child")
	//  - "teenager" when age is 13 through 17
	case age <= 17:
		fmt.Println("teenager")
	//  - "adult" when age is 18+
	default:
		fmt.Println("Adult")
	}

}
