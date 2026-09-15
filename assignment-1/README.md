# Car Builder

A simple Builder pattern example in Go. Instead of one huge `Car` constructor with a million params, you build cars step by step with chained method calls.

## Files

- `car.go` — the `Car` struct itself, plus a `Print()` to dump its specs
- `builder.go` — the `CarBuilder` interface and two builders: `SportsCarBuilder` and `ElectricCarBuilder`
- `director.go` — `Director`, which knows the standard build steps for each car type so you don't have to repeat them
- `main.go` — runs it

## How it works

`SportsCarBuilder` makes a 2-seat, high-horsepower car and won't let you build one under 300 HP.

`ElectricCarBuilder` makes a 5-seat EV with a battery, and has an extra `SetBatteryCapacity` step that's not part of the shared interface. Won't build with less than 40 kWh.

Example:

```go
sportsBuilder := NewSportsCarBuilder("Porsche", "911 GT3")
director.BuildSportsCar(sportsBuilder)
sportsCar, err := sportsBuilder.Build()
```

```go
evBuilder := NewElectricCarBuilder("Tesla", "Model S")
evBuilder.SetBatteryCapacity(100)
director.BuildEV(evBuilder)
evCar, err := evBuilder.Build()
```

`Build()` returns an error if the car doesn't meet the builder's rules.

## How to run

```bash
go run .
```

Builds one sports car and one EV, prints both.

## Requirements

Go 1.18+