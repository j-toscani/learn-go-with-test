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
	rectangle := Rectangle{5.0, 10.0}
	got := Area(rectangle)
	want := 50.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}
