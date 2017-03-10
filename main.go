package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
        "flag"

	pt "github.com/epicdangerfish/gotracer/pathtracer"
)

func init() {
	// generate a new rand seed each run
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// max color value
	const c = 255.99

        // image constraints with default values
        nx := flag.Int("w", 800, "image width")
        ny := flag.Int("h", 600, "image height")
        ns := flag.Int("s", 100, "sample rate")

        // camera constraints with default values
        fov := flag.Float64("fov", 20, "camera vertical field of view")
        aperture := flag.Float64("ap", 0.05, "camera aperture")
        loc := flag.Int("cam", 1, "camera location")

        flag.Parse()

        // choose camara locations
        var lookfrom pt.Vector
        if *loc == 1 {
                lookfrom = pt.Vector{13.0, 2.0, 4.0}
        } else if *loc == 2 {
                lookfrom = pt.Vector{-13.0, 2.0, 4.0}
        } else {
                lookfrom = pt.Vector{2.0, 2.0, 2.0}
        }

	// camera contraints
	lookat := pt.Vector{0.0, 0.0, 0.0}
	orientation := pt.Vector{0.0, 1.0, 0.0}
	distToFocus := (lookfrom.Subtract(lookat)).Length()

	f, err := os.Create("out.ppm")
	check(err, "Error opening file: %v\n")

	defer f.Close()

	_, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", *nx, *ny)
	check(err, "Error writing to file: %v\n")

	// create the scene to render
	world := *createScene()
	camera := pt.CreateCamera(lookfrom, lookat, orientation, *fov, float64(*nx)/float64(*ny), *aperture, distToFocus)

	// main render loop
	for j := *ny - 1; j >= 0; j-- {
		for i := 0; i < *nx; i++ {

			// anti aliasing loop
			col := pt.Vector{0.0, 0.0, 0.0}
			for s := 0; s < *ns; s++ {
				var u float64 = (float64(i) + rand.Float64()) / float64(*nx)
				var v float64 = (float64(j) + rand.Float64()) / float64(*ny)
				r := camera.GetRay(u, v)
				col = col.Add(color(&r, &world, 0))
			}

			// gamma correction
			col = col.ScalarDiv(float64(*ns))
			col = pt.Vector{math.Sqrt(col.X), math.Sqrt(col.Y), math.Sqrt(col.Z)}
			var ir int = int(c * col.X)
			var ig int = int(c * col.Y)
			var ib int = int(c * col.Z)

			_, err = fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)

			check(err, "Error writing to file: %v\n")
		}
	}
}

func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, s, e)
		os.Exit(1)
	}
}

func color(r *pt.Ray, world pt.Hitable, depth int) pt.Vector {

	// check to see if the ray hits anything and store the information
	hit, rec := world.Hit(r, 1e-9, 1e9)

	// if there is a hit
	if hit {
		if depth < 50 {
			// scatter the ray
			success, scattered := rec.Scatter(*r, rec)

			// continue the loop if scattered
			if success {
				newColor := color(&scattered, world, depth+1)
				return rec.Material.Color().Multiply(newColor)
			}
			return rec.Material.Color()
		}
	}

	// no hit return the background color
	unitDirection := r.Direction().MakeUnitVector()
	var t float64 = 0.5 * (unitDirection.Y + 1.0)
	return pt.Vector{1.0, 1.0, 1.0}.ScalarMulti(1.0 - t).Add(pt.Vector{0.5, 0.7, 1.0}.ScalarMulti(t))
}

func createScene() *pt.HitableList {
	var list pt.HitableList = pt.HitableList{}

	// create and add the floor
	floor := pt.Sphere{Center: pt.Vector{0.0, -2000.0, 0.0},
		Radius:   2000.0,
		Material: &pt.Lambertian{pt.Vector{0.9, 0.9, 0.9}}}
	list.Add(&floor)

	// random sphere generation
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := pt.Vector{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()}
			if center.Subtract(pt.Vector{4.0, 0.2, 0.0}).Length() > 0.9 {
				// choose diffuse material
				if chooseMat < 0.8 {
					list.Add(&pt.Sphere{center, 0.2, &pt.Lambertian{pt.Vector{X: rand.Float64() * rand.Float64(),
						Y: rand.Float64() * rand.Float64(),
						Z: rand.Float64() * rand.Float64()}}})
					// choose metal material
				} else if chooseMat < 0.95 {
					list.Add(&pt.Sphere{center, 0.2, &pt.Metal{pt.Vector{X: (1.0 + rand.Float64()) * 0.5,
						Y: (1.0 + rand.Float64()) * 0.5,
						Z: (1.0 + rand.Float64()) * 0.5}, rand.Float64() * 0.5}})
					// choose glass material
				} else {
					list.Add(&pt.Sphere{center, 0.2, &pt.Dielectric{1.5}})
				}
			}
		}
	}

	// create three large sphere, one of each material
	sphere1 := pt.Sphere{pt.Vector{0.0, 1.0, 0.0}, 1.0, &pt.Dielectric{1.5}}
	sphere2 := pt.Sphere{pt.Vector{-4.0, 1.0, 0.0}, 1.0, &pt.Lambertian{pt.Vector{0.4, 0.2, 0.1}}}
	sphere3 := pt.Sphere{pt.Vector{4.0, 1.0, 0.0}, 1.0, &pt.Metal{pt.Vector{0.7, 0.6, 0.5}, 0.0}}

	list.Add(&sphere1)
	list.Add(&sphere2)
	list.Add(&sphere3)

	return &list
}
