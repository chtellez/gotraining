//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//* The program must use variables to configure:
	//  - number of times to roll the dice
	//  - number of dice used in the rolls
	//  - number of sides of all the dice (6-sided, 10-sided, etc determined
	//    with a variable). All dice must use the same variable for number
	//    of sides, so they all have the same number of sides.
	var diceQty, rollTimes, diceSizes int

	fmt.Println("Bienvenido al mi programa de dados")
	fmt.Println("Para que el programa funcione adecuadamente, es necesario que ingreses los siguientes valores")
	fmt.Print("1/3 Cantidad de dados :")
	fmt.Scanf("%d", &diceQty)
	fmt.Print("2/3 Cantidad de lances :")
	fmt.Scanf("%d", &rollTimes)
	fmt.Print("3/3 Cantidad de caras :")
	fmt.Scanf("%d", &diceSizes)
	fmt.Printf("Usted ingreso:\n     Cantidad de dados: %d,\n     Cantidad de lances: %d,\n     Caras del dado: %d. \n", diceQty, rollTimes, diceSizes)
	fmt.Println("------------------------------------")

	var sumaLance, sumaTotal, caraDado int = 0, 0, 0
	for lance := 1; lance <= rollTimes; lance++ {
		sumaLance = 0
		fmt.Println("Lanzamiento No:", lance)
		for dado := 1; dado <= diceQty; dado++ {
			caraDado = 0
			rand.Seed(time.Now().UnixNano())
			for caraDado == 0 {
				caraDado = rand.Intn(diceSizes)
			}
			sumaLance += caraDado
			sumaTotal += caraDado
			fmt.Printf("	dado No: %d, valor: %d.\n", dado, caraDado)
		}
		//* Print the sum of the dice roll
		fmt.Printf("Total lancamiento No %d: %d \n", lance, sumaLance)

		//* Print additional information in these cirsumstances:
		//  - "Snake eyes": when the total roll is 2, and total dice is 2
		//  - "Lucky 7": when the total roll is 7
		//  - "Even": when the total roll is even
		//  - "Odd": when the total roll is odd
		if sumaLance == 2 && diceQty == 2 {
			fmt.Println("Snake eyes")
		} else if sumaLance == 7 {
			fmt.Println("Lucky 7")
		}
		if sumaLance%2 == 0 {
			fmt.Println("Resultado Par")
		} else {
			fmt.Println("Resultado Inpar")
		}
		fmt.Println("------------------------------------")
	}

}
