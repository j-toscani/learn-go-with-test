package main

import (
	mocking "main/mocking"
	"os"
	"time"
)

func main() {
	sleeper := &mocking.ConfigurableSleeper{Duration: 5 * time.Second, Sleeper: time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
