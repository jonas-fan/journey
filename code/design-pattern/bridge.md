# Bridge Pattern

## Go

```go
package main

import (
	"fmt"
)

type Lens interface {
	Focus()
}

type Camera interface {
	Take()
}

type NikonLens struct{}

func (l *NikonLens) Focus() {
	fmt.Println("Lens focusing ...")
}

func NewNikonLens() *NikonLens {
	return &NikonLens{}
}

type NikonCamera struct {
	Lens
}

func (c *NikonCamera) Take() {
	c.Lens.Focus()

	fmt.Println("Taking a photo ...")
}

func NewNikonCamera(lens Lens) *NikonCamera {
	return &NikonCamera{
		Lens: lens,
	}
}

func main() {
	var cam Camera

	cam = NewNikonCamera(NewNikonLens())
	cam.Take()
}
```

```
Lens focusing ...
Taking a photo ...
```
