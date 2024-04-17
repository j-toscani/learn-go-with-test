package iteration

import (
	"fmt"
	"testing"
)

func TestRepat(t *testing.T) {
	t.Run("adds the value 5 times per default", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("adds the value n times if provided", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}

func ExampleRepeat() {
	repeatDefault := Repeat("a", 0)
	fmt.Println(repeatDefault)
	// Output: aaaaa
}
