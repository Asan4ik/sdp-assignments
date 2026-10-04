package main
 
import "fmt"
 
func main() {
	vector := &VectorRenderer{}
	raster := &RasterRenderer{}
 
	/* The client composes a RefinedAbstraction with a ConcreteImplementor
	   at runtime. Circle and Square know nothing about VectorRenderer
	   or RasterRenderer directly, only about the Renderer interface */
	circle := NewCircle(vector, 5)
	square := NewSquare(raster, 10)
 
	fmt.Println("Initial renderers:")
	circle.Draw()
	square.Draw()
 
	fmt.Println("\nResizing shapes (abstraction-side logic, untouched by rendering):")
	circle.Resize(2)
	square.Resize(0.5)
	circle.Draw()
	square.Draw()
 
	fmt.Println("\nSwitching the circle to a different renderer at runtime,")
	fmt.Println("with no change to the Circle type itself:")
	circle.SetRenderer(raster)
	circle.Draw()
 
	fmt.Println("\nAdding a new shape built on the existing renderers,")
	fmt.Println("with no change to Renderer, VectorRenderer, or RasterRenderer:")
	square2 := NewSquare(vector, 3)
	square2.Draw()
}
 