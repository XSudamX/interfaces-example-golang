package main

import "fmt"

// 						Interfaces allow a value to be of more than one type.					//

type registry interface { // Anything that has  the method inregistry also has the type registry
	inregistry()
}

type vehicle struct {
	modelname  string
	wheels     int
	passengers int
}
type airvehicle struct {
	vehicle
	carryextracargo bool
}

// A function when defined with an interface, will associate the interface as a TYPE to any value that has the method defined in the (same) function.
func (v vehicle) inregistry()    {}
func (p airvehicle) inregistry() {}

// POLYMORPHISM//

func bar(inr registry) {
	fmt.Println("I am in the registry and my details are ", inr)
}
func main() {

	vehicle1 := vehicle{
		"Range Rover",
		4,
		6,
	}
	vehicle2 := airvehicle{
		vehicle: vehicle{
			"AirBus 737",
			6,
			189,
		},
		carryextracargo: true,
	}
	bar(vehicle1)
	bar(vehicle2)
}

// We just passed in different "objects" to A SINGLE FUNCTION of DIFFERENT TYPES
// What we have achieved here is Polymorphism.
