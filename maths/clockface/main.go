package main

import (
	"main/maths"
	"os"
	"time"
)

func main() {
	t := time.Now()
	maths.SVGWriter(os.Stdout, t)
}
