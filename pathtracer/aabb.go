package pathtracer

import (
        "math"
)

type AABB struct {
        Min, Max Vector
}

func (box *AABB) Hit(r *Ray, tMin, tMax float64) bool {

        var invD, t0, t1 float64

        // test x direction
        invD = 1.0 / r.Direction().X
        t0 = (box.Min.X - r.Origin().X) * invD
        t1 = (box.Max.X - r.Origin().X) * invD
        if testBounds(&t0, &t1, &invD, &tMin, &tMax) {
                return false
        }

        // test y direction
        invD = 1.0 / r.Direction().Y
        t0 = (box.Min.Y - r.Origin().Y) * invD
        t1 = (box.Max.Y - r.Origin().Y) * invD
        if testBounds(&t0, &t1, &invD, &tMin, &tMax) {
                return false
        }

        // test z direction
        invD = 1.0 / r.Direction().Z
        t0 = (box.Min.Z - r.Origin().Z) * invD
        t1 = (box.Max.Z - r.Origin().Z) * invD
        if testBounds(&t0, &t1, &invD, &tMin, &tMax) {
                return false
        }
        return true
}

func testBounds(t0, t1, invD, tMin, tMax *float64) bool {
        if *invD < 0.0 {
                t0, t1 = t1, t0
        }
        if *t0 > *tMin {
                tMin = t0
        }
        if *t1 < *tMax {
                tMax = t1
        }
        if *tMax <= *tMin {
                return true
        }
        return false
}


func surroundingBox(box1, box2 *AABB) *AABB {
        a := Vector{X: math.Min(box1.Min.X, box2.Min.X),
                        Y: math.Min(box1.Min.Y, box2.Min.Y),
                        Z: math.Min(box1.Min.Z, box2.Min.Z)}
        b := Vector{X: math.Max(box1.Max.X, box2.Max.X),
                        Y: math.Max(box1.Max.Y, box2.Max.Y),
                        Z: math.Max(box1.Max.Z, box2.Max.Z)}
        surround := AABB{a, b}
        return &surround
}
