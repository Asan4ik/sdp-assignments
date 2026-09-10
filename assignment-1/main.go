package main

import "fmt"

func main() {
	normalCar, err := NewCarBuilder("Kia", "K5").
			SetSeats(4).
			AddGPS().
			Build()
	
	if err != nil {
		fmt.Println("Error building normal car:", err)
	} else {
		normalCar.Print()
	}

	sportsCar, err := NewCarBuilder("Porsche", "911").
		SetSeats(2).
		SetEngine("Twin-Turbo Flat-6", 443).
		AddGPS().
		AddSunroof().
		AddSportPackage().
		Build()
	
	if err != nil {
		fmt.Println("Error building sports car:", err)
	} else {
		sportsCar.Print()
	}
}
