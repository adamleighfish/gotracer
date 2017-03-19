package pathtracer

import (
	"math"
)

type Sphere struct {
	Center   Vector
	Radius   float64
	Material Material
}

func NewSphere(center Vector, radius float64, material Material) *Sphere {
	return &Sphere{center, radius, material}
}

func (s *Sphere) Hit(r *Ray, tMin float64, tMax float64) (bool, HitRecord) {

	// determine the number of intersections with the sphere
	oc := r.Origin().Subtract(s.Center)
	var a float64 = r.Direction().Dot(r.Direction())
	var b float64 = oc.Dot(r.Direction())
	var c float64 = oc.Dot(oc) - s.Radius*s.Radius
	var d float64 = b*b - a*c

	rec := HitRecord{Material: s.Material}

	if d > tMin {
		var temp float64 = (-b - math.Sqrt(d)) / a

		// test both the positive and negative case
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

func (s *Sphere) BoundingBox(t0, t1 float64) (bool, AABB) {
        a := s.Center.Subtract(Vector{s.Radius, s.Radius, s.Radius})
        b := s.Center.Add(Vector{s.Radius, s.Radius, s.Radius})
        return true, AABB{a, b}
}
