package transport

/* Logistics is the creator interface. It declares the factory method
   that every concrete creator has to implement. Notice it never
   mentions Truck or Ship directly - it only works with the Transport
   interface. */
type Logistics interface {
	CreateTransport() Transport
}

/* PlanDelivery is the shared logic that works with any Logistics
   implementation. In a language with abstract classes this would live
   as a method on the base Creator class; in Go we just write a plain
   function that takes the interface. It never needs to know which
   concrete transport it ends up using. */
func PlanDelivery(logistics Logistics) string {
	deliveryTransport := logistics.CreateTransport()
	return deliveryTransport.Deliver()
}
