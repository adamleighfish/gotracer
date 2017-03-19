package pathtracer

import (
	"math"
	"math/rand"
)

type Dielectric struct {
	ReflectIndex float64
}

func (d *Dielectric) Scatter(r Ray, rec HitRecord) (bool, Ray, Vector) {
	var outwardNormal, refracted Vector
	var snell, reflectProb, cosine float64
	var success bool

	if r.Direction().Dot(rec.Normal) > 0.0 {
		outwardNormal = rec.Normal.ScalarMulti(-1.0)
		snell = d.ReflectIndex

		a := r.Direction().Dot(rec.Normal) * d.ReflectIndex
		b := r.Direction().Length()

		cosine = a / b
	} else {
		outwardNormal = rec.Normal
		snell = 1.0 / d.ReflectIndex

		a := r.Direction().Dot(rec.Normal) * d.ReflectIndex
		b := r.Direction().Length()

		cosine = -a / b
	}

	success, refracted = r.Direction().Refract(outwardNormal, snell)

	if success {
		reflectProb = d.Schlick(cosine)
	} else {
		reflectProb = 1.0
	}

	if rand.Float64() < reflectProb {
		reflected := r.Direction().Reflect(rec.Normal)
		return true, Ray{rec.P, reflected, r.Time()}, Vector{1.0, 1.0, 1.0}
	}
	return true, Ray{rec.P, refracted, r.Time()}, Vector{1.0, 1.0, 1.0}
}

// Schlick's approx. of the specular reflection coeff.
func (d *Dielectric) Schlick(cosine float64) float64 {
	var r0 float64 = (1.0 - d.ReflectIndex) / (1.0 + d.ReflectIndex)
	r0 = r0 * r0
	return r0 + (1.0-r0)*math.Pow(1.0-cosine, 5.0)
}
