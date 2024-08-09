package main

import (
	mocking "main/mocking"
	"os"
)

func main() {
	mocking.Countdown(os.Stdout)
}
