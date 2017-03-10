package pathtracer

import (
	"image"
	"image/color"
	"math"
	"math/rand"
)

func getColor(r *Ray, world Hitable, depth int) Vector {

	// check to see if the ray hits anything and store the information
	hit, rec := world.Hit(r, 1e-9, 1e9)

	// if there is a hit
	if hit {
		if depth < 50 {
			// scatter the ray
			success, scattered := rec.Scatter(*r, rec)

			// continue the loop if scattered
			if success {
				newColor := getColor(&scattered, world, depth+1)
				return rec.Material.Color().Multiply(newColor)
			}
			return Vector{1.0, 1.0, 1.0}
		}
	}

	// no hit return the background color
	unitDirection := r.Direction().MakeUnitVector()
	var t float64 = 0.5 * (unitDirection.Y + 1.0)
	return Vector{1.0, 1.0, 1.0}.ScalarMulti(1.0 - t).Add(Vector{0.5, 0.7, 1.0}.ScalarMulti(t))
}

func Render(world Hitable, camera *Camera, nx int, ny int, ns int, cpus int) image.Image {

	const c = 255.99

	img := image.NewNRGBA(image.Rect(0, 0, nx, ny))

	// main render loop
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {

			// anti aliasing loop
			col := Vector{0.0, 0.0, 0.0}
			for s := 0; s < ns; s++ {
				var u float64 = (float64(i) + rand.Float64()) / float64(nx)
				var v float64 = (float64(j) + rand.Float64()) / float64(ny)
				r := camera.GetRay(u, v)
				col = col.Add(getColor(&r, world, 0))
			}

			// gamma correction
			col = col.ScalarDiv(float64(ns))
			col = Vector{math.Sqrt(col.X), math.Sqrt(col.Y), math.Sqrt(col.Z)}
			var ir uint8 = uint8(c * col.X)
			var ig uint8 = uint8(c * col.Y)
			var ib uint8 = uint8(c * col.Z)

			rgb := color.NRGBA{ir, ig, ib, 0xff}

			img.Set(i, ny-j-1, rgb)
		}
	}

	return img
}
