package transport

// SeaLogistics is a concrete creator that produces ships instead of trucks.
type SeaLogistics struct{}

func (seaLogistics SeaLogistics) CreateTransport() Transport {
	return Ship{}
}
