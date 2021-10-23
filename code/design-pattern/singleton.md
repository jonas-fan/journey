# Singleton Pattern

Make sure the number of class instances is limited so they are allowed to be shared in an entire system.

## Go

```go
package main

import (
	"fmt"
	"sync"
)

type Camera struct {
	Vendor string
}

var (
	camera     *Camera
	cameraOnce sync.Once
)

func GetCamera() *Camera {
	cameraOnce.Do(func() {
		camera = &Camera{
			Vendor: "Nikon",
		}
	})

	return camera
}

func main() {
	fmt.Printf("Instance: %p\n", GetCamera())
	fmt.Printf("Instance: %p\n", GetCamera())
	fmt.Printf("Instance: %p\n", GetCamera())
}

```
