package pathtracer

import (
	"math"
	"math/rand"
)

type Camera struct {
	LowerLeft  Vector
	Horizontal Vector
	Vertical   Vector
	Origin     Vector
	U, V, W    Vector
	LensRadius float64
}

// create a camera with the given constraints and return a pointer to it
func CreateCamera(lookfrom Vector, lookat Vector, vup Vector, vfov float64, aspect float64, aperture float64, focusDist float64) *Camera {
	var u, v, w Vector

	// variables used for framing picture
	theta := vfov * math.Pi / 180.0
	halfHeight := math.Tan(theta / 2.0)
	halfWidth := aspect * halfHeight

	lensRadius := aperture / 2.0

	// determine the orientation of the camera
	w = lookfrom.Subtract(lookat).MakeUnitVector()
	u = vup.Cross(w).MakeUnitVector()
	v = w.Cross(u)

	// camera framing
	origin := lookfrom
	lowerLeft := origin.Subtract(u.ScalarMulti(halfWidth * focusDist)).Subtract(v.ScalarMulti(halfHeight * focusDist)).Subtract(w.ScalarMulti(focusDist))
	horizontal := u.ScalarMulti(2.0 * halfWidth * focusDist)
	vertical := v.ScalarMulti(2.0 * halfHeight * focusDist)

	return &Camera{lowerLeft, horizontal, vertical, origin, u, v, w, lensRadius}
}

//
func (c *Camera) GetRay(u, v float64) Ray {
	rd := RandomInUnitDisk().ScalarMulti(c.LensRadius)
	offset := c.U.ScalarMulti(rd.X).Add(c.V.ScalarMulti(rd.Y))
	return Ray{c.Origin.Add(offset), c.LowerLeft.Add(c.Horizontal.ScalarMulti(u)).Add(c.Vertical.ScalarMulti(v)).Subtract(c.Origin).Subtract(offset)}
}

// return random unit vector in the xy unit disk
// used for depth of field generation
func RandomInUnitDisk() Vector {
	var p Vector
	for {
		p = Vector{2.0*rand.Float64() - 1.0, 2.0*rand.Float64() - 1.0, 0.0}
		if p.Dot(p) < 1.0 {
			return p
		}
	}
}
