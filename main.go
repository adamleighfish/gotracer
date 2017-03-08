package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	pt "github.com/epicdangerfish/gotracer/pathtracer"
)

func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, s, e)
		os.Exit(1)
	}
}

func color(r *pt.Ray, world pt.Hitable, depth int) pt.Vector {
	hit, rec := world.Hit(r, 0.001, math.MaxFloat64)

	if hit {
		if depth < 50 {
			bounced, bouncedRay := rec.Scatter(*r, rec)
			if bounced {
				newColor := color(&bouncedRay, world, depth+1)
				return rec.Material.Color().Multiply(newColor)
			}
			return pt.Vector{0.0, 0.0, 0.0}
		}
	}
	unitDirection := r.Direction().MakeUnitVector()
	var t float64 = 0.5 * (unitDirection.Y + 1.0)
	return pt.Vector{1.0, 1.0, 1.0}.ScalarMulti(1.0 - t).Add(pt.Vector{0.5, 0.7, 1.0}.ScalarMulti(t))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// width and heigh of image
	var nx int = 800
	var ny int = 600
	var ns int = 100

	// camera contraints
	lookfrom := pt.Vector{13.0, 2.0, 4.0}
	lookat := pt.Vector{0.0, 0.0, 0.0}
	distToFocus := (lookfrom.Subtract(lookat)).Length()
	aperture := 0.05
	fov := 20.0

	const max = 255.99

	f, err := os.Create("out.ppm")
	check(err, "Error opening file: %v\n")

	defer f.Close()

	_, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", nx, ny)
	check(err, "Error writing to file: %v\n")

	world := CreateScene()

	// HitableList{[]Hitable{&middle, &floor, &right, &left, &inner}}

	camera := pt.CreateCamera(lookfrom, lookat, pt.Vector{0.0, 1.0, 0.0}, fov, float64(nx)/float64(ny), aperture, distToFocus)

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {

			col := pt.Vector{0.0, 0.0, 0.0}
			for s := 0; s < ns; s++ {
				var u float64 = (float64(i) + rand.Float64()) / float64(nx)
				var v float64 = (float64(j) + rand.Float64()) / float64(ny)
				r := camera.GetRay(u, v)
				col = col.Add(color(&r, &world, 0))
			}

			col = col.ScalarDiv(float64(ns))
			col = pt.Vector{math.Sqrt(col.X), math.Sqrt(col.Y), math.Sqrt(col.Z)}
			var ir int = int(max * col.X)
			var ig int = int(max * col.Y)
			var ib int = int(max * col.Z)

			_, err = fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)

			check(err, "Error writing to file: %v\n")
		}
	}
}

func CreateScene() pt.HitableList {
	var list pt.HitableList = pt.HitableList{}
	floor := pt.Sphere{pt.Vector{0.0, -1000.0, 0.0}, 1000.0, &pt.Lambertian{pt.Vector{0.5, 0.5, 0.5}}}
	list.Add(&floor)

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := pt.Vector{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()}
			if center.Subtract(pt.Vector{4.0, 0.2, 0.0}).Length() > 0.9 {
				if chooseMat < 0.8 {
					list.Add(&pt.Sphere{center, 0.2, &pt.Lambertian{pt.Vector{rand.Float64() * rand.Float64(), rand.Float64() * rand.Float64(), rand.Float64() * rand.Float64()}}})
				} else if chooseMat < 0.95 {
					list.Add(&pt.Sphere{center, 0.2, &pt.Metal{pt.Vector{1.0 + rand.Float64(), 1.0 + rand.Float64(), 1.0 + rand.Float64()}.ScalarMulti(0.5), 0.5 * rand.Float64()}})
				} else {
					list.Add(&pt.Sphere{center, 0.2, &pt.Dielectric{1.5}})
				}
			}
		}
	}

	sphere1 := pt.Sphere{pt.Vector{0.0, 1.0, 0.0}, 1.0, &pt.Dielectric{1.5}}
	sphere2 := pt.Sphere{pt.Vector{-4.0, 1.0, 0.0}, 1.0, &pt.Lambertian{pt.Vector{0.4, 0.2, 0.1}}}
	sphere3 := pt.Sphere{pt.Vector{4.0, 1.0, 0.0}, 1.0, &pt.Metal{pt.Vector{0.7, 0.6, 0.5}, 0.0}}

	list.Add(&sphere1)
	list.Add(&sphere2)
	list.Add(&sphere3)

	return list
}
