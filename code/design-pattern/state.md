# State Pattern

## Go

```go
package main

import (
	"fmt"
)

type StateFn func() StateFn

func Day() StateFn {
	fmt.Println("Day")
	return Night
}

func Night() StateFn {
	fmt.Println("Night")
	return Day
}

type Timer struct {
	state StateFn
}

func (t *Timer) Lapse() {
	t.state = t.state()
}

func NewTimer() *Timer {
	return &Timer{
		state: Day,
	}
}

func main() {
	timer := NewTimer()

	timer.Lapse()
	timer.Lapse()
	timer.Lapse()
	timer.Lapse()
}
```

```
Day
Night
Day
Night
```
