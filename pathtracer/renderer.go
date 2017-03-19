package pathtracer

import (
	"image"
	"image/color"
	"math"
	"math/rand"
        "sync"
)

func Render(world Hitable, camera *Camera, nx int, ny int, ns int, cpus int) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, nx, ny))

        // create a waitgroup to hold the goroutines
        var wg sync.WaitGroup

        // for each core in use, add a gorountine to the waitgroup
        for i := 0; i < cpus; i++ {
                wg.Add(1)

                // concurrently sample each pixel
                go func(i int) {
                        defer wg.Done()

	                // main render loop
	                for y := i; y < ny; y += cpus {
		                for x := 0; x < nx; x++ {
			                rgb := Sample(world, camera, nx, ny, ns, x, y)
			                img.Set(x, ny-y-1, rgb)
                                }
                        }
		}(i)
	}

        // wait for all gorountines to finish before returning
        wg.Wait()
	return img
}

func Sample(world Hitable, camera *Camera, nx, ny, ns, x, y int) color.Color {
	col := Vector{0.0, 0.0, 0.0}

        // anti-aliasing and render loop
	for s := 0; s < ns; s++ {
                // generate the ray to be sent, sampling nearby pixels ns amount of times
		var u float64 = (float64(x) + rand.Float64()) / float64(nx)
		var v float64 = (float64(y) + rand.Float64()) / float64(ny)
		r := camera.GetRay(u, v)

                // trace the ray to generate the color for the pixel (x,y)
		col = col.Add(getColor(&r, world, 0))
	}
	col = col.ScalarDiv(float64(ns))

	// gamma correction
	col = Vector{math.Sqrt(col.X), math.Sqrt(col.Y), math.Sqrt(col.Z)}
	var ir uint8 = uint8(255.99 * col.X)
	var ig uint8 = uint8(255.99 * col.Y)
	var ib uint8 = uint8(255.99 * col.Z)

	return color.NRGBA{ir, ig, ib, 0xff}
}

func getColor(r *Ray, world Hitable, depth int) Vector {
	// check to see if the ray hits anything and store the information
	hit, rec := world.Hit(r, 1e-9, 1e9)

	// if there is a hit
	if hit {
		if depth < 50 {
			// scatter the ray
			success, scattered, tempColor := rec.Material.Scatter(*r, rec)

			// continue the loop if scattered
                        if success {
			        newColor := getColor(&scattered, world, depth+1)
			        return tempColor.Multiply(newColor)
                        }
                        return Vector{0.0, 0.0, 0.0}
		}
	}

	// no hit return the background color
	unitDirection := r.Direction().MakeUnitVector()
	var t float64 = 0.5 * (unitDirection.Y + 1.0)
	return Vector{1.0, 1.0, 1.0}.ScalarMulti(1.0 - t).Add(Vector{0.5, 0.7, 1.0}.ScalarMulti(t))
}

