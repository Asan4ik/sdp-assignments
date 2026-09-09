package main

import "fmt"

/* Car represents the complex object 
   being constructed              */
type Car struct {
	Make        string
	Model       string
	Seats       int
	EngineType  string
	HoursePower int
	HasGPS      bool
	HasSunroof  bool
	IsSportPkg  bool
}


func (c *Car) PrintSpecs() {
	fmt.Printf("--- %s %s Specs ---", c.Make, c.Model)
	fmt.Printf("Engine: %s (%d HP), c.EngineType, c.Horsepower")
	fmt.Printf("Seats: %d\n", c.Seats)
	fmt.Printf("GPS: %t | Sunroof: %t | Sport Package: %t\n\n", c.HasGPS, c.HasSunroof, c.IsSportPkg)
}