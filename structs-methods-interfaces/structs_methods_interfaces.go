package structsmethodsinterfaces

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimiter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}
