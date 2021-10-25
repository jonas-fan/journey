# Interpreter Pattern

## Go

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func push(stack *[]int, value int) {
	(*stack) = append(*stack, value)
}

func pop(stack *[]int) int {
	index := len(*stack) - 1
	top := (*stack)[index]

	(*stack) = (*stack)[:index]

	return top
}

type MathExpr func(values *[]int)

func MathAddExpr() MathExpr {
	return func(values *[]int) {
		rhs := pop(values)
		lhs := pop(values)

		push(values, lhs+rhs)
	}
}

func MathSubExpr() MathExpr {
	return func(values *[]int) {
		rhs := pop(values)
		lhs := pop(values)

		push(values, lhs-rhs)
	}
}

func MathValueExpr(value int) MathExpr {
	return func(values *[]int) {
		push(values, value)
	}
}

type MathInterpreter struct{}

func (i *MathInterpreter) Interpret(express string) int {
	var expr MathExpr
	tokens := strings.Split(express, " ")
	values := make([]int, 0, len(tokens)/2+1)

	for _, token := range tokens {
		switch token {
		case "+":
			expr = MathAddExpr()
		case "-":
			expr = MathSubExpr()
		default:
			value, _ := strconv.Atoi(token)
			expr = MathValueExpr(value)
		}

		expr(&values)
	}

	return pop(&values)
}

func main() {
	interpreter := &MathInterpreter{}
	express := "5 1 7 + -"

	fmt.Printf("%q => %d", express, interpreter.Interpret(express))
}
```

```
"5 1 7 + -" => -3
```
