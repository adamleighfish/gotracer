package pathtracer

import (
	"math"
)

type Plane struct {
	Point    Vector
	Normal   Vector
	Material Material
}

func NewPlane(p Vector, n Vector, m Material) *Plane {
	n = n.MakeUnitVector()
	return &Plane{p, n, m}
}

func (p *Plane) Hit(r *Ray, tMin float64, tMax float64) (bool, HitRecord) {
	l := r.Direction().Dot(p.Normal)

	rec := HitRecord{Material: p.Material}

	if math.Abs(l) < tMin {
		return false, rec
	}

	d := (p.Point.Subtract(r.Origin())).Dot(p.Normal)
	d = d / l

	if math.Abs(d) < tMin {
		return false, rec
	}

	return true, HitRecord{d, r.PointAtParameter(d), p.Normal, p.Material}
}
