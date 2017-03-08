package pathtracer

type Metal struct {
	Albedo Vector
	Fuzz   float64
}

func (m *Metal) Scatter(r Ray, rec HitRecord) (bool, Ray) {
	reflected := r.Direction().Reflect(rec.Normal)
	bouncedRay := Ray{rec.P, reflected.Add(RandomInUnitSphere().ScalarMulti(m.Fuzz))}
	bounced := bouncedRay.Direction().Dot(rec.Normal) > 0
	return bounced, bouncedRay
}

func (m *Metal) Color() Vector {
	return m.Albedo
}
