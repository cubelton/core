package cube

type Color uint8

const (
	GREEN     Color = 0
	YELLOW    Color = 1
	RED       Color = 2
	BLUE      Color = 3
	WHITE     Color = 4
	ORANGE    Color = 5
	COLORLESS Color = 6
)

var colors = []Color{RED, GREEN, ORANGE, BLUE, YELLOW, WHITE}
