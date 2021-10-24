# Mediator Pattern

## Go

```go
package main

import (
	"fmt"
)

type Sales struct {
	Customers map[string]bool
}

func (s *Sales) Engage(customer string) {
	s.Customers[customer] = true
}

func (s *Sales) Communicate(customer string, message string) {
	if _, ok := s.Customers[customer]; ok {
		fmt.Printf("Sent %q to %q\n", message, customer)
	}
}

func NewSales() *Sales {
	return &Sales{
		Customers: make(map[string]bool),
	}
}

func main() {
	sales := NewSales()

	sales.Engage("Company A")
	sales.Engage("Company B")
	sales.Engage("Company C")

	sales.Communicate("Company A", "How is it going")
	sales.Communicate("Company C", "How is it going")
}
```

```
Sent "How is it going" to "Company A"
Sent "How is it going" to "Company C"
```
