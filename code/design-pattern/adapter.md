# Adapter Pattern

## Go

```go
package main

import (
	"fmt"
)

type Camera interface {
	Take()
}

type NikonCamera struct{}

func (c *NikonCamera) TakePhoto() {
	fmt.Println("Taking a photo")
}

func NewNikonCamera() *NikonCamera {
	return &NikonCamera{}
}

type NikonCameraAdapter struct {
	*NikonCamera
}

func (a *NikonCameraAdapter) Take() {
	a.NikonCamera.TakePhoto()
}

func NewNikonCameraAdapter(cam *NikonCamera) *NikonCameraAdapter {
	return &NikonCameraAdapter{
		NikonCamera: cam,
	}
}

func main() {
	var cam Camera

	cam = NewNikonCameraAdapter(NewNikonCamera())
	cam.Take()
}
```

```
Taking a photo
```
