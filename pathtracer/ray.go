package pathtracer

type Ray struct {
	A Vector
	B Vector
}

func (r Ray) Origin() Vector {
	return r.A
}

func (r Ray) Direction() Vector {
	return r.B
}

func (r Ray) PointAtParameter(t float64) Vector {
	return r.A.Add(r.B.ScalarMulti(t))
}
