# Visitor Pattern

## Go

```go
package main

import (
	"fmt"
)

type CameraVisitor interface {
	VisitModel(m *CameraModel)
	VisitLens(l *CameraLens)
}

type NikonCameraVisitor struct{}

func (v *NikonCameraVisitor) VisitModel(m *CameraModel) {
	fmt.Printf("Nikon: visiting model %q\n", m.Model)
}

func (v *NikonCameraVisitor) VisitLens(l *CameraLens) {
	fmt.Printf("Nikon: visiting lens %q\n", l.Lens)
}

type CameraModel struct {
	Model string
}

type CameraLens struct {
	Lens string
}

type Camera struct {
	*CameraModel
	*CameraLens
}

func (c *Camera) Visit(v CameraVisitor) {
	v.VisitModel(c.CameraModel)
	v.VisitLens(c.CameraLens)
}

func main() {
	var vistor CameraVisitor = &NikonCameraVisitor{}

	camera := &Camera{
		CameraModel: &CameraModel{Model: "Nikon D750"},
		CameraLens:  &CameraLens{Lens: "35mm f/1.8"},
	}

	camera.Visit(vistor)
}
```

```
Nikon: visiting model "Nikon D750"
Nikon: visiting lens "35mm f/1.8"
```
