# N minutes in Go

- https://golang.org/doc/
- https://play.golang.org/

```go
// Declaration of a package
package main

// An inline comment
/*
A block comment
*/

// Importing packages
import (
	"fmt"
)

// Global variables
var app = "N minutes in Go"
var (
	name     = "Jonas"
	fullname = name + " Fan"
)

// The entrypoint of an application
func main() {
	fmt.Println("Hello,", fullname)
}
```

