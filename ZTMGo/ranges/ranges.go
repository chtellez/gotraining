package main

import (
	"fmt"
)

func main() {
	slice := []string{"Hello", "World", "!"}

	for i, element := range slice {
		fmt.Println(i, element, ":")
		for _, char := range element {
			fmt.Printf("\t%q\n", char)
		}
	}

}
