package cell

// Cell struct
type Cell struct {
	X       int  // X coordinate of the Cell
	Y       int  // Y coordinate of the Cell
	Changed bool // Have the Cell already be changed ?
}

// New return a Cell
func New(x, y int) Cell {
	return Cell{x, y, false}
}
