# Prototype Pattern

## Go

```go
package main

import (
	"fmt"
)

type Camera struct {
	Vendor string
}

type CameraCloneable interface {
	Clone() *Camera
}

func (c *Camera) Clone() *Camera {
	// deep copy will be better
	return &Camera{
		Vendor: c.Vendor,
	}
}

func main() {
	blueprint := &Camera{Vendor: "Nikon"}
	cam1 := blueprint.Clone()
	cam2 := blueprint.Clone()

	fmt.Printf("Produced a \"%s\" camera\n", cam1.Vendor)
	fmt.Printf("Produced a \"%s\" camera\n", cam2.Vendor)
}
```

```
Produced a "Nikon" camera
Produced a "Nikon" camera
```
