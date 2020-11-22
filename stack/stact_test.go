package stack

import (
	"testing"

	"github.com/keftcha/floodfill/cell"
)

func TestCreatingAStackShouldBeEmpty(t *testing.T) {
	s := New()

	if len(s) != 0 {
		t.Error("The new stack isn't empty")
	}
}

func TestPopEmptyStackShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Pop an empty Stack did not panic")
		}
	}()

	s := New()
	s.Pop()
}

func TestPushACellShouldAddTheElementOnTheStack(t *testing.T) {
	s := New()
	s.Push(cell.New(0, 0))

	if len(s) != 1 {
		t.Error("The Cell isn't added on the Stack")
	}
}

func TestPushMultipleElementsShouldKeepThemInOrder(t *testing.T) {
	s := New()

	c0 := cell.New(0, 0)
	c1 := cell.New(1, 1)
	c2 := cell.New(2, 2)

	s.Push(c0)
	s.Push(c1)
	s.Push(c2)

	if len(s) != 3 {
		t.Error("All Cells haven't been added")
	}

	if s[0] != c0 || s[1] != c1 || s[2] != c2 {
		t.Error("Cells aren't added in order")
	}
}

func TestPopACellShouldReturnTheRemovedCellAndLetOtherCellsInStackInOrder(t *testing.T) {
	s := New()

	c0 := cell.New(0, 0)
	c1 := cell.New(1, 1)
	c2 := cell.New(2, 2)

	s.Push(c0)
	s.Push(c1)
	s.Push(c2)

	cellGot := s.Pop()

	if len(s) != 2 {
		t.Error("The Cell isn't popped")
	}

	if cellGot != c2 {
		t.Error("The popped cell isn't the last added one")
	}

	if s[0] != c0 || s[1] != c1 {
		t.Error("Cell order in Stack is changed")
	}
}

func TestPopMultipleCellsShouldReturnTheRemovedCellsAndLetOtherCellsInStackOrder(t *testing.T) {
	s := New()

	c0 := cell.New(0, 0)
	c1 := cell.New(1, 1)
	c2 := cell.New(2, 2)
	c3 := cell.New(3, 3)

	s.Push(c0)
	s.Push(c1)
	s.Push(c2)
	s.Push(c3)

	c3Got := s.Pop()
	c2Got := s.Pop()

	if len(s) != 2 {
		t.Error("Cells aren't popped")
	}

	if c3Got != c3 {
		t.Error("The popped cell isn't the last added one")
	}
	if c2Got != c2 {
		t.Error("The popped cell isn't the before lase added one")
	}

	if s[0] != c0 || s[1] != c1 {
		t.Error("Cell order in Stack is changed")
	}
}
