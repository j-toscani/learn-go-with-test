package mocking

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep(duration time.Duration)
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep(duration time.Duration) {
	s.Calls++
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(buffer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(buffer, i)
		sleeper.Sleep(1 * time.Second)
	}

	fmt.Fprint(buffer, finalWord)
}
