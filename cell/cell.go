package cell

// Cell struct
type Cell struct {
	X              int  // X coordinate of the Cell
	Y              int  // Y coordinate of the Cell
	Changed        bool // Have the Cell already be changed ? If set to true it will stop the propagaition
	ChangeNextStep bool // The cell will change at next step
}

// New return a Cell
func New(x, y int, changed bool) Cell {
	return Cell{x, y, changed, false}
}
