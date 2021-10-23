# Builder Pattern

## Go

```go
package main

import (
	"fmt"
)

type Camera struct {
	Vendor string
	Model  string
}

type CameraBuilder struct {
	*Camera
}

func (b *CameraBuilder) Vendor(name string) *CameraBuilder {
	b.Camera.Vendor = name

	return b
}

func (b *CameraBuilder) Model(name string) *CameraBuilder {
	b.Camera.Model = name

	return b
}

func (b *CameraBuilder) Build() *Camera {
	return b.Camera
}

func NewCameraBuilder() *CameraBuilder {
	return &CameraBuilder{
		Camera: &Camera{},
	}
}

func main() {
	builder := NewCameraBuilder()
	builder.Vendor("Nikon").Model("D750")

	camera := builder.Build()

	fmt.Printf("Produced a \"%s %s\" camera\n", camera.Vendor, camera.Model)
}
```

```
Produced a "Nikon D750" camera
```
