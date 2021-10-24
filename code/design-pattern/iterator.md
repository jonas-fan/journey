# Iterator Pattern

## Go

```go
package main

import (
	"fmt"
)

type Photo struct {
	Name string
}

func (p *Photo) String() string {
	return p.Name
}

type PhotoIterator struct {
	photos []*Photo
	index  int
}

func (i *PhotoIterator) HasNext() bool {
	return i.index < len(i.photos)
}

func (i *PhotoIterator) Next() *Photo {
	photo := i.photos[i.index]

	i.index++

	return photo
}

type Camera struct {
	photos []*Photo
}

func (c *Camera) Take(name string) {
	c.photos = append(c.photos, &Photo{Name: name})
}

func (c *Camera) Iterator() *PhotoIterator {
	return &PhotoIterator{
		photos: c.photos,
	}
}

func main() {
	camera := &Camera{}
	camera.Take("photo-001.jpeg")
	camera.Take("photo-002.jpeg")
	camera.Take("photo-003.jpeg")

	iter := camera.Iterator()

	for iter.HasNext() {
		fmt.Printf("Photo: %s\n", iter.Next())
	}
}
```

```
Photo: photo-001.jpeg
Photo: photo-002.jpeg
Photo: photo-003.jpeg
```
