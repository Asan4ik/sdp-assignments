package main

import (
	"fmt"
	"errors"
)

// CarBuilder holds intermediate state during construction
type CarBuilder struct {
	car Car
}

// NewCarBuilder creates a builder with mandatory fields and predefined defaults
func NewCarBuilder(make, model string) *CarBuilder {
	return &CarBuilder {
		car: Car {
			Make: make,
			Model: model,
			Seats: 4,
			EngineType: "V4",
			HorsePower: 150,
			HasGPS: false,
			HasSunroof: false,
			IsSportPkg: false,
		},
	}
}

// Getters and Setters
func (b *CarBuilder) SetSeats(seats int) *CarBuilder {
	b.car.seats = seats
	return b
}

func (b *CarBuilder) SetEngine(engineType string, horsePower int) *CarBuilder {
	b.car.engineType = engineType
	b.car.horsePower = horsePower
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

// Build validates Car constraints and returns its final instance
func (b *CarBuilder) Build() (*Car, error) {
	if b.car.seats <= 0 {
		return nil, errors.New("A car must have at least 1 seat")
	}

	if b.car.IsSportPkg && b.car.horsePower < 300 {
		return nil, fmt.Errorf("Sport package requires at least 300 HP. Current HP is %d", b.car.horsePower)
	}

	builtCar := b.car
	return &builtCar, nil
}