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
	greeting(app)
	greeting(fullname)

	variables()
	choices(2)
	// loops()
}

func greeting(name string) {
	fmt.Println("Hello,", name)
}

func variables() {
	var b1 bool                  // false
	var b2 bool = true           // true
	var b3 = b1 && b2            // false
	b4 := !b3                    // true

	var n1 int = 3               // 3
	n2 := 5                      // 5 (int)
	var n3 uint64 = 1<<64 - 1    // 18446744073709551615

	var p1 *int                  // nil
	p2 := &n2                    // address of n2 (*int)

	r := 'a'                     // 0x61 (rune or int32) 
	str := "cake"                // cake (string)

	s1 := []int{2, 4, 6}         // [2, 4, 6] (slice, len=3, cap=3)
	s2 := append(s1, 8)          // [2, 4, 6, 8] (slice, len=4, cap=6)
	s3 := s2[1:3]                // [4, 6] (slice, len=2, cap=5)

	a1 := [3]int{2, 4, 6}        // [2, 4, 6] (array, len=3, cap=3)
	a2 := [3]int{8}              // [8, 0, 0] (array, len=3, cap=3)

	var m1 map[string]string             // nil map
	var m2 = make(map[string]string)     // map
	m2["Name"] = "Jonas Fan"
	m2["Role"] = "Software Developer"
	m3 := map[string]string{             // map
		"Name": m2["Name"],
		"Role": m2["Role"],
	}

	var i1 interface{}                   // nil (interface)
	var i2 interface{} = m1              // interface to m1

	fmt.Println(b1, b2, b3, b4)
	fmt.Println(n1, n2, n3)
	fmt.Println(p1, p2, *p2)
	fmt.Println(r, str)
	fmt.Println(s1, s2, s3)
	fmt.Println(a1, a2)
	fmt.Println(m1, m2, m3)
	fmt.Println(i1, i2)
}

func choices(n int) {
	// if-else statement
	if n < 0 {
		fmt.Println("n is negative")
	} else {
		fmt.Println("n is positive")
	}

	if n < 0 {
		fmt.Println("n < 0", n)
	} else if n < 10 {
		fmt.Println("0 <= n < 10")
	} else if n < 100 {
		fmt.Println("10 <= n < 100")
	} else {
		fmt.Println("100 <= n")
	}

	// switch statment
	switch (n) {
	case 0:
		fmt.Println("n is 0")
	case 1:
		fmt.Print("n is 1 or ")
		fallthrough
	case 2:
		fmt.Println("n is 2")
	case 3, 4, 5:
		fmt.Println("n is in [3, 4, 5]")
	}

	switch {
	case n < 0:
		fmt.Println("n < 0")
	case n < 10:
		fmt.Println("0 <= n < 10")
	case n < 100:
		fmt.Println("10 <= n < 100")
	default:
		fmt.Println("100 <= n")
	}
}


```
