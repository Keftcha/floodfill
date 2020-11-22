package stack

import (
	"github.com/keftcha/floodfill/cell"
)

// Stack of Cell
type Stack []cell.Cell

// New Stack of Cell
func New() Stack {
	return make(Stack, 0)
}

// Push a Cell on the Stack
func (s *Stack) Push(c cell.Cell) {
	*s = append((*s), c)
}

// Pop a Cell of the Stack
func (s *Stack) Pop() cell.Cell {
	lastIdx := len(*s) - 1
	c := (*s)[lastIdx]
	*s = (*s)[:lastIdx]
	return c
}
