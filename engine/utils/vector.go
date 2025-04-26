package utils

import "math"

type Vector struct {
	X float64
	Y float64
}

func (v Vector) GetEuclideanDistance(target *Vector) float64 {
	return math.Sqrt(v.GetEuclideanQuadraticDistance(target))
}

func (v Vector) GetEuclideanQuadraticDistance(target *Vector) float64 {
	dx := v.X - target.X
	dy := v.Y - target.Y

	return dx*dx + dy*dy
}

func (v *Vector) Add(o *Vector) *Vector {
	return &Vector{X: o.X + v.X, Y: o.Y + v.Y}
}

func (v *Vector) Divide(scalar float64) *Vector {
	return &Vector{X: v.X / scalar, Y: v.Y / scalar}
}

func (v *Vector) MultiplyByScalar(scalar float64) *Vector {
	return &Vector{X: v.X * scalar, Y: v.Y * scalar}
}
