package main

import (
	"flag"
	"fmt"
	"image/png"
	"math/rand"
	"os"
	"time"

	pt "github.com/epicdangerfish/gotracer/pathtracer"
)

func init() {
	// generate a new rand seed each run
	rand.Seed(time.Now().UnixNano())
}

func main() {
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
		lookfrom = pt.Vector{0.0, 1.0, 10.0}
	}

	// camera contraints
	lookat := pt.Vector{0.0, 0.0, 0.0}
	orientation := pt.Vector{0.0, 1.0, 0.0}
	distToFocus := (lookfrom.Subtract(lookat)).Length()

	f, err := os.Create("out.png")
	check(err, "Error opening file: %v\n")

	defer f.Close()

	// create the scene to render
	world := *createScene()
	camera := pt.CreateCamera(lookfrom, lookat, orientation, *fov, float64(*nx)/float64(*ny), *aperture, distToFocus)

	image := pt.Render(&world, camera, *nx, *ny, *ns, 0)

	err = png.Encode(f, image)
	check(err, "Error writing to file: %v\n")
}

func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, s, e)
		os.Exit(1)
	}
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
