package cube

import (
	"testing"
)

func TestCellBits(t *testing.T) {
	t.Log("Cell binary conversion")
	vectors := []struct {
		color       Color
		binaryValue string
	}{
		{GREEN, "000"},
		{YELLOW, "001"},
		{RED, "010"},
		{BLUE, "011"},
		{WHITE, "100"},
		{ORANGE, "101"},
		{COLORLESS, "110"},
	}
	for _, vector := range vectors {
		cell := NewCell(vector.color)
		bits := cell.Bits()

		if bits != vector.binaryValue {
			t.Error("Cell binary value should be "+vector.binaryValue+", but ", bits)
		} else {
			t.Log("binary value is " + bits)
		}
	}
}
