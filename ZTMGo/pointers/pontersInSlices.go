package main

import "fmt"

type Product struct {
	brand string
}

func changeBrandArray(productos [2]Product) {
	productos[0].brand = "Nueva Marca 1"
	productos[1].brand = "Nueva Marca 2"
	fmt.Println("\nSe cambio la marca Array\n", productos)
}

func changeBrandSlice(productos []Product) {
	productos[0].brand = "Nueva Marca 1"
	productos[1].brand = "Nueva Marca 2"
	fmt.Println("\nSe cambio la marca Slice\n", productos)
}

func main() {
	bmw := Product{"BMW"}
	mazda := Product{"Mazda"}

	carrosArray := [...]Product{bmw, mazda}
	carrosSlice := []Product{bmw, mazda}
	fmt.Println(carrosArray)

	changeBrandArray(carrosArray)
	changeBrandSlice(carrosSlice)

	fmt.Println("\nImprimiendo marca de carros luego de la función Array", carrosArray)
	fmt.Println("\nImprimiendo marca de carros luego de la función Slice", carrosSlice)

}
