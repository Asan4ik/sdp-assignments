package main

import (
	"fmt"

	"assignment-2/gui"
	"assignment-2/transport"
)

/* Application is the Abstract Factory client. It only ever talks to
   the GUIFactory interface and the abstract Button/Checkbox
   interfaces - it never mentions WindowsButton or MacButton directly,
   so it works the same way no matter which family it was given. */
type Application struct {
	factory gui.GUIFactory
}

func NewApplication(factory gui.GUIFactory) Application {
	return Application{factory: factory}
}

func (application Application) Render() {
	button := application.factory.CreateButton()
	checkbox := application.factory.CreateCheckbox()

	fmt.Println(button.Render())
	fmt.Println(checkbox.Render())
}

func main() {
	fmt.Println("--- Part A: Factory Method ---")

	roadDeliveryResult := transport.PlanDelivery(transport.RoadLogistics{})
	fmt.Println(roadDeliveryResult)

	seaDeliveryResult := transport.PlanDelivery(transport.SeaLogistics{})
	fmt.Println(seaDeliveryResult)

	fmt.Println()
	fmt.Println("--- Part B: Abstract Factory ---")

	windowsApplication := NewApplication(gui.WindowsFactory{})
	windowsApplication.Render()

	macApplication := NewApplication(gui.MacFactory{})
	macApplication.Render()
}
