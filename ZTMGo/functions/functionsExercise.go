package main

import (
	"fmt"
	"math/rand"
)

//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func saludarA(name string) {
	fmt.Println("Hola, mucho gusto en conocerte", name)
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func saludo() string {
	return "Este es un saludo standard desde la aplicación!"
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func sumar3(pro, sdo, tro int) int {
	return pro + sdo + tro
}

// * Write a function that returns any number
func numero() int {
	return rand.Int()
}

// * Write a function that returns any two numbers
func numeros() (int, int) {
	return rand.Int(), rand.Int()
}

func main() {
	//* Add three numbers together using any combination of the existing functions.
	a, b := numeros()
	c := sumar3(numero(), a, b)
	//  * Print the result
	fmt.Println(c)
	//* Call every function at least once
	saludarA("Christian")
	fmt.Println(saludo())

}
