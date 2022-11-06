//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//

package main

import "fmt"

//--Requirements:

const (
	SmallLift = iota
	StandarLift
	LargeLift
)


const (
	Motorcycles = iota
	Cars
	Trucks
)
type Lift struct {
	name string
	liftType int
}

var lifts = []Lift{
	{name: "SmallLift", liftType: SmallLift},
	{name: "StandarLift", liftType: StandarLift},
	{name: "LargeLift", liftType: LargeLift},
}

type LiftAsigner interface {
	LiftPicker()
	PrintLift()
}
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
type Vehicle struct {
	vType int
	model string
	lift Lift
	LiftAsigner
}

//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts

func (v *Vehicle) LiftPicker(){
	switch v.vType {
	case Motorcycles:
		v.lift = lifts[SmallLift]
	case Cars:
		v.lift = lifts[StandarLift]
	case Trucks:
		v.lift = lifts[LargeLift]
	default:
		fmt.Println("Vehicle type not handle")
	}

}

func (v *Vehicle) PrintLift() {
	fmt.Println(v.model, "assigned to", v.lift.name )
}

//* Write a single function to handle all of the vehicles
//  that the shop works on.

func assignLift( vehicles []LiftAsigner) {
	for i := 0; i < len(vehicles); i++ {
		vehicles[i].LiftPicker()
		vehicles[i].PrintLift()
	}
}

//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
func main (){
	
	mazda := Vehicle{model:"Mazda", vType:Cars}
	Kawa := Vehicle{model:"Kawasaky", vType:Motorcycles}
	Kentworth := Vehicle{model:"Kentworth", vType:Trucks}

	vehicles := []LiftAsigner{&mazda, &Kawa, &Kentworth}

	assignLift(vehicles)
}

//--Notes:
//* Use any names for vehicle models