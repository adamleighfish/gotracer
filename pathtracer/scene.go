package pathtracer

import (
        "math/rand"
)


func CreateScene() *HitableList {
	var list HitableList = HitableList{}

	// create and add the floor
        floor := Sphere{Vector{0.0, -1000.0, 0}, 1000.0, &Lambertian{Vector{0.5, 0.5, 0.5}}}
        list.Add(&floor)

	// random sphere generation
	for a := -10; a < 10; a++ {
		for b := -10; b < 10; b++ {
			chooseMat := rand.Float64()
			center := Vector{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()}
			if center.Subtract(Vector{4.0, 0.2, 0.0}).Length() > 0.9 {
                                // choose diffuse material
                                if chooseMat < 0.7 {
                                        list.Add(&Sphere{center, 0.2, &Lambertian{Vector{
                                                X: rand.Float64() * rand.Float64(),
                                                Y: rand.Float64() * rand.Float64(),
                                                Z: rand.Float64() * rand.Float64()}}})
				// choose moving diffuse material
                                } else if chooseMat < 0.8 {
					list.Add(&MovingSphere{center, center.Add(Vector{0.0, 0.5 * rand.Float64(), 0.0}), 0.0, 1.0, 0.2,
                                                &Lambertian{Vector{
                                                X: rand.Float64() * rand.Float64(),
						Y: rand.Float64() * rand.Float64(),
						Z: rand.Float64() * rand.Float64()}}})
				// choose metal material
				} else if chooseMat < 0.95 {
					list.Add(&Sphere{center, 0.2, &Metal{Vector{
                                                X: (1.0 + rand.Float64()) * 0.5,
						Y: (1.0 + rand.Float64()) * 0.5,
						Z: (1.0 + rand.Float64()) * 0.5}, rand.Float64() * 0.5}})
				// choose glass material
				} else {
					list.Add(&Sphere{center, 0.2, &Dielectric{1.5}})
				}
			}
		}
	}

	// create three large sphere, one of each material
	sphere1 := Sphere{Vector{0.0, 1.0, 0.0}, 1.0, &Dielectric{1.5}}
	sphere2 := Sphere{Vector{-4.0, 1.0, 0.0}, 1.0, &Lambertian{Vector{0.4, 0.2, 0.1}}}
	sphere3 := Sphere{Vector{4.0, 1.0, 0.0}, 1.0, &Metal{Vector{0.7, 0.6, 0.5}, 0.0}}

	list.Add(&sphere1)
	list.Add(&sphere2)
	list.Add(&sphere3)

	return &list
}
