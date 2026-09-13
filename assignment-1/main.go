package main

import "fmt"

func main() {
	director := NewDirector()

	// Direct usage of SportsCarBuilder
	sportsBuilder := NewSportsCarBuilder("Porsche", "911 GT3")
	director.BuildSportsCar(sportsBuilder)

	sportsCar, err := sportsBuilder.Build()
	if err != nil {
		fmt.Println("Error building sports car:", err)
	} else {
		sportsCar.Print()
	}

	// Direct usage of ElectricCarBuilder with specific battery setup
	evBuilder := NewElectricCarBuilder("Tesla", "Model S")
	evBuilder.SetBatteryCapacity(100)
	director.BuildEV(evBuilder)

	evCar, err := evBuilder.Build()
	if err != nil {
		fmt.Println("Error building EV:", err)
	} else {
		evCar.Print()
	}
}