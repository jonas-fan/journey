# Abstract Factory Pattern

## Go

```go
package main

import (
	"fmt"
)

type EngineFactory interface {
	NewEngine() string
}

type WheelFactory interface {
	NewWheel() string
}

type CarFactory interface {
	EngineFactory
	WheelFactory
}

type ToyotaCarFactory struct{}

func (f *ToyotaCarFactory) NewEngine() string {
	return "an engine made by Toyota"
}

func (f *ToyotaCarFactory) NewWheel() string {
	return "a wheel made by Toyota"
}

func FromToyota() CarFactory {
	return &ToyotaCarFactory{}
}

func main() {
	var factory CarFactory

	factory = FromToyota()

	fmt.Printf("Produced %s\n", factory.NewEngine())
	fmt.Printf("Produced %s\n", factory.NewWheel())
}
```

```
Produced an engine made by Toyota
Produced a wheel made by Toyota
```
