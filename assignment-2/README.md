# Factory Method & Abstract Factory

Assignment 2 for Software Design Patterns. Topic: Logistics / Transport

Implemented in Go

## Structure

- `transport/` - Factory Method.
  - `Transport` is the product interface, `Truck` and `Ship` are the concrete products
  - `Logistics` is the creator interface, `RoadLogistics` and `SeaLogistics` are the
    concrete creators
  - `PlanDelivery` is the shared logic that works with any `Logistics`, without
    knowing whether it got a truck or a ship

- `gui/` - Abstract Factory
  - `Button` and `Checkbox` are the abstract products
  - `WindowsFactory`, `MacFactory` and `LinuxFactory` are the concrete factories, each producing
    a matching set of Windows or Mac components

- `main.go` - a small client that exercises both patterns

## How to run

```
go run main.go
```

## Why these patterns

Part A only needs to decide which single transport to create, so Factory Method
is enough. Part B needs a whole family of components (button + checkbox) that
must always come from the same platform - Abstract Factory guarantees that
consistency, which a single Factory Method could not do on its own

## Clean Code examples

- Names spell out roles instead of abbreviating (`RoadLogistics`, `deliveryTransport`,
  `windowsApplication`) so a reader can tell creator from product from client at a glance
- No type switches in client code - `Application` and `main` only call interface
  methods, so adding a new transport or platform never touches existing code
- Each factory method does exactly one thing: build and return one object
- Concrete factories (`WindowsFactory`, `MacFactory`, `LinuxFactory`) build their whole family the
  same way, so there is no duplicated logic between them
