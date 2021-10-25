# Strategy Pattern

## Go

```go
package main

import (
	"fmt"
)

type FilterFn func()

func UVFilter() {
	fmt.Println("Applying UV filter ...")
}

func PolarizingFilter() {
	fmt.Println("Applying Polarizing filter ...")
}

type Camera struct {
	filter FilterFn
}

func (c *Camera) SetFilter(filter FilterFn) {
	c.filter = filter
}

func (c *Camera) Take() {
	fmt.Println("Taking a photo ...")

	if c.filter != nil {
		c.filter()
	}
}

func main() {
	camera := &Camera{}

	camera.SetFilter(UVFilter)
	camera.Take()
	camera.SetFilter(PolarizingFilter)
	camera.Take()
}
```

```
Taking a photo ...
Applying UV filter ...
Taking a photo ...
Applying Polarizing filter ...
```
