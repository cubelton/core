package cube

import (
	"fmt"
	"strconv"
)

type Side uint8

const size = 9

type Surface struct {
	cells [size]*Cell
}

func NewSurface(color Color) *Surface {
	surface := new(Surface)
	for _, position := range positions {
		surface.cells[position] = NewCell(color)
	}
	return surface
}

func (s *Surface) Bits() string {
	result := ""
	for _, position := range positions {
		result += s.cells[position].Bits()
	}
	return result
}

func (s *Surface) Hex() string {
	u, _ := strconv.ParseUint(s.Bits(), 2, 64)
	return fmt.Sprintf("%08x", u)
}

func (s *Surface) Rotate() {
	//rotate main
	northCell := s.cells[North]
	s.cells[North] = s.cells[West]
	s.cells[West] = s.cells[South]
	s.cells[South] = s.cells[East]
	s.cells[East] = northCell
	//rotate angels
	northEastCell := s.cells[NorthEast]
	s.cells[NorthEast] = s.cells[NorthWest]
	s.cells[NorthWest] = s.cells[SouthWest]
	s.cells[SouthWest] = s.cells[SouthEast]
	s.cells[SouthEast] = northEastCell
}

func (s *Surface) RotateBack() {
	//rotate main
	northCell := s.cells[North]
	s.cells[North] = s.cells[East]
	s.cells[East] = s.cells[South]
	s.cells[South] = s.cells[West]
	s.cells[West] = northCell
	//rotate angels
	northEastCell := s.cells[NorthEast]
	s.cells[NorthEast] = s.cells[SouthEast]
	s.cells[SouthEast] = s.cells[SouthWest]
	s.cells[SouthWest] = s.cells[NorthWest]
	s.cells[NorthWest] = northEastCell
}

func (s *Surface) DoubleRotate() {
	//rotate main
	northCell := s.cells[North]
	s.cells[North] = s.cells[South]
	s.cells[South] = northCell
	eastCell := s.cells[East]
	s.cells[East] = s.cells[West]
	s.cells[West] = eastCell
	//rotate angels
	northEastCell := s.cells[NorthEast]
	s.cells[NorthEast] = s.cells[SouthWest]
	s.cells[SouthWest] = northEastCell
	southEastCell := s.cells[SouthEast]
	s.cells[SouthEast] = s.cells[NorthWest]
	s.cells[NorthWest] = southEastCell
}
