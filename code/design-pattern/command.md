# Command Pattern

## Go

```go
package main

import (
	"fmt"
)

type Camera struct{}

func (c *Camera) On() {
	fmt.Println("Turning on camera ...")
}

func (c *Camera) Off() {
	fmt.Println("Turning off camera ...")
}

type Command func()

func NewOpenCameraCommand(camera *Camera) Command {
	return func() {
		camera.On()
	}
}

func NewCloseCameraCommand(camera *Camera) Command {
	return func() {
		camera.Off()
	}
}

func main() {
	camera := &Camera{}
	openCommand := NewOpenCameraCommand(camera)
	closeCommand := NewCloseCameraCommand(camera)

	openCommand()
	closeCommand()
}
```

```
Turning on camera ...
Turning off camera ...
```
