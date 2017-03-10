package pathtracer

type Metal struct {
	Albedo Vector
	Fuzz   float64
}

func (m *Metal) Scatter(r Ray, rec HitRecord) (bool, Ray) {

	// find the reflected vector
	reflected := r.Direction().Reflect(rec.Normal)

	// account for the fuzziness of the material
	bouncedRay := Ray{rec.P, reflected.Add(RandomInUnitSphere().ScalarMulti(m.Fuzz))}

	// make sure the bounce succeeded
	bounced := bouncedRay.Direction().Dot(rec.Normal) > 0
	return bounced, bouncedRay
}

func (m *Metal) Color() Vector {
	return m.Albedo
}
