package transport

// Truck is a concrete product. It delivers goods by road.
type Truck struct{}

func (truck Truck) Deliver() string {
	return "Delivering by land in a box"
}
