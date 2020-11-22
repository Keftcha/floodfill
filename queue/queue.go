package queue

import (
	"github.com/keftcha/floodfill/cell"
)

// Queue of Cell
type Queue []cell.Cell

// New Queue of Cell
func New() Queue {
	return make(Queue, 0)
}

// Enqueue a Cell in the Queue
func (q *Queue) Enqueue(c cell.Cell) {
	*q = append((*q), c)
}

// Dequeue a Cell of the Queue
func (q *Queue) Dequeue() cell.Cell {
	c := (*q)[0]
	*q = (*q)[1:]
	return c
}
