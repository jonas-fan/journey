# Factory Method Pattern

## Go

```go
package main

import (
	"fmt"
)

type Camera interface {
	Vendor() string
}

type CameraFactoryFn func() Camera

type NikonCamera struct{}

func (n *NikonCamera) Vendor() string {
	return "Nikon"
}

func FromNikon() Camera {
	return &NikonCamera{}
}

func main() {
	var factoryFn CameraFactoryFn

	factoryFn = FromNikon

	for i := 0; i < 3; i++ {
		camera := factoryFn()

		fmt.Printf("Produced a camera of %q\n", camera.Vendor())
	}
}
```

```
Produced a camera of "Nikon"
Produced a camera of "Nikon"
Produced a camera of "Nikon"
```
