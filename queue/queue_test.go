package queue

import (
	"testing"

	"github.com/keftcha/floodfill/cell"
)

func TestCreatingAQueueShouldBeEmpty(t *testing.T) {
	q := New()

	if len(q) != 0 {
		t.Error("The new queue isn't empty")
	}
}

func TestDequeueEmptyQueueShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Dequeue an empty Queue did not panic")
		}
	}()

	q := New()
	q.Dequeue()
}

func TestEnqueueACellShouldAddTheElementInTheQueue(t *testing.T) {
	q := New()
	q.Enqueue(cell.New(0, 0, false))

	if len(q) != 1 {
		t.Error("The Cell isn't added to the Queue")
	}
}

func TestEnqueueMultipleElementsShouldKeepThemInOrder(t *testing.T) {
	q := New()

	c0 := cell.New(0, 0, false)
	c1 := cell.New(1, 1, false)
	c2 := cell.New(2, 2, false)

	q.Enqueue(c0)
	q.Enqueue(c1)
	q.Enqueue(c2)

	if len(q) != 3 {
		t.Error("All Cells haven't been added")
	}

	if q[0] != c0 || q[1] != c1 || q[2] != c2 {
		t.Error("Cells aren't added in order")
	}
}

func TestDequeueACellShouldReturnTheRemovedCellAndLetOtherCellsInQueueInOrder(t *testing.T) {
	q := New()

	c0 := cell.New(0, 0, false)
	c1 := cell.New(1, 1, false)
	c2 := cell.New(2, 2, false)

	q.Enqueue(c0)
	q.Enqueue(c1)
	q.Enqueue(c2)

	cellGot := q.Dequeue()

	if len(q) != 2 {
		t.Error("The Cell isn't dequeued")
	}

	if cellGot != c0 {
		t.Error("The dequeued cell isn't the first added cell")
	}

	if q[0] != c1 || q[1] != c2 {
		t.Error("Cells order in Queue is changed")
	}
}

func TestDequeueMultipleCellsShouldReturnTheRemovedCellAndLetOtherCellsInQueueInOrder(t *testing.T) {
	q := New()

	c0 := cell.New(0, 0, false)
	c1 := cell.New(1, 1, false)
	c2 := cell.New(2, 2, false)
	c3 := cell.New(3, 3, false)

	q.Enqueue(c0)
	q.Enqueue(c1)
	q.Enqueue(c2)
	q.Enqueue(c3)

	c0Got := q.Dequeue()
	c1Got := q.Dequeue()

	if len(q) != 2 {
		t.Error("Cells aren't dequeued")
	}

	if c0Got != c0 {
		t.Error("The dequeued cell isn't the first added cell")
	}
	if c1Got != c1 {
		t.Error("The dequeued cell isn't the second added cell")
	}

	if q[0] != c2 || q[1] != c3 {
		t.Error("Cells order in Queue is changed")
	}
}
