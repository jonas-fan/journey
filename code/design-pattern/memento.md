# Memento Pattern

## Go

```go
package main

import (
	"fmt"
)

type Camera struct {
	state string
}

func (c *Camera) State() string {
	return c.state
}

func (c *Camera) SetState(state string) {
	fmt.Printf("Setting to %q state\n", state)
	c.state = state
}

func main() {
	states := make([]string, 0)
	camera := &Camera{}

	camera.SetState("auto mode")
	states = append(states, camera.State())
	camera.SetState("s mode")
	states = append(states, camera.State())
	camera.SetState("a mode")
	states = append(states, camera.State())
	camera.SetState("m mode")
	states = append(states, camera.State())

	camera.SetState(states[0])
}
```

```
Setting to "auto mode" state
Setting to "s mode" state
Setting to "a mode" state
Setting to "m mode" state
Setting to "auto mode" state
```
