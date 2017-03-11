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

        // multicore constraints with default value
        cpus := flag.Int("cpu", 4, "cpu core count")

	// camera constraints with default values
	fov := flag.Float64("fov", 20, "camera vertical field of view")
	aperture := flag.Float64("ap", 0.05, "camera aperture")
	loc := flag.Int("cam", 1, "camera location")
        ratio := float64(*nx)/float64(*ny)

	flag.Parse()

	// choose camara locations
	var lookfrom pt.Vector

        switch *loc {
        case 1:
                lookfrom = pt.Vector{13.0, 2.0, 4.0}
        case 2:
                lookfrom = pt.Vector{-13.0, 2.0, 4.0}
        default:
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
	world := *pt.CreateScene()
	camera := pt.CreateCamera(lookfrom, lookat, orientation, *fov, ratio, *aperture, distToFocus)

	image := pt.Render(&world, camera, *nx, *ny, *ns, *cpus)

	err = png.Encode(f, image)
	check(err, "Error writing to file: %v\n")
}

func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, s, e)
		os.Exit(1)
	}
}
