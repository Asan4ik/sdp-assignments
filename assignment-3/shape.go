package main

type Shape interface {
	Draw()
	Resize(factor float64)
}

type baseShape struct {
	renderer Renderer
}

func (b *baseShape) SetRenderer(renderer Renderer) {
	b.renderer = renderer
}

type Circle struct {
	baseShape
	Radius float64
}

func NewCircle(renderer Renderer, radius float64) *Circle {
	return &Circle{baseShape: baseShape{renderer: renderer}, Radius: radius}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.Radius)
}

func (c *Circle) Resize(factor float64) {
	c.Radius *= factor
}

/* Square is a RefinedAbstraction, structured the same way as Circle but
   with its own shape-specific state */
type Square struct {
	baseShape
	Side float64
}
 
// NewSquare composes a Square with a chosen Renderer at construction time
func NewSquare(renderer Renderer, side float64) *Square {
	return &Square{baseShape: baseShape{renderer: renderer}, Side: side}
}
 
func (s *Square) Draw() {
	s.renderer.RenderSquare(s.Side)
}
 
func (s *Square) Resize(factor float64) {
	s.Side *= factor
}