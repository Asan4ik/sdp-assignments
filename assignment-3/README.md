# Bridge Pattern: Shape and Renderer

A Go implementation of the Bridge structural design pattern for Assignment 3
(ShP-2216, Software Design Patterns).

## What this is

A `Shape` abstraction (`Circle`, `Square`) that can be drawn by two
interchangeable `Renderer` implementations (`VectorRenderer`,
`RasterRenderer`). The two hierarchies are connected by composition, not
inheritance, so a shape can switch renderers at runtime and a new renderer
can be added without changing any shape code.

## Project structure

```
.
├── shape.go      # Abstraction: Shape interface, baseShape, Circle, Square
├── renderer.go   # Implementor: Renderer interface, VectorRenderer, RasterRenderer
└── main.go       # Client: composes shapes with renderers and demonstrates the pattern
```

## Requirements

- Go 1.22 or later

## How to run

```bash
go run .
```

## What the demo shows

- `Circle` drawn with `VectorRenderer`, `Square` drawn with `RasterRenderer`
- Resizing both shapes without touching any rendering code
- Switching `Circle` to `RasterRenderer` at runtime via `SetRenderer`,
  with no change to the `Circle` type itself
- Adding a second `Square` reusing an existing renderer, showing that new
  shapes and new renderers can be combined freely
