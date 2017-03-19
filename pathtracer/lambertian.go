package pathtracer

type Lambertian struct {
	Albedo Texture
}

func (l *Lambertian) Scatter(r Ray, rec HitRecord) (bool, Ray, Vector) {
	target := rec.Normal.Add(RandomInUnitSphere())
        scattered := Ray{rec.P, target, r.Time()}
        color := l.Albedo.Value(0.0, 0.0, rec.P)
	return true, scattered, color
}
