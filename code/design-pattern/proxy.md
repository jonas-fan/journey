# Proxy Pattern

## Go

```go
package main

import (
	"fmt"
)

type Camera interface {
	fmt.Stringer
}

type NikonCameraPrivate struct{}

func (c *NikonCameraPrivate) String() string {
	return "Nikon"
}

type NikonCamera struct {
	priv *NikonCameraPrivate
}

func (c *NikonCamera) String() string {
	return c.priv.String()
}

func NewNikonCamera() *NikonCamera {
	return &NikonCamera{
		priv: &NikonCameraPrivate{},
	}
}

func spec(camera Camera) {
	fmt.Printf("Camera(%p): %s\n", camera, camera)
}

func main() {
	spec(NewNikonCamera())
}

```

```
Camera(0xc000122018): Nikon
```
