package cube

import (
	"fmt"
)

type Position uint8

const (
	North     Position = 0
	NorthEast Position = 1
	East      Position = 2
	SouthEast Position = 3
	South     Position = 4
	SouthWest Position = 5
	West      Position = 6
	NorthWest Position = 7
	Center    Position = 8
)

var positions = []Position{North, NorthEast, East, SouthEast, South, SouthWest, West, NorthWest, Center}

type Cell struct {
	color Color
}

func NewCell(color Color) *Cell {
	cell := new(Cell)
	cell.color = color
	return cell
}

func (cell *Cell) Bits() string {
	return fmt.Sprintf("%03b", cell.color)
}
