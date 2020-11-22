package field

import (
	"testing"

	"github.com/keftcha/floodfill/cell"
)

func equalSlice(s0, s1 []cell.Cell) bool {
	if len(s0) != len(s1) {
		return false
	}

	for idx, c0 := range s0 {
		c1 := s1[idx]
		if c0 != c1 {
			return false
		}
	}

	return true
}

func getField() Field {
	// The field look like this
	// # # # # # # # # # #
	// #           ¤     #
	// #             ¤ ¤ #
	// #   ¤ ¤ ¤         #
	// #   ¤   ¤         #
	// #   ¤   ¤   S     #
	// #   ¤ ¤           #
	// #                 #
	// # # # # # # # # # #
	// # → border
	// ¤ → cell that have Changed to true (wall)
	// S → start cell
	cells := [][]cell.Cell{
		[]cell.Cell{
			cell.New(0, 0, false),
			cell.New(1, 0, false),
			cell.New(2, 0, false),
			cell.New(3, 0, false),
			cell.New(4, 0, false),
			cell.New(5, 0, true),
			cell.New(6, 0, false),
			cell.New(7, 0, false),
		},
		[]cell.Cell{
			cell.New(0, 1, false),
			cell.New(1, 1, false),
			cell.New(2, 1, false),
			cell.New(3, 1, false),
			cell.New(4, 1, false),
			cell.New(5, 1, false),
			cell.New(6, 1, true),
			cell.New(7, 1, true),
		},
		[]cell.Cell{
			cell.New(0, 2, false),
			cell.New(1, 2, true),
			cell.New(2, 2, true),
			cell.New(3, 2, true),
			cell.New(4, 2, false),
			cell.New(5, 2, false),
			cell.New(6, 2, false),
			cell.New(7, 2, false),
		},
		[]cell.Cell{
			cell.New(0, 3, false),
			cell.New(1, 3, true),
			cell.New(2, 3, false),
			cell.New(3, 3, true),
			cell.New(4, 3, false),
			cell.New(5, 3, false),
			cell.New(6, 3, false),
			cell.New(7, 3, false),
		},
		[]cell.Cell{
			cell.New(0, 4, false),
			cell.New(1, 4, true),
			cell.New(2, 4, false),
			cell.New(3, 4, true),
			cell.New(4, 4, false),
			cell.New(5, 4, false),
			cell.New(6, 4, false),
			cell.New(7, 4, false),
		},
		[]cell.Cell{
			cell.New(0, 5, false),
			cell.New(1, 5, true),
			cell.New(2, 5, true),
			cell.New(3, 5, false),
			cell.New(4, 5, false),
			cell.New(5, 5, false),
			cell.New(6, 5, false),
			cell.New(7, 5, false),
		},
		[]cell.Cell{
			cell.New(0, 6, false),
			cell.New(1, 6, false),
			cell.New(2, 6, false),
			cell.New(3, 6, false),
			cell.New(4, 6, false),
			cell.New(5, 6, false),
			cell.New(6, 6, false),
			cell.New(7, 6, false),
		},
	}

	return New(cells, cells[4][5])
}

func TestCreateANewField(t *testing.T) {
	f := getField()

	if f.Width != 8 {
		t.Errorf("Field width isn't finded well, expected: 8, got: %d", f.Width)
	}
	if f.Height != 7 {
		t.Errorf("Field height isn't finded well, expected: 7, got: %d", f.Height)
	}

	expectedInitCell := cell.New(5, 4, false)
	expectedInitCell.ChangeNextStep = true
	if len(f.toChange) != 1 || f.toChange[0] != expectedInitCell {
		t.Errorf(
			"The initial cell to change is wrong, expected %v, got: %v",
			[]cell.Cell{expectedInitCell},
			f.toChange,
		)
	}

	if f.Filled {
		t.Error("The field is declared to filled")
	}
}

func TestDoOneStepShouldChangeTheInitialStateAndTheSliceOfToChangeCell(t *testing.T) {
	f := getField()

	changedCells := f.Step()

	expectedInitCell := cell.New(5, 4, true)
	expectedInitCell.ChangeNextStep = false
	if len(changedCells) != 1 || changedCells[0] != expectedInitCell {
		t.Errorf(
			"The returned changed cells is wrong, expected: %v, got: %v",
			expectedInitCell,
			changedCells[0],
		)
	}

	if !f.Cells[4][5].Changed {
		t.Error("The start cell isn't changed")
	}
	if f.Cells[4][5].ChangeNextStep {
		t.Error("The start cell will not change next step")
	}

	// Expected to change cell
	etc0 := cell.New(5, 3, false)
	etc0.ChangeNextStep = true
	etc1 := cell.New(6, 4, false)
	etc1.ChangeNextStep = true
	etc2 := cell.New(5, 5, false)
	etc2.ChangeNextStep = true
	etc3 := cell.New(4, 4, false)
	etc3.ChangeNextStep = true
	expectedToChange := []cell.Cell{etc0, etc1, etc2, etc3}

	if !equalSlice(expectedToChange, f.toChange) {
		t.Errorf(
			"The toChange cell for the next step is wrong, expected: %v, got: %v",
			expectedToChange,
			f.toChange,
		)
	}

	if f.Filled {
		t.Error("The field is indicated filled")
	}
}

func TestDoASecondStep(t *testing.T) {
	f := getField()
	f.Step()
	changedCellsGot := f.Step()

	// Expected changed cells
	cC0 := cell.New(5, 3, true)
	cC1 := cell.New(6, 4, true)
	cC2 := cell.New(5, 5, true)
	cC3 := cell.New(4, 4, true)
	changedCellsExpected := []cell.Cell{cC0, cC1, cC2, cC3}

	if !equalSlice(changedCellsExpected, changedCellsGot) {
		t.Errorf("The changed cells is wrong, expected: %v, got: %v", changedCellsExpected, changedCellsGot)
	}

	// Expected to change cells
	etc0 := cell.New(5, 2, false)
	etc0.ChangeNextStep = true
	etc1 := cell.New(6, 3, false)
	etc1.ChangeNextStep = true
	etc2 := cell.New(4, 3, false)
	etc2.ChangeNextStep = true

	etc3 := cell.New(7, 4, false)
	etc3.ChangeNextStep = true
	etc4 := cell.New(6, 5, false)
	etc4.ChangeNextStep = true

	etc5 := cell.New(5, 6, false)
	etc5.ChangeNextStep = true
	etc6 := cell.New(4, 5, false)
	etc6.ChangeNextStep = true

	// Don't add because there is a wall
	// etc7 := cell.New(3, 4, false),

	toChangeCellsExpected := []cell.Cell{etc0, etc1, etc2, etc3, etc4, etc5, etc6}
	toChangeCellsGot := f.toChange

	if !equalSlice(toChangeCellsExpected, f.toChange) {
		t.Errorf("The toChange cells is wrong, expected: %v, got: %v", toChangeCellsExpected, toChangeCellsGot)
	}

	if f.Filled {
		t.Error("The field is indicated filled")
	}
}

func TestDoEightSteps(t *testing.T) {
	f := getField()

	f.Step()
	f.Step()
	f.Step()
	f.Step()
	f.Step()
	f.Step()
	f.Step()
	changedCellsGot := f.Step()

	// Expected changed cells
	cC0 := cell.New(2, 0, true)
	cC1 := cell.New(1, 1, true)
	cC2 := cell.New(0, 6, true)
	changedCellsExpected := []cell.Cell{cC0, cC1, cC2}

	if !equalSlice(changedCellsExpected, changedCellsGot) {
		t.Errorf("The changed cells is wrong, expected: %v, got: %v", changedCellsExpected, changedCellsGot)
	}

	// Check field line by line

	// Line 0
	{
		c0 := cell.New(0, 0, false)
		c1 := cell.New(1, 0, false)
		c1.ChangeNextStep = true
		c2 := cell.New(2, 0, true)
		c3 := cell.New(3, 0, true)
		c4 := cell.New(4, 0, true)
		c5 := cell.New(5, 0, true)
		c6 := cell.New(6, 0, false)
		c7 := cell.New(7, 0, false)
		expectedFieldLine0 := []cell.Cell{c0, c1, c2, c3, c4, c5, c6, c7}

		if !equalSlice(expectedFieldLine0, f.Cells[0]) {
			t.Errorf("Line 0 of field is wrong, expected: %v, got: %v", expectedFieldLine0, f.Cells[0])
		}
	}
	// Line 1
	{
		c0 := cell.New(0, 1, false)
		c0.ChangeNextStep = true
		c1 := cell.New(1, 1, true)
		c2 := cell.New(2, 1, true)
		c3 := cell.New(3, 1, true)
		c4 := cell.New(4, 1, true)
		c5 := cell.New(5, 1, true)
		c6 := cell.New(6, 1, true)
		c7 := cell.New(7, 1, true)
		expectedFieldLine1 := []cell.Cell{c0, c1, c2, c3, c4, c5, c6, c7}

		if !equalSlice(expectedFieldLine1, f.Cells[1]) {
			t.Errorf("Line 1 of field is wrong, expected: %v, got: %v", expectedFieldLine1, f.Cells[1])
		}
	}
	// Line 2
	{
		c0 := cell.New(0, 2, false)
		c1 := cell.New(1, 2, true)
		c2 := cell.New(2, 2, true)
		c3 := cell.New(3, 2, true)
		c4 := cell.New(4, 2, true)
		c5 := cell.New(5, 2, true)
		c6 := cell.New(6, 2, true)
		c7 := cell.New(7, 2, true)
		expectedFieldLine2 := []cell.Cell{c0, c1, c2, c3, c4, c5, c6, c7}

		if !equalSlice(expectedFieldLine2, f.Cells[2]) {
			t.Errorf("Line 2 of field is wrong, expected: %v, got: %v", expectedFieldLine2, f.Cells[2])
		}
	}
	// Line 3
	{
		c0 := cell.New(0, 3, false)
		c1 := cell.New(1, 3, true)
		c2 := cell.New(2, 3, false)
		c3 := cell.New(3, 3, true)
		c4 := cell.New(4, 3, true)
		c5 := cell.New(5, 3, true)
		c6 := cell.New(6, 3, true)
		c7 := cell.New(7, 3, true)
		expectedFieldLine3 := []cell.Cell{c0, c1, c2, c3, c4, c5, c6, c7}

		if !equalSlice(expectedFieldLine3, f.Cells[3]) {
			t.Errorf("Line 3 of field is wrong, expected: %v, got: %v", expectedFieldLine3, f.Cells[3])
		}
	}
	// Line 4
	{
		c0 := cell.New(0, 4, false)
		c1 := cell.New(1, 4, true)
		c2 := cell.New(2, 4, false)
		c3 := cell.New(3, 4, true)
		c4 := cell.New(4, 4, true)
		c5 := cell.New(5, 4, true)
		c6 := cell.New(6, 4, true)
		c7 := cell.New(7, 4, true)
		expectedFieldLine4 := []cell.Cell{c0, c1, c2, c3, c4, c5, c6, c7}

		if !equalSlice(expectedFieldLine4, f.Cells[4]) {
			t.Errorf("Line 4 of field is wrong, expected: %v, got: %v", expectedFieldLine4, f.Cells[4])
		}
	}
	// Line 5
	{
		c0 := cell.New(0, 5, false)
		c0.ChangeNextStep = true
		c1 := cell.New(1, 5, true)
		c2 := cell.New(2, 5, true)
		c3 := cell.New(3, 5, true)
		c4 := cell.New(4, 5, true)
		c5 := cell.New(5, 5, true)
		c6 := cell.New(6, 5, true)
		c7 := cell.New(7, 5, true)
		expectedFieldLine5 := []cell.Cell{c0, c1, c2, c3, c4, c5, c6, c7}

		if !equalSlice(expectedFieldLine5, f.Cells[5]) {
			t.Errorf("Line 5 of field is wrong, expected: %v, got: %v", expectedFieldLine5, f.Cells[5])
		}
	}
	// Line 6
	{
		c0 := cell.New(0, 6, true)
		c1 := cell.New(1, 6, true)
		c2 := cell.New(2, 6, true)
		c3 := cell.New(3, 6, true)
		c4 := cell.New(4, 6, true)
		c5 := cell.New(5, 6, true)
		c6 := cell.New(6, 6, true)
		c7 := cell.New(7, 6, true)
		expectedFieldLine6 := []cell.Cell{c0, c1, c2, c3, c4, c5, c6, c7}

		if !equalSlice(expectedFieldLine6, f.Cells[6]) {
			t.Errorf("Line 6 of field is wrong, expected: %v, got: %v", expectedFieldLine6, f.Cells[6])
		}
	}

	if f.Filled {
		t.Error("The field is indicated filled")
	}
}

func TestDoElevenStepsAndTheFieldShouldBeFilled(t *testing.T) {
	f := getField()

	f.Step()
	f.Step()
	f.Step()
	f.Step()
	f.Step()
	f.Step()
	f.Step()
	f.Step()
	f.Step()
	f.Step()
	changedCellsGot := f.Step()

	// Expected changed cells
	changedCellsExpected := []cell.Cell{cell.New(0, 3, true)}

	if !equalSlice(changedCellsExpected, changedCellsGot) {
		t.Errorf("The changed cells is wrong, expected: %v, got: %v", changedCellsExpected, changedCellsGot)
	}

	// Check field line by line

	// Line 0
	{
		c0 := cell.New(0, 0, true)
		c1 := cell.New(1, 0, true)
		c2 := cell.New(2, 0, true)
		c3 := cell.New(3, 0, true)
		c4 := cell.New(4, 0, true)
		c5 := cell.New(5, 0, true)
		c6 := cell.New(6, 0, false)
		c7 := cell.New(7, 0, false)
		expectedFieldLine0 := []cell.Cell{c0, c1, c2, c3, c4, c5, c6, c7}

		if !equalSlice(expectedFieldLine0, f.Cells[0]) {
			t.Errorf("Line 0 of field is wrong, expected: %v, got: %v", expectedFieldLine0, f.Cells[0])
		}
	}
	// Line 1
	{
		c0 := cell.New(0, 1, true)
		c1 := cell.New(1, 1, true)
		c2 := cell.New(2, 1, true)
		c3 := cell.New(3, 1, true)
		c4 := cell.New(4, 1, true)
		c5 := cell.New(5, 1, true)
		c6 := cell.New(6, 1, true)
		c7 := cell.New(7, 1, true)
		expectedFieldLine1 := []cell.Cell{c0, c1, c2, c3, c4, c5, c6, c7}

		if !equalSlice(expectedFieldLine1, f.Cells[1]) {
			t.Errorf("Line 1 of field is wrong, expected: %v, got: %v", expectedFieldLine1, f.Cells[1])
		}
	}
	// Line 2
	{
		c0 := cell.New(0, 2, true)
		c1 := cell.New(1, 2, true)
		c2 := cell.New(2, 2, true)
		c3 := cell.New(3, 2, true)
		c4 := cell.New(4, 2, true)
		c5 := cell.New(5, 2, true)
		c6 := cell.New(6, 2, true)
		c7 := cell.New(7, 2, true)
		expectedFieldLine2 := []cell.Cell{c0, c1, c2, c3, c4, c5, c6, c7}

		if !equalSlice(expectedFieldLine2, f.Cells[2]) {
			t.Errorf("Line 2 of field is wrong, expected: %v, got: %v", expectedFieldLine2, f.Cells[2])
		}
	}
	// Line 3
	{
		c0 := cell.New(0, 3, true)
		c1 := cell.New(1, 3, true)
		c2 := cell.New(2, 3, false)
		c3 := cell.New(3, 3, true)
		c4 := cell.New(4, 3, true)
		c5 := cell.New(5, 3, true)
		c6 := cell.New(6, 3, true)
		c7 := cell.New(7, 3, true)
		expectedFieldLine3 := []cell.Cell{c0, c1, c2, c3, c4, c5, c6, c7}

		if !equalSlice(expectedFieldLine3, f.Cells[3]) {
			t.Errorf("Line 3 of field is wrong, expected: %v, got: %v", expectedFieldLine3, f.Cells[3])
		}
	}
	// Line 4
	{
		c0 := cell.New(0, 4, true)
		c1 := cell.New(1, 4, true)
		c2 := cell.New(2, 4, false)
		c3 := cell.New(3, 4, true)
		c4 := cell.New(4, 4, true)
		c5 := cell.New(5, 4, true)
		c6 := cell.New(6, 4, true)
		c7 := cell.New(7, 4, true)
		expectedFieldLine4 := []cell.Cell{c0, c1, c2, c3, c4, c5, c6, c7}

		if !equalSlice(expectedFieldLine4, f.Cells[4]) {
			t.Errorf("Line 4 of field is wrong, expected: %v, got: %v", expectedFieldLine4, f.Cells[4])
		}
	}
	// Line 5
	{
		c0 := cell.New(0, 5, true)
		c1 := cell.New(1, 5, true)
		c2 := cell.New(2, 5, true)
		c3 := cell.New(3, 5, true)
		c4 := cell.New(4, 5, true)
		c5 := cell.New(5, 5, true)
		c6 := cell.New(6, 5, true)
		c7 := cell.New(7, 5, true)
		expectedFieldLine5 := []cell.Cell{c0, c1, c2, c3, c4, c5, c6, c7}

		if !equalSlice(expectedFieldLine5, f.Cells[5]) {
			t.Errorf("Line 5 of field is wrong, expected: %v, got: %v", expectedFieldLine5, f.Cells[5])
		}
	}
	// Line 6
	{
		c0 := cell.New(0, 6, true)
		c1 := cell.New(1, 6, true)
		c2 := cell.New(2, 6, true)
		c3 := cell.New(3, 6, true)
		c4 := cell.New(4, 6, true)
		c5 := cell.New(5, 6, true)
		c6 := cell.New(6, 6, true)
		c7 := cell.New(7, 6, true)
		expectedFieldLine6 := []cell.Cell{c0, c1, c2, c3, c4, c5, c6, c7}

		if !equalSlice(expectedFieldLine6, f.Cells[6]) {
			t.Errorf("Line 6 of field is wrong, expected: %v, got: %v", expectedFieldLine6, f.Cells[6])
		}
	}

	if !f.Filled {
		t.Error("The field is indicated not filled")
	}
}

func TestContinueSteppingWhenTheFieldIsFilled(t *testing.T) {
	f := getField()

	f.Step()
	f.Step()
	f.Step()
	f.Step()
	f.Step()
	f.Step()
	f.Step()
	f.Step()
	f.Step()
	f.Step()
	f.Step()
	f.Step()
	changedCellsGot := f.Step()

	// Expected changed cells
	changedCellsExpected := make([]cell.Cell, 0)

	if !equalSlice(changedCellsExpected, changedCellsGot) {
		t.Errorf("The changed cells is wrong, expected: %v, got: %v", changedCellsExpected, changedCellsGot)
	}

	// Check field line by line

	// Line 0
	{
		c0 := cell.New(0, 0, true)
		c1 := cell.New(1, 0, true)
		c2 := cell.New(2, 0, true)
		c3 := cell.New(3, 0, true)
		c4 := cell.New(4, 0, true)
		c5 := cell.New(5, 0, true)
		c6 := cell.New(6, 0, false)
		c7 := cell.New(7, 0, false)
		expectedFieldLine0 := []cell.Cell{c0, c1, c2, c3, c4, c5, c6, c7}

		if !equalSlice(expectedFieldLine0, f.Cells[0]) {
			t.Errorf("Line 0 of field is wrong, expected: %v, got: %v", expectedFieldLine0, f.Cells[0])
		}
	}
	// Line 1
	{
		c0 := cell.New(0, 1, true)
		c1 := cell.New(1, 1, true)
		c2 := cell.New(2, 1, true)
		c3 := cell.New(3, 1, true)
		c4 := cell.New(4, 1, true)
		c5 := cell.New(5, 1, true)
		c6 := cell.New(6, 1, true)
		c7 := cell.New(7, 1, true)
		expectedFieldLine1 := []cell.Cell{c0, c1, c2, c3, c4, c5, c6, c7}

		if !equalSlice(expectedFieldLine1, f.Cells[1]) {
			t.Errorf("Line 1 of field is wrong, expected: %v, got: %v", expectedFieldLine1, f.Cells[1])
		}
	}
	// Line 2
	{
		c0 := cell.New(0, 2, true)
		c1 := cell.New(1, 2, true)
		c2 := cell.New(2, 2, true)
		c3 := cell.New(3, 2, true)
		c4 := cell.New(4, 2, true)
		c5 := cell.New(5, 2, true)
		c6 := cell.New(6, 2, true)
		c7 := cell.New(7, 2, true)
		expectedFieldLine2 := []cell.Cell{c0, c1, c2, c3, c4, c5, c6, c7}

		if !equalSlice(expectedFieldLine2, f.Cells[2]) {
			t.Errorf("Line 2 of field is wrong, expected: %v, got: %v", expectedFieldLine2, f.Cells[2])
		}
	}
	// Line 3
	{
		c0 := cell.New(0, 3, true)
		c1 := cell.New(1, 3, true)
		c2 := cell.New(2, 3, false)
		c3 := cell.New(3, 3, true)
		c4 := cell.New(4, 3, true)
		c5 := cell.New(5, 3, true)
		c6 := cell.New(6, 3, true)
		c7 := cell.New(7, 3, true)
		expectedFieldLine3 := []cell.Cell{c0, c1, c2, c3, c4, c5, c6, c7}

		if !equalSlice(expectedFieldLine3, f.Cells[3]) {
			t.Errorf("Line 3 of field is wrong, expected: %v, got: %v", expectedFieldLine3, f.Cells[3])
		}
	}
	// Line 4
	{
		c0 := cell.New(0, 4, true)
		c1 := cell.New(1, 4, true)
		c2 := cell.New(2, 4, false)
		c3 := cell.New(3, 4, true)
		c4 := cell.New(4, 4, true)
		c5 := cell.New(5, 4, true)
		c6 := cell.New(6, 4, true)
		c7 := cell.New(7, 4, true)
		expectedFieldLine4 := []cell.Cell{c0, c1, c2, c3, c4, c5, c6, c7}

		if !equalSlice(expectedFieldLine4, f.Cells[4]) {
			t.Errorf("Line 4 of field is wrong, expected: %v, got: %v", expectedFieldLine4, f.Cells[4])
		}
	}
	// Line 5
	{
		c0 := cell.New(0, 5, true)
		c1 := cell.New(1, 5, true)
		c2 := cell.New(2, 5, true)
		c3 := cell.New(3, 5, true)
		c4 := cell.New(4, 5, true)
		c5 := cell.New(5, 5, true)
		c6 := cell.New(6, 5, true)
		c7 := cell.New(7, 5, true)
		expectedFieldLine5 := []cell.Cell{c0, c1, c2, c3, c4, c5, c6, c7}

		if !equalSlice(expectedFieldLine5, f.Cells[5]) {
			t.Errorf("Line 5 of field is wrong, expected: %v, got: %v", expectedFieldLine5, f.Cells[5])
		}
	}
	// Line 6
	{
		c0 := cell.New(0, 6, true)
		c1 := cell.New(1, 6, true)
		c2 := cell.New(2, 6, true)
		c3 := cell.New(3, 6, true)
		c4 := cell.New(4, 6, true)
		c5 := cell.New(5, 6, true)
		c6 := cell.New(6, 6, true)
		c7 := cell.New(7, 6, true)
		expectedFieldLine6 := []cell.Cell{c0, c1, c2, c3, c4, c5, c6, c7}

		if !equalSlice(expectedFieldLine6, f.Cells[6]) {
			t.Errorf("Line 6 of field is wrong, expected: %v, got: %v", expectedFieldLine6, f.Cells[6])
		}
	}

	if !f.Filled {
		t.Error("The field is indicated not filled")
	}
}
