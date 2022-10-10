//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//

package main

import "fmt"

//--Requirements:
//* Implement a Player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name

type Player struct {
	name                                 string
	health, maxHealth, energy, maxEnergy uint
}

//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the Player.

func (p *Player) increaseHealth(health uint) {
	p.health += health
	p.printStatistics()
}

func (p *Player) decreaseHealth(health uint) {
	p.health -= health
	p.printStatistics()
}

func (p *Player) increaseEnergy(energy uint) {
	p.energy += energy
	p.printStatistics()
}

func (p *Player) decreaseEnergy(energy uint) {
	p.energy -= energy
	p.printStatistics()
}

// - Print out the statistic change within each function
func (p *Player) printStatistics() {
	fmt.Println(p.name, "health:", p.health, "/", p.maxHealth)
	fmt.Println(p.name, "energy:", p.energy, "/", p.maxEnergy)
}

//  - Execute each function at least once

func main() {
	christian := Player{name: "Christian", health: 100, energy: 100, maxHealth: 100, maxEnergy: 100}
	fmt.Println("Jugador", christian.name, "fue creado!")
	christian.printStatistics()

	fmt.Println("\nImpacto directo recibido de 80")
	christian.decreaseHealth(80)

	fmt.Println("\nAtaque de 50 de energía")
	christian.decreaseEnergy(50)

	fmt.Println("\nPosión de vida de 60")
	christian.increaseHealth(60)

	fmt.Println("\nPosión de energía de 40")
	christian.increaseEnergy(40)
}
