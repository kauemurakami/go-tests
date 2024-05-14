package form

import (
	"math"
)

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}
func (c Circle) Area() float64 {
	return math.Pi * (c.Rad * c.Rad)
}

type Circle struct {
	Rad float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

type Form interface {
	Area() float64
}
