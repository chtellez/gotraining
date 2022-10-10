//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//

package rcvFunc

import "fmt"

//--Requirements:
//* Implement a Player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name

type Player struct {
	name                                 string
	health, maxHealth, energy, maxEnergy int
}

//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the Player.

func (p *Player) increaseHealth(amount int) {
	if p.health < p.maxHealth {
		p.health += amount
	}
	//p.printStatistics()
}

func (p *Player) decreaseHealth(amount int) {
	if p.health > 0 {
        p.health -= amount
    }
	//p.printStatistics()
}

func (p *Player) increaseEnergy(amount int) {
	if p.energy < p.maxEnergy {
        p.energy += amount
    }
	//p.printStatistics()
}

func (p *Player) decreaseEnergy(amount int) {
	if p.energy > 0 {
        p.energy -= amount
    }
	//p.printStatistics()
}

// - Print out the statistic change within each function
func (p *Player) printStatistics() {
	fmt.Println(p.name, "health:", p.health, "/", p.maxHealth)
	fmt.Println(p.name, "energy:", p.energy, "/", p.maxEnergy)
}

//  - Execute each function at least once
