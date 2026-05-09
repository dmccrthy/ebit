package utils

import "math"

type Vector struct {
	X float64
	Y float64
}

func (v *Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector) Normalize() {
	magnitude := v.Magnitude()
	v.X /= magnitude
	v.Y /= magnitude
}

func (v *Vector) Scale(scale float64) {
	v.Normalize()
	v.X *= scale
	v.Y *= scale
}
