package main

import (
	mocking "main/mocking"
	"os"
	"time"
)

type DefaultSleeper struct {}

func (d *DefaultSleeper) Sleep(duration time.Duration){
	time.Sleep(duration)
}

func main() {
	sleeper := &DefaultSleeper{}
	mocking.Countdown(os.Stdout, sleeper)
}
