package pathtracer

import (
	"math"
)

type Sphere struct {
	Center Vector
	Radius float64
	Material
}

func NewSphere(center Vector, radius float64, material Material) *Sphere {
	return &Sphere{center, radius, material}
}

func (s *Sphere) Hit(r *Ray, tMin float64, tMax float64) (bool, HitRecord) {
	oc := r.Origin().Subtract(s.Center)
	var a float64 = r.Direction().Dot(r.Direction())
	var b float64 = oc.Dot(r.Direction())
	var c float64 = oc.Dot(oc) - s.Radius*s.Radius
	var d float64 = b*b - a*c

	rec := HitRecord{Material: s.Material}

	if d > 0.0 {
		var temp float64 = (-b - math.Sqrt(d)) / a
		if temp < tMax && temp > tMin {
			rec.T = temp
			rec.P = r.PointAtParameter(rec.T)
			rec.Normal = (rec.P.Subtract(s.Center)).ScalarDiv(s.Radius)
			return true, rec
		}
		temp = (-b + math.Sqrt(d)) / a
		if temp < tMax && temp > tMin {
			rec.T = temp
			rec.P = r.PointAtParameter(rec.T)
			rec.Normal = (rec.P.Subtract(s.Center)).ScalarDiv(s.Radius)
			return true, rec
		}
	}
	return false, rec
}
