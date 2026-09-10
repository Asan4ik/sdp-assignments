package main

import (
	"fmt"
	"errors"
)

// Holds intermediate state during construction
type CarBuilder struct {
	car Car
}

// Creates a builder with mandatory fields and predefined defaults
func NewCarBuilder(make, model string) *CarBuilder {
	return &CarBuilder {
		car: Car {
			Make:       make,
			Model:      model,
			Seats:      4,
			EngineType: "V4",
			Horsepower: 150,
			HasGPS:     false,
			HasSunroof: false,
			IsSportPkg: false,
		},
	}
}

// Getters and Setters
func (b *CarBuilder) SetSeats(seats int) *CarBuilder {
	b.car.Seats = seats
	return b
}

func (b *CarBuilder) SetEngine(engineType string, horsepower int) *CarBuilder {
	b.car.EngineType = engineType
	b.car.Horsepower = horsepower
	return b
}

func (b *CarBuilder) AddGPS() *CarBuilder {
	b.car.HasGPS = true
	return b
}

func (b *CarBuilder) AddSunroof() *CarBuilder {
	b.car.HasSunroof = true
	return b
}

func (b *CarBuilder) AddSportPackage() *CarBuilder {
	b.car.IsSportPkg = true
	return b
}
//

// Validates car constraints and returns its final instance
func (b *CarBuilder) Build() (*Car, error) {
	if b.car.Seats <= 0 {
		return nil, errors.New("A car must have at least 1 seat")
	}

	if b.car.IsSportPkg && b.car.Horsepower < 300 {
		return nil, fmt.Errorf("Sport package requires at least 300 HP. Current HP is %d", b.car.Horsepower)
	}

	builtCar := b.car
	return &builtCar, nil
}
