package main

import "fmt"

// Car represents the complex object being constructed
type Car struct {
	Make        string
	Model       string
	Seats       int
	EngineType  string
	Horsepower  int
	BatteryKWh  int
	HasGPS      bool
	HasSunroof  bool
	IsSportPkg  bool
}

// Prints specs of a Car
func (car *Car) Print() {
	fmt.Printf("--- %s %s Specs ---\n", car.Make, car.Model)
	fmt.Printf("Engine/Power: %s (%d HP)\n", car.EngineType, car.Horsepower)
	if car.BatteryKWh > 0 {
		fmt.Printf("Battery Capacity: %d kWh\n", car.BatteryKWh)
	}
	fmt.Printf("Seats: %d\n", car.Seats)
	fmt.Printf("GPS: %t | Sunroof: %t | Sport Package: %t\n\n", car.HasGPS, car.HasSunroof, car.IsSportPkg)
}