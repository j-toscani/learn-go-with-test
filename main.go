package main

import (
	mocking "main/mocking"
	"os"
	"time"
)

type DefaultSleeper struct {}

func (d *DefaultSleeper) Sleep(){
	time.Sleep(1*time.Second)
}

func main() {
	sleeper := &DefaultSleeper{}
	mocking.Countdown(os.Stdout, sleeper)
}
