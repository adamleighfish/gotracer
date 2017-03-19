package pathtracer

import (
        "math"
)

type Texture interface {
        Value(u, v float64, p Vector) Vector
}

type ConstTexture struct {
        Color Vector
}

func (c *ConstTexture) Value(u, v float64, p Vector) Vector {
        return c.Color
}

type CheckeredTexture struct {
        Odd, Even Texture
}

func (c *CheckeredTexture) Value(u, v float64, p Vector) Vector {
        var sine float64 = math.Sin(10 * p.X) * math.Sin(10 * p.Y) * math.Sin(10 * p.Z)
        if sine < 0 {
                return c.Odd.Value(u, v, p)
        } else {
                return c.Even.Value(u, v, p)
        }
}
