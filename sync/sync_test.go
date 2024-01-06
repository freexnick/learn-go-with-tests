package sync

import (
	"sync"
	"testing"
)

func NewCounter() *Counter {
	return &Counter{}
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	value := got.Value()

	if value != want {
		t.Errorf("got %d, want %d", value, want)
	}
}

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()

		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		counter := NewCounter()
		var wg sync.WaitGroup
		wantedCount := 1000

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
