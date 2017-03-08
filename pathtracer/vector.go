package pathtracer

import (
	"math"
	"math/rand"
)

type Vector struct {
	X, Y, Z float64
}

func (a Vector) MakeUnitVector() Vector {
	var k float64 = 1.0 / math.Sqrt(a.X*a.X+a.Y*a.Y+a.Z*a.Z)
	return Vector{a.X * k, a.Y * k, a.Z * k}
}

func (a Vector) SquaredLength() float64 {
	return a.X*a.X + a.Y*a.Y + a.Z*a.Z
}

func (a Vector) Length() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

func (a Vector) Add(b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (a Vector) Subtract(b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func (a Vector) Dot(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a Vector) Cross(b Vector) Vector {
	return Vector{a.Y*b.Z - a.Z*b.Y, a.Z*b.X - a.X*b.Z, a.X*b.Y - a.Y*b.X}
}

func (a Vector) ScalarMulti(t float64) Vector {
	return Vector{a.X * t, a.Y * t, a.Z * t}
}

func (a Vector) ScalarDiv(t float64) Vector {
	return Vector{a.X / t, a.Y / t, a.Z / t}
}

func (a Vector) Multiply(b Vector) Vector {
	return Vector{a.X * b.X, a.Y * b.Y, a.Z * b.Z}
}

func (a Vector) Reflect(b Vector) Vector {
	temp := 2 * a.Dot(b)
	return a.Subtract(b.ScalarMulti(temp))
}

func (a Vector) Refract(b Vector, n float64) (bool, Vector) {
	ua := a.MakeUnitVector()
	ub := b.MakeUnitVector()
	dt := ua.Dot(ub)
	d := 1.0 - (n * n * (1 - dt*dt))
	if d > 0.0 {
		x := ua.Subtract(b.ScalarMulti(dt)).ScalarMulti(n)
		y := b.ScalarMulti(math.Sqrt(d))
		return true, x.Subtract(y)
	}
	return false, Vector{}
}

func RandomInUnitSphere() Vector {
	for {
		r := Vector{rand.Float64(), rand.Float64(), rand.Float64()}
		p := r.ScalarMulti(2.0).Subtract(Vector{1.0, 1.0, 1.0})
		if p.SquaredLength() < 1.0 {
			return p
		}
	}
}
