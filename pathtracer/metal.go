package pathtracer

type Metal struct {
	Albedo Texture
	Fuzz   float64
}

func (m *Metal) Scatter(r Ray, rec HitRecord) (bool, Ray, Vector) {

	// find the reflected vector
	reflected := r.Direction().Reflect(rec.Normal)

	// account for the fuzziness of the material (max fuzziness at 1.0)
        if m.Fuzz > 1.0 { m.Fuzz = 1.0 }
	bouncedRay := Ray{rec.P, reflected.Add(RandomInUnitSphere().ScalarMulti(m.Fuzz)), r.Time()}

	// make sure the bounce succeeded
	bounced := bouncedRay.Direction().Dot(rec.Normal) > 0
	return bounced, bouncedRay, m.Albedo.Value(0.0, 0.0, rec.P)
}
