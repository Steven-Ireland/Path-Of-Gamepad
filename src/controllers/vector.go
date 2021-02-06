package controllers

import "math"

var Zero2 = Vector2{}

type Vector2 struct{ X, Y float64 }

// XY returns both components
func (a Vector2) XY() (x, y float64) { return a.X, a.Y }

// Add adds two vectors and returns the result
func (a Vector2) Add(b Vector2) Vector2                 { return Vector2{a.X + b.X, a.Y + b.Y} }
func (a Vector2) AddScale(b Vector2, s float64) Vector2 { return Vector2{a.X + b.X*s, a.Y + b.Y*s} }

// Sub subtracts two vectors and returns the result
func (a Vector2) Sub(b Vector2) Vector2 { return Vector2{a.X - b.X, a.Y - b.Y} }

// Dot calculates the dot product
func (a Vector2) Dot(b Vector2) float64 { return a.X*b.X + a.Y*b.Y }

// Scale scales each component and returns the result
func (a Vector2) Scale(s float64) Vector2 { return Vector2{a.X * s, a.Y * s} }

// Length returns the length of the vector
func (a Vector2) Length() float64 { return math.Sqrt(a.X*a.X + a.Y*a.Y) }

// Length2 returns the squared length of the vector
func (a Vector2) Length2() float64 { return a.X*a.X + a.Y*a.Y }

// Distance returns the distance to vector b
func (a Vector2) Distance(b Vector2) float64 {
	dx, dy := a.X-b.X, a.Y-b.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// Distance2 returns the squared distance to vector b
func (a Vector2) Distance2(b Vector2) float64 {
	dx, dy := a.X-b.X, a.Y-b.Y
	return dx*dx + dy*dy
}

func (a Vector2) Normalize() Vector2 {
	m := a.Length()
	if m < 1 {
		m = 1
	}
	return Vector2{a.X / m, a.Y / m}
}

func (a Vector2) Negate() Vector2 { return Vector2{-a.X, -a.Y} }

// Cross product of a and b
func (a Vector2) Cross(b Vector2) float64 { return a.X*b.Y - a.Y*b.X }

func (a Vector2) NearZero() bool { return a.Length2() < 0.0001 }

func (a Vector2) Rotate(angle float64) Vector2 {
	cs, sn := math.Cos(angle), math.Sin(angle)
	return Vector2{a.X*cs - a.Y*sn, a.X*sn + a.Y*cs}
}

func (a Vector2) Angle() float64 { return math.Atan2(a.Y, a.X) }

func (a Vector2) Rotate90() Vector2  { return Vector2{-a.Y, a.X} }
func (a Vector2) Rotate90c() Vector2 { return Vector2{a.Y, -a.X} }
func (a Vector2) Rotate180() Vector2 { return Vector2{-a.X, -a.Y} }
