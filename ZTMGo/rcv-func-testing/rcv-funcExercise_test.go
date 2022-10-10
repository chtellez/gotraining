//--Requirements:

package rcvFunc

import "testing"

//* Write unit tests that ensure:

//  - Health & energy can not go above their maximums
func TestIncreaseHealthAndEnergy(t *testing.T) {
	//New player with health and energy of 0 and maxHealth and maxEnergy of 100
	p := Player{"Carlos", 0,100,0,100}

	for i := 0; i < 100; i++ {
		//Check health of player
		if i != p.health {
			t.Errorf("Health should not be %v, wait %v", p.health, i)
		}

		//Check energy of player
        if i!= p.energy {
			t.Errorf("Energy should not be %v, wait %v", p.energy, i)
		}

		//increase health by 1
		p.increaseHealth(1)
		if i >= p.health {
			t.Errorf("Error increasing health: got %v, want %v", p.health, i)
		}

		//increase energy by 1
        p.increaseEnergy(1)
		if i >= p.energy {
			t.Errorf("Error increasing energy: got %v, want %v", p.energy, i)
		}
		
	}
	
	// Health should be maxHealth
	if p.health != p.maxHealth {
		t.Errorf("Health must be maxHealth, got %v, wait %v", p.health, p.maxHealth)
	}

	// Energy should be maxEnergy
    if p.energy!= p.maxEnergy {
		t.Errorf("Energy must be maxEnergy, got %v, wait %v", p.energy, p.maxEnergy)
	}

	
	//Helth could not increase when is already at maxHealth
	p.increaseHealth(1)
	if p.health > p.maxHealth {
		t.Errorf("Health should not be greater than maxHealth, got %v, wait %v", p.health, p.maxHealth)
	}

	//Energy could not increase when is already at maxEnergy
	p.increaseEnergy(1)
    if p.energy > p.maxEnergy {
        t.Errorf("Energy should not be greater than maxEnergy, got %v, wait %v", p.energy, p.maxEnergy)
    }

}

//  - Health & energy can not go below 0
func TestDecreaseHealthAndEnergy(t *testing.T) {
	//New player with health and energy of 100 and maxHealth and maxEnergy of 100
	p := Player{"Carlos", 100,100,100,100}

	for i := 100; i > 0; i-- {
		//Check health of player
		if i != p.health {
			t.Errorf("Health should not be %v, wait %v", p.health, i)
		}

		//Check energy of player
        if i!= p.energy {
			t.Errorf("Energy should not be %v, wait %v", p.energy, i)
		}

		//decrease health by 1
		p.decreaseHealth(1)
		if i <= p.health {
			t.Errorf("Error decreasing health: got %v, want %v", p.health, i)
		}

		//increase energy by 1
        p.decreaseEnergy(1)
		if i <= p.energy {
			t.Errorf("Error decreasing energy: got %v, want %v", p.energy, i)
		}
		
	}
	
	// Health should be 0
	if p.health != 0 {
		t.Errorf("Health must be 0, got %v, wait %v", p.health, p.maxHealth)
	}

	// Energy should be 0
    if p.energy != 0 {
		t.Errorf("Energy must be 0, got %v, wait %v", p.energy, p.maxEnergy)
	}

	
	//Helth could not decrease when is already at 0
	p.decreaseHealth(1)
	if p.health > p.maxHealth {
		t.Errorf("Health should not be less than 0, got %v, wait %v", p.health, p.maxHealth)
	}

	//Energy could not decrease when is already at 0
	p.decreaseEnergy(1)
    if p.energy > p.maxEnergy {
        t.Errorf("Energy should not be less than 0, got %v, wait %v", p.energy, p.maxEnergy)
    }

}





//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests