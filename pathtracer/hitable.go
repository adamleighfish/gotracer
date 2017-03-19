package pathtracer

type HitRecord struct {
	T               float64
	P, Normal       Vector
	Material        Material
}

type Hitable interface {
	Hit(r *Ray, tMin float64, tMax float64) (bool, HitRecord)
        BoundingBox(t0, t1 float64) (bool, AABB)
}
