# Composite Pattern

## Go

```go
package main

import (
	"fmt"
)

type Filter interface {
	Do()
}

type LightFilter struct{}

func (f *LightFilter) Do() {
	fmt.Println("Light filter")
}

type LomoFilter struct{}

func (f *LomoFilter) Do() {
	fmt.Println("Lomo filter")
}

type RetroFilter struct{}

func (f *RetroFilter) Do() {
	fmt.Println("Retro filter")
}

type FilterSet struct {
	Filters []Filter
}

func (f *FilterSet) Add(filter Filter) {
	f.Filters = append(f.Filters, filter)
}

func (f *FilterSet) Do() {
	for _, filter := range f.Filters {
		filter.Do()
	}
}

func main() {
	fmt.Println("Filtering ...")
	filter1 := &FilterSet{}
	filter1.Add(&LightFilter{})
	filter1.Add(&LomoFilter{})
	filter1.Do()

	fmt.Println("\nFiltering ...")
	filter2 := &FilterSet{}
	filter2.Add(filter1)
	filter2.Add(&RetroFilter{})
	filter2.Do()
}
```

```
Filtering ...
Light filter
Lomo filter

Filtering ...
Light filter
Lomo filter
Retro filter
```
