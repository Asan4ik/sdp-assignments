package transport

/* RoadLogistics is a concrete creator. Its only job is deciding which
   transport gets created. */
type RoadLogistics struct{}

func (roadLogistics RoadLogistics) CreateTransport() Transport {
	return Truck{}
}
