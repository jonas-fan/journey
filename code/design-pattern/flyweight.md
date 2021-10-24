# Flyweight Pattern

## Go

```go
package main

import (
	"fmt"
)

type Camera struct {
	Model string
}

func (c *Camera) String() string {
	return c.Model
}

func NewCamera(model string) *Camera {
	return &Camera{
		Model: model,
	}
}

type CameraFactory struct {
	Cameras map[string]*Camera
}

func (f *CameraFactory) Camera(model string) *Camera {
	camera, ok := f.Cameras[model]

	if !ok {
		camera = NewCamera(model)
		f.Cameras[model] = camera
	}

	return camera
}

func NewCameraFactory() *CameraFactory {
	return &CameraFactory{
		Cameras: make(map[string]*Camera),
	}
}

func spec(camera *Camera) {
	fmt.Printf("Camera(%p): %s\n", camera, camera)
}

func main() {
	factory := NewCameraFactory()

	spec(factory.Camera("Nikon z8"))
	spec(factory.Camera("Nikon z5"))
	spec(factory.Camera("Nikon z5"))
}
```

```
Camera(0xc000010230): Nikon z8
Camera(0xc000010240): Nikon z5
Camera(0xc000010240): Nikon z5
```
