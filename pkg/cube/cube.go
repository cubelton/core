package cube

type Cube struct {
	surfaces [6]*Surface
}

func NewCube() *Cube {
	cube := new(Cube)
	for _, color := range colors {
		cube.surfaces[color] = NewSurface(color)
	}
	return cube
}
func (c *Cube) GetSurfaces() [6]*Surface {
	return c.surfaces
}

func (c *Cube) Bits() string {
	result := ""
	for _, surface := range c.surfaces {
		result += surface.Bits()
	}
	return result
}
func (c *Cube) StateCode() string {
	result := ""
	for _, surface := range c.surfaces {
		result += surface.Hex()
	}
	return result
}

func (c *Cube) rotate(color Color) {
	c.surfaces[color].Rotate()
}
