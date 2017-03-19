package pathtracer

import (
        "math"
)

type MovingSphere struct {
        Center1, Center2        Vector
        Time1, Time2            float64
        Radius                  float64
        Material                Material
}

func NewMovingSphere(c1, c2 Vector, t1, t2, r float64, mat Material) *MovingSphere {
        return &MovingSphere{c1, c2, t1, t2, r, mat}
}

func (s *MovingSphere) Hit(r *Ray, tMin, tMax float64) (bool, HitRecord) {
        oc := r.Origin().Subtract(s.Center(r.Time()))
        a := r.Direction().Dot(r.Direction())
        b := oc.Dot(r.Direction())
        c := oc.Dot(oc) - s.Radius * s.Radius
        d := b * b - a * c

        rec := HitRecord{Material: s.Material}
        if d > 0.0 {
                temp := (-b - math.Sqrt(d)) / a
                if temp > tMin && temp < tMax {
                        rec.T = temp
                        rec.P = r.PointAtParameter(temp)
                        rec.Normal = (rec.P.Subtract(s.Center(r.Time()))).ScalarDiv(s.Radius)
                        return true, rec
                }
                temp = (-b + math.Sqrt(d)) / a
                if temp > tMin && temp < tMax {
                        rec.T = temp
                        rec.P = r.PointAtParameter(temp)
                        rec.Normal = (rec.P.Subtract(s.Center(r.Time()))).ScalarDiv(s.Radius)
                        return true, rec
                }
        }
        return false, rec
}

func (s *MovingSphere) BoundingBox(t0, t1 float64) (bool, AABB) {
        s1 := Sphere{s.Center1, s.Radius, s.Material}
        s2 := Sphere{s.Center2, s.Radius, s.Material}
        _, b1 := s1.BoundingBox(t0, t1)
        _, b2 := s2.BoundingBox(t0, t1)
        return true, *surroundingBox(&b1, &b2)
}

func (s *MovingSphere) Center(time float64) Vector {
        dt := (time - s.Time1) / (s.Time2 - s.Time1)
        dc := s.Center2.Subtract(s.Center1)
        return s.Center1.Add(dc.ScalarMulti(dt))
}
