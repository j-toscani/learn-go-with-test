package clock

import (
	"math"
	"time"
)

const (
	hoursInHalfClock   = 6
	hoursInClock       = hoursInHalfClock * 2
	minutesInHalfClock = 30
	minutesInClock     = minutesInHalfClock * 2
	secondsInHalfClock = 30
	secondsInClock     = secondsInHalfClock * 2
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(tm time.Time) Point {
	p := secondHandPoint(tm)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // scale
	p = Point{p.X, -p.Y}                                      // flip - account for origin difference in coordinate system and svg
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}         // translate
	return p
}

func secondsInRadians(rm time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / (float64(rm.Second()))))
}

func hoursInRadians(rm time.Time) float64 {
	return (minutesInRadians(rm) / hoursInClock) + (math.Pi / (hoursInHalfClock / float64(rm.Hour()%hoursInClock)))
}

func minutesInRadians(rm time.Time) float64 {
	return (secondsInRadians(rm) / minutesInClock) +
		(math.Pi / (minutesInClock / float64(rm.Minute())))
}

// x = sin(a)
// y = cos(a)
func secondHandPoint(rm time.Time) Point {
	radians := secondsInRadians(rm)
	return anglePoint(radians)
}

func minuteHandPoint(rm time.Time) Point {
	radians := minutesInRadians(rm)
	return anglePoint(radians)
}
func hourHandPoint(rm time.Time) Point {
	radians := hoursInRadians(rm)
	return anglePoint(radians)
}

func anglePoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
