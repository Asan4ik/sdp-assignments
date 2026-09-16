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

func (builder *SportsCarBuilder) SetSeats(seats int) CarBuilder {
	builder.car.Seats = seats
	return builder
}

func (builder *SportsCarBuilder) SetEngine(engineType string, horsepower int) CarBuilder {
	builder.car.EngineType = engineType
	builder.car.Horsepower = horsepower
	return builder
}

func (builder *SportsCarBuilder) AddGPS() CarBuilder {
	builder.car.HasGPS = true
	return builder
}

func (builder *SportsCarBuilder) AddSunroof() CarBuilder {
	builder.car.HasSunroof = true
	return builder
}

func (builder *SportsCarBuilder) AddSportPackage() CarBuilder {
	builder.car.IsSportPkg = true
	return builder
}

func (builder *SportsCarBuilder) Build() (*Car, error) {
	if builder.car.Seats <= 0 {
		return nil, errors.New("sports car must have at least 1 seat")
	}
	if builder.car.Horsepower < 300 {
		return nil, fmt.Errorf("sports car requires at least 300 HP, got %d", builder.car.Horsepower)
	}
	builtCar := builder.car
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

func (builder *ElectricCarBuilder) SetSeats(seats int) CarBuilder {
	builder.car.Seats = seats
	return builder
}

func (builder *ElectricCarBuilder) SetEngine(engineType string, horsepower int) CarBuilder {
	builder.car.EngineType = engineType
	builder.car.Horsepower = horsepower
	return builder
}

func (builder *ElectricCarBuilder) SetBatteryCapacity(kWh int) *ElectricCarBuilder {
	builder.car.BatteryKWh = kWh
	return builder
}

func (builder *ElectricCarBuilder) AddGPS() CarBuilder {
	builder.car.HasGPS = true
	return builder
}

func (builder *ElectricCarBuilder) AddSunroof() CarBuilder {
	builder.car.HasSunroof = true
	return builder
}

func (builder *ElectricCarBuilder) AddSportPackage() CarBuilder {
	builder.car.IsSportPkg = true
	return builder
}

func (builder *ElectricCarBuilder) Build() (*Car, error) {
	if builder.car.Seats <= 0 {
		return nil, errors.New("Electric car must have at least 1 seat")
	}
	if builder.car.BatteryKWh < 40 {
		return nil, fmt.Errorf("Electric car battery must be at least 40 kWh, got %d", builder.car.BatteryKWh)
	}
	builtCar := builder.car
	return &builtCar, nil
}