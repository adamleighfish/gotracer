package pathtracer

type Material interface {
	Scatter(r Ray, rec HitRecord) (bool, Ray, Vector)
}
