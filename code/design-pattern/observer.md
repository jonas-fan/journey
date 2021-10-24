# Observer Pattern

## Go

```go
package main

import (
	"fmt"
)

type Notifier interface {
	Notify(message string)
}

type News struct {
	subscribers map[string]Notifier
}

func (n *News) Register(name string, notifier Notifier) {
	n.subscribers[name] = notifier
}

func (n *News) Notify(message string) {
	for _, each := range n.subscribers {
		each.Notify(message)
	}
}

func NewNews() *News {
	return &News{
		subscribers: make(map[string]Notifier),
	}
}

type Person struct {
	Name string
}

func (p *Person) Notify(message string) {
	fmt.Printf("%s: got a message %q\n", p.Name, message)
}

func main() {
	person1 := &Person{Name: "Jonas"}
	person2 := &Person{Name: "Jane"}

	news := NewNews()

	news.Register(person1.Name, person1)
	news.Register(person2.Name, person2)
	news.Notify("hello world")
}
```

```
Jane: got a message "hello world"
Jonas: got a message "hello world"
```
