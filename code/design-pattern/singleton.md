# Singleton Pattern

## Go

```go
package main

import (
	"fmt"
	"sync"
)

var (
	camera     *Camera
	cameraOnce sync.Once
)

type Camera struct {
	Vendor string
}

func NewCamera(vendor string) *Camera {
	return &Camera{
		Vendor: vendor,
	}
}

func GetCamera() *Camera {
	cameraOnce.Do(func() {
		camera = NewCamera("Nikon")
	})

	return camera
}

func main() {
	for i := 0; i < 3; i++ {
		cam := GetCamera()

		fmt.Printf("A camera(%p) made by %s\n", cam, cam.Vendor)
	}
}
```

```
A camera(0xc0000a6210) made by Nikon
A camera(0xc0000a6210) made by Nikon
A camera(0xc0000a6210) made by Nikon
```
