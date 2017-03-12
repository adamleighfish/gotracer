package pathtracer

type Lambertian struct {
	Albedo Vector
}

func (l *Lambertian) Scatter(r Ray, rec HitRecord) (bool, Ray) {
	target := rec.Normal.Add(RandomInUnitSphere())
	return true, Ray{rec.P, target, r.Time()}
}

func (l *Lambertian) Color() Vector {
	return l.Albedo
}
