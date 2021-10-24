# Facade Pattern

## Go

```go
package main

import (
	"fmt"
)

type Sensor struct{}

func (s *Sensor) Read() {
	fmt.Println("Reading data")
}

type Card struct{}

func (c *Card) Save() {
	fmt.Println("Saving data")
}

type Capturer struct {
	*Sensor
	*Card
}

func (c *Capturer) Do() {
	c.Sensor.Read()
	c.Card.Save()
}

func NewCapturer() *Capturer {
	return &Capturer{
		Sensor: &Sensor{},
		Card:   &Card{},
	}
}

type Camera struct {
	*Capturer
}

func (c *Camera) Take() {
	c.Capturer.Do()
}

func NewCamera() *Camera {
	return &Camera{
		Capturer: NewCapturer(),
	}
}

func main() {
	cam := NewCamera()
	cam.Take()
}
```

```
Reading data
Saving data
```
