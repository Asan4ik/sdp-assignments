package main

// Orchestrates common build steps using the CarBuilder interface
type Director struct{}

func NewDirector() *Director {
	return &Director{}
}

// Configures sports car options
func (d *Director) BuildSportsCar(builder CarBuilder) {
	builder.SetSeats(2).
		SetEngine("V10 Naturally Aspirated", 610).
		AddGPS().
		AddSportPackage()
}

// Configures electric vehicle options
func (d *Director) BuildEV(builder CarBuilder) {
	builder.SetSeats(5).
		SetEngine("Tri-Motor Plaid", 1020).
		AddGPS().
		AddSunroof()
}