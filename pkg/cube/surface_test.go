package cube

import "testing"

func TestInitialSurfaceBits(t *testing.T) {
	t.Log("Surface binary conversion")
	vectors := []struct {
		color       Color
		binaryValue string
	}{
		{GREEN, "000000000000000000000000000"},
		{YELLOW, "001001001001001001001001001"},
		{RED, "010010010010010010010010010"},
		{BLUE, "011011011011011011011011011"},
		{WHITE, "100100100100100100100100100"},
		{ORANGE, "101101101101101101101101101"},
		{COLORLESS, "110110110110110110110110110"},
		{COLORLESS, "110110110110110110110110110"},
	}
	for _, vector := range vectors {
		surface := NewSurface(vector.color)
		bits := surface.Bits()

		if bits != vector.binaryValue {
			t.Error("Surface binary value should be "+vector.binaryValue+", but ", bits)
		} else {
			t.Log("binary value is " + bits)
		}
	}
}

func TestCustomSurfaceBits(t *testing.T) {
	vectors := []struct {
		cells       [9]*Cell
		binaryValue string
	}{
		{
			[9]*Cell{
				NewCell(GREEN), NewCell(ORANGE), NewCell(GREEN),
				NewCell(GREEN), NewCell(GREEN), NewCell(GREEN),
				NewCell(GREEN), NewCell(GREEN), NewCell(GREEN),
			},
			"000101000000000000000000000",
		},
		{
			[9]*Cell{
				NewCell(GREEN), NewCell(RED), NewCell(GREEN),
				NewCell(RED), NewCell(GREEN), NewCell(RED),
				NewCell(GREEN), NewCell(RED), NewCell(GREEN),
			},
			"000010000010000010000010000",
		},
		{
			[9]*Cell{
				NewCell(GREEN), NewCell(ORANGE), NewCell(GREEN),
				NewCell(ORANGE), NewCell(GREEN), NewCell(ORANGE),
				NewCell(GREEN), NewCell(ORANGE), NewCell(GREEN),
			},
			"000101000101000101000101000",
		},
		{
			[9]*Cell{
				NewCell(GREEN), NewCell(YELLOW), NewCell(RED),
				NewCell(BLUE), NewCell(WHITE), NewCell(ORANGE),
				NewCell(COLORLESS), NewCell(GREEN), NewCell(GREEN),
			},
			"000001010011100101110000000",
		},
		{
			[9]*Cell{
				NewCell(GREEN), NewCell(RED), NewCell(WHITE),
				NewCell(GREEN), NewCell(RED), NewCell(WHITE),
				NewCell(GREEN), NewCell(RED), NewCell(WHITE),
			},
			"000010100000010100000010100",
		},
	}

	for _, vector := range vectors {
		surface := new(Surface)
		surface.cells = vector.cells
		bits := surface.Bits()
		if bits != vector.binaryValue {
			t.Error("Surface binary value should be "+vector.binaryValue+", but ", bits)
		} else {
			t.Log("binary value is " + bits)
		}
	}
}

func TestInitialSurfaceHex(t *testing.T) {
	t.Log("Surface binary conversion")
	vectors := []struct {
		color    Color
		hexValue string
	}{
		{GREEN, "00000000"},
		{YELLOW, "01249249"},
		{RED, "02492492"},
		{BLUE, "036db6db"},
		{WHITE, "04924924"},
		{ORANGE, "05b6db6d"},
		{COLORLESS, "06db6db6"},
		{COLORLESS, "06db6db6"},
	}
	for _, vector := range vectors {
		surface := NewSurface(vector.color)
		hex := surface.Hex()

		if hex != vector.hexValue {
			t.Error("Surface hex value should be "+vector.hexValue+", but ", hex)
		} else {
			t.Log("hex value is " + hex)
		}
	}
}

func TestCustomSurfaceHex(t *testing.T) {
	vectors := []struct {
		cells    [9]*Cell
		hexValue string
	}{
		{
			[9]*Cell{
				NewCell(GREEN), NewCell(RED), NewCell(GREEN),
				NewCell(RED), NewCell(GREEN), NewCell(RED),
				NewCell(GREEN), NewCell(RED), NewCell(GREEN),
			},
			"00410410",
		},
		{
			[9]*Cell{
				NewCell(GREEN), NewCell(ORANGE), NewCell(GREEN),
				NewCell(ORANGE), NewCell(GREEN), NewCell(ORANGE),
				NewCell(GREEN), NewCell(ORANGE), NewCell(GREEN),
			},
			"00a28a28",
		},
		{
			[9]*Cell{
				NewCell(GREEN), NewCell(YELLOW), NewCell(RED),
				NewCell(BLUE), NewCell(WHITE), NewCell(ORANGE),
				NewCell(COLORLESS), NewCell(GREEN), NewCell(GREEN),
			},
			"0029cb80",
		},
		{
			[9]*Cell{
				NewCell(GREEN), NewCell(RED), NewCell(WHITE),
				NewCell(GREEN), NewCell(RED), NewCell(WHITE),
				NewCell(GREEN), NewCell(RED), NewCell(WHITE),
			},
			"00502814",
		},
	}

	for _, vector := range vectors {
		surface := new(Surface)
		surface.cells = vector.cells
		hex := surface.Hex()
		if hex != vector.hexValue {
			t.Error("Surface hex value should be "+vector.hexValue+", but ", hex)
		} else {
			t.Log("hex value is " + hex)
		}
	}
}

func TestCustomSurfaceRotate(t *testing.T) {
	vectors := []struct {
		cells       [9]*Cell
		binaryValue string
		hexValue    string
	}{
		{
			[9]*Cell{
				NewCell(GREEN), NewCell(GREEN), NewCell(GREEN),
				NewCell(GREEN), NewCell(GREEN), NewCell(GREEN),
				NewCell(GREEN), NewCell(GREEN), NewCell(COLORLESS),
			},
			"000000000000000000000000110",
			"00000006",
		},
		{
			[9]*Cell{
				NewCell(GREEN), NewCell(ORANGE), NewCell(GREEN),
				NewCell(ORANGE), NewCell(GREEN), NewCell(ORANGE),
				NewCell(GREEN), NewCell(ORANGE), NewCell(COLORLESS),
			},
			"000101000101000101000101110",
			"00a28a2e",
		},
		{
			[9]*Cell{
				NewCell(ORANGE), NewCell(GREEN), NewCell(GREEN),
				NewCell(GREEN), NewCell(GREEN), NewCell(GREEN),
				NewCell(GREEN), NewCell(GREEN), NewCell(COLORLESS),
			},
			"000000101000000000000000110",
			"00140006",
		},
		{
			[9]*Cell{
				NewCell(GREEN), NewCell(ORANGE), NewCell(GREEN),
				NewCell(GREEN), NewCell(GREEN), NewCell(GREEN),
				NewCell(GREEN), NewCell(GREEN), NewCell(COLORLESS),
			},
			"000000000101000000000000110",
			"00028006",
		},
		{
			[9]*Cell{
				NewCell(GREEN), NewCell(GREEN), NewCell(GREEN),
				NewCell(GREEN), NewCell(GREEN), NewCell(GREEN),
				NewCell(ORANGE), NewCell(GREEN), NewCell(COLORLESS),
			},
			"101000000000000000000000110",
			"05000006",
		},
		{
			[9]*Cell{
				NewCell(GREEN), NewCell(GREEN), NewCell(GREEN),
				NewCell(GREEN), NewCell(GREEN), NewCell(GREEN),
				NewCell(GREEN), NewCell(ORANGE), NewCell(COLORLESS),
			},
			"000101000000000000000000110",
			"00a00006",
		},
	}
	for _, vector := range vectors {
		surface := new(Surface)
		surface.cells = vector.cells
		surface.Rotate()
		bits := surface.Bits()

		if bits != vector.binaryValue {
			t.Error("Surface binary value should be "+vector.binaryValue+", but ", bits)
		} else {
			t.Log("binary value is " + bits)
		}
		hex := surface.Hex()
		if hex != vector.hexValue {
			t.Error("Surface hex value should be "+vector.hexValue+", but ", hex)
		} else {
			t.Log("hex value is " + hex)
		}
	}
}
