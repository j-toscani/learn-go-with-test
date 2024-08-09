package mocking

import (
	"fmt"
	"io"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(buffer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(buffer, i)
		sleeper.Sleep()
	}

	fmt.Fprint(buffer, finalWord)
}
