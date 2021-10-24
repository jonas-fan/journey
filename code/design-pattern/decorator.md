# Decorator Pattern

## Go

```go
package main

import (
	"fmt"
)

type CameraPackage interface {
	fmt.Stringer
	Price() uint
}

type Camera struct{}

func (c *Camera) String() string {
	return "Camera"
}

func (c *Camera) Price() uint {
	return 1000
}

func NewCamera() *Camera {
	return &Camera{}
}

type CameraWithLens struct {
	*Camera
}

func (c *CameraWithLens) String() string {
	return fmt.Sprintf("%s + Lens", c.Camera)
}

func (c *CameraWithLens) Price() uint {
	return c.Camera.Price() + 499
}

func NewCameraWithLens(camera *Camera) *CameraWithLens {
	return &CameraWithLens{
		Camera: camera,
	}
}

func main() {
	var item1 CameraPackage = NewCamera()
	var item2 CameraPackage = NewCameraWithLens(NewCamera())

	fmt.Printf("Package %s: $%d\n", item1, item1.Price())
	fmt.Printf("Package %s: $%d\n", item2, item2.Price())
}
```

```
Package Camera: $1000
Package Camera + Lens: $1499
```
