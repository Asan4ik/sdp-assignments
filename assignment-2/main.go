package main

import (
	"fmt"
	"assignment-2/transport"
)

func main() {
	fmt.Println("--- Part A: Factory Method ---")

	roadDeliveryResult := transport.PlanDelivery(transport.RoadLogistics{})
	fmt.Println(roadDeliveryResult)

	seaDeliveryResult := transport.PlanDelivery(transport.SeaLogistics{})
	fmt.Println(seaDeliveryResult)

	fmt.Println()
}
