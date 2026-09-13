package main

import (
	"errors"
	"fmt"
)

// CarBuilder defines common steps for building a Car
type CarBuilder interface {
	SetSeats(seats int) CarBuilder
	SetEngine(engineType string, horsepower int) CarBuilder
	AddGPS() CarBuilder
	AddSunroof() CarBuilder
	AddSportPackage() CarBuilder
	Build() (*Car, error)
}

// Concrete Builder 1: SportsCarBuilder
type SportsCarBuilder struct {
	car Car
}

func NewSportsCarBuilder(make, model string) *SportsCarBuilder {
	return &SportsCarBuilder{
		car: Car{
			Make:       make,
			Model:      model,
			Seats:      2,
			EngineType: "V8 Twin-Turbo",
			Horsepower: 450,
			HasGPS:     true,
			HasSunroof: false,
			IsSportPkg: true,
		},
	}
}

func (b *SportsCarBuilder) SetSeats(seats int) CarBuilder {
	b.car.Seats = seats
	return b
}

func (b *SportsCarBuilder) SetEngine(engineType string, horsepower int) CarBuilder {
	b.car.EngineType = engineType
	b.car.Horsepower = horsepower
	return b
}

func (b *SportsCarBuilder) AddGPS() CarBuilder {
	b.car.HasGPS = true
	return b
}

func (b *SportsCarBuilder) AddSunroof() CarBuilder {
	b.car.HasSunroof = true
	return b
}

func (b *SportsCarBuilder) AddSportPackage() CarBuilder {
	b.car.IsSportPkg = true
	return b
}

func (b *SportsCarBuilder) Build() (*Car, error) {
	if b.car.Seats <= 0 {
		return nil, errors.New("sports car must have at least 1 seat")
	}
	if b.car.Horsepower < 300 {
		return nil, fmt.Errorf("sports car requires at least 300 HP, got %d", b.car.Horsepower)
	}
	builtCar := b.car
	return &builtCar, nil
}

// Concrete Builder 2: ElectricCarBuilder
type ElectricCarBuilder struct {
	car Car
}

func NewElectricCarBuilder(make, model string) *ElectricCarBuilder {
	return &ElectricCarBuilder{
		car: Car{
			Make:       make,
			Model:      model,
			Seats:      5,
			EngineType: "Dual Motor EV",
			Horsepower: 300,
			BatteryKWh: 75,
			HasGPS:     true,
			HasSunroof: true,
			IsSportPkg: false,
		},
	}
}

func (b *ElectricCarBuilder) SetSeats(seats int) CarBuilder {
	b.car.Seats = seats
	return b
}

func (b *ElectricCarBuilder) SetEngine(engineType string, horsepower int) CarBuilder {
	b.car.EngineType = engineType
	b.car.Horsepower = horsepower
	return b
}

func (b *ElectricCarBuilder) SetBatteryCapacity(kWh int) *ElectricCarBuilder {
	b.car.BatteryKWh = kWh
	return b
}

func (b *ElectricCarBuilder) AddGPS() CarBuilder {
	b.car.HasGPS = true
	return b
}

func (b *ElectricCarBuilder) AddSunroof() CarBuilder {
	b.car.HasSunroof = true
	return b
}

func (b *ElectricCarBuilder) AddSportPackage() CarBuilder {
	b.car.IsSportPkg = true
	return b
}

func (b *ElectricCarBuilder) Build() (*Car, error) {
	if b.car.Seats <= 0 {
		return nil, errors.New("Electric car must have at least 1 seat")
	}
	if b.car.BatteryKWh < 40 {
		return nil, fmt.Errorf("Electric car battery must be at least 40 kWh, got %d", b.car.BatteryKWh)
	}
	builtCar := b.car
	return &builtCar, nil
}