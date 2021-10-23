# Abstract Factory Pattern

## Go

```go
package main

import (
	"fmt"
)

// Car and factory definition
type Engine interface {
	Vendor() string
}

type Wheel interface {
	Vendor() string
}

type CarFactory interface {
	NewEngine() Engine
	NewWheel() Wheel
}

// Toyota factory
type ToyotaEngine struct{}

func (e *ToyotaEngine) Vendor() string {
	return "Toyota"
}

type ToyotaWheel struct{}

func (w *ToyotaWheel) Vendor() string {
	return "Toyota"
}

type ToyotaFactory struct{}

func (f *ToyotaFactory) NewEngine() Engine {
	return &ToyotaEngine{}
}

func (f *ToyotaFactory) NewWheel() Wheel {
	return &ToyotaWheel{}
}

func FromToyota() CarFactory {
	return &ToyotaFactory{}
}

// Producer
func ProduceCar(factory CarFactory) {
	engine := factory.NewEngine()
	wheel := factory.NewWheel()

	fmt.Printf("Produced an engine of %q\n", engine.Vendor())
	fmt.Printf("Produced a wheel of %q\n", wheel.Vendor())
}

func main() {
	ProduceCar(FromToyota())
}
```
