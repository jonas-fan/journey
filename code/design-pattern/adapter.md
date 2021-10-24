# Adapter Pattern

## Go

```go
package main

import (
	"fmt"
)

type Camera interface {
	Take() string
}

type NikonCamera struct{}

func (c *NikonCamera) TakePhoto() string {
	return "took a photo"
}

func NewNikonCamera() *NikonCamera {
	return &NikonCamera{}
}

type NikonCameraAdapter struct {
	*NikonCamera
}

func (a *NikonCameraAdapter) Take() string {
	return a.NikonCamera.TakePhoto()
}

func NewNikonCameraAdapter(cam *NikonCamera) *NikonCameraAdapter {
	return &NikonCameraAdapter{
		NikonCamera: cam,
	}
}

func main() {
	var cam Camera

	cam = NewNikonCameraAdapter(NewNikonCamera())

	for i := 0; i < 3; i++ {
		fmt.Printf("Taking ... %s\n", cam.Take())
	}
}
```

```
Taking ... took a photo
Taking ... took a photo
Taking ... took a photo
```
