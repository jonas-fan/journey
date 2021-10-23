# Factory Method Pattern

## Go

```go
package main

import (
	"fmt"
)

// Camera interface and factory definition
type Camera interface {
	Vendor() string
}

type CameraFactoryFn func() Camera

// Nikon camera and factory
type NikonCamera struct{}

func (n *NikonCamera) Vendor() string {
	return "Nikon"
}

func FromNikon() Camera {
	return &NikonCamera{}
}

// Canon camera and factory
type CanonCamera struct{}

func (c *CanonCamera) Vendor() string {
	return "Canon"
}

func FromCanon() Camera {
	return &CanonCamera{}
}

// Producer
func ProduceCamera(factoryFn CameraFactoryFn, num uint) {
	for ; num != 0; num-- {
		camera := factoryFn()

		fmt.Printf("Produced a camera of %q\n", camera.Vendor())
	}
}

func main() {
	ProduceCamera(FromNikon, 2)
	ProduceCamera(FromCanon, 2)
}
```
