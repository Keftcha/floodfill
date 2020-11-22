package field

import (
	"github.com/keftcha/floodfill/cell"
)

// Field of cell
type Field struct {
	Width    int           // width of the field
	Height   int           // height of the field
	Cells    [][]cell.Cell // cells in the field
	toChange []cell.Cell   // cells to change if the fiels
	Filled   bool          // Is the area filled
}

// New return a Field
func New(cells [][]cell.Cell, startCell cell.Cell) Field {
	cells[startCell.Y][startCell.X].ChangeNextStep = true
	return Field{
		Width:    len(cells[0]),
		Height:   len(cells),
		Cells:    cells,
		toChange: []cell.Cell{cells[startCell.Y][startCell.X]},
		Filled:   false,
	}
}

// Step change cells state and return changed one
func (f *Field) Step() []cell.Cell {
	nextGen := make([]cell.Cell, 0)
	changedCells := make([]cell.Cell, len(f.toChange))
	for idx, c := range f.toChange {
		// Cell above
		if 0 < c.Y {
			above := f.Cells[c.Y-1][c.X]
			if !above.Changed && !above.ChangeNextStep {
				f.Cells[above.Y][above.X].ChangeNextStep = true
				nextGen = append(nextGen, f.Cells[above.Y][above.X])
			}
		}
		// Cell on right
		if c.X < f.Width-1 {
			right := f.Cells[c.Y][c.X+1]
			if !right.Changed && !right.ChangeNextStep {
				f.Cells[right.Y][right.X].ChangeNextStep = true
				nextGen = append(nextGen, f.Cells[right.Y][right.X])
			}
		}
		// Cell bellow
		if c.Y < f.Height-1 {
			bellow := f.Cells[c.Y+1][c.X]
			if !bellow.Changed && !bellow.ChangeNextStep {
				f.Cells[bellow.Y][bellow.X].ChangeNextStep = true
				nextGen = append(nextGen, f.Cells[bellow.Y][bellow.X])
			}
		}
		// Cell on left
		if 0 < c.X {
			left := f.Cells[c.Y][c.X-1]
			if !left.Changed && !left.ChangeNextStep {
				f.Cells[left.Y][left.X].ChangeNextStep = true
				nextGen = append(nextGen, f.Cells[left.Y][left.X])
			}
		}
		f.Cells[c.Y][c.X].Changed = true
		f.Cells[c.Y][c.X].ChangeNextStep = false
		changedCells[idx] = f.Cells[c.Y][c.X]
	}

	if len(nextGen) == 0 {
		f.Filled = true
	}

	f.toChange = nextGen
	return changedCells
}
