package transport

// Ship is a concrete product. It delivers goods by sea.
type Ship struct{}

func (ship Ship) Deliver() string {
	return "Delivering by sea in a container"
}
