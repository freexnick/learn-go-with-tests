package main

import (
	"bytes"
	"slices"
	"testing"
	"time"
)

const write = "write"
const sleep = "sleep"

type SpyCountdownOperations struct {
	Calls []string
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountDown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		var buffer bytes.Buffer
		var spyCountdownOperations SpyCountdownOperations
		Countdown(&buffer, &spyCountdownOperations)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		var spySleepPrinter SpyCountdownOperations

		Countdown(&spySleepPrinter, &spySleepPrinter)

		want := []string{
			write, sleep, write, sleep, write, sleep, write,
		}

		if !slices.Equal(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	var spyTime SpyTime
	sleepTime := 5 * time.Second

	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()
}
