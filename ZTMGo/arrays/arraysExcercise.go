// --Summary:
//
//	Create a program that can store a shopping list and print
//	out information about the list.
package main

import "fmt"

// - Products must include the price and the name
type Product struct {
	name  string
	price int
}

func checkProductList(list [4]Product) {
	var sumaTotal, ultimoProducto int
	//* Print to the terminal:
	for i := 0; i < len(list); i++ {
		//  - The last item on the list
		sumaTotal += list[i].price
		if list[i].name != "" {
			ultimoProducto = i
		}
		if i+1 == len(list) {
			fmt.Println(list[ultimoProducto].name, "es el ultimo producto.")
			//  - The total number of items
			fmt.Println("El costo total fue de", sumaTotal)
		}
	}
}

func main() {
	//--Requirements:
	//* Using an array, create a shopping list with enough room
	//  for 4 products
	shoppingList := [4]Product{
		//* Insert 3 products into the array
		{"carro", 1000},
		{"avion", 2000},
		{"cohete", 4000},
	}

	checkProductList(shoppingList)

	//* Add a fourth product to the list and print out the
	shoppingList[3].name = "nave espacial"
	shoppingList[3].price = 5000
	//  information again
	checkProductList(shoppingList)

}
