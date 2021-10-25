# Template Method Pattern

## Go

```go
package main

import (
	"fmt"
)

type Writer interface {
	Write(data string)
}

type StdoutWriter struct{}

func (w *StdoutWriter) Write(data string) {
	fmt.Println(data)
}

func NewStdoutWriter() Writer {
	return &StdoutWriter{}
}

type Camera struct {
	Writer
}

func (c *Camera) Take() {
	c.Writer.Write("Saving photo.jpeg ...")
}

func main() {
	camera := &Camera{Writer: NewStdoutWriter()}

	camera.Take()
	camera.Take()
}
```

```
Saving photo.jpeg ...
Saving photo.jpeg ...
```
