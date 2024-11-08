package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		maxItteration := 3
		counter := NewCounter()

		for range maxItteration {
			counter.Inc()
		}

		assertCounter(t, counter, maxItteration)
	})
	t.Run("each incrementation increases caount by 1", func(t *testing.T) {
		counter := NewCounter()
		maxItterations := 5

		for itteration := range maxItterations {
			assertCounter(t, counter, itteration)
			counter.Inc()
		}

		assertCounter(t, counter, maxItterations)
	})
	t.Run("it runs save cocurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, counter *Counter, want int) {
	t.Helper()

	if counter.Value() != want {
		t.Errorf("got %d, want %d", counter.Value(), want)
	}
}
