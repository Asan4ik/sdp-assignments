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

func (c *Cirlce) Draw() {
	c.renderer.RenderCircle(c.Radius)
}

func (c *Circle) Resize(factor float64) {
	c.Radius *= factor
}