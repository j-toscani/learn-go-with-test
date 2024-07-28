package structsmethodsinterfaces

import "testing"

func TestPerimiter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimiter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	t.Run("rectangle", func(t *testing.T) {

		rectangle := Rectangle{5.0, 10.0}
		got := rectangle.Area()
		want := 50.0

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})

	t.Run("circle", func(t *testing.T) {

		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})
}
