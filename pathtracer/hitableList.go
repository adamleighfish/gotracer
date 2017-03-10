package pathtracer

type HitableList struct {
	Elements []Hitable
}

func (l *HitableList) Hit(r *Ray, tMin float64, tMax float64) (bool, HitRecord) {
	hitAnything := false
	closest := tMax
	record := HitRecord{}

	// test if any of the elements in the list are hit by r
	for _, element := range l.Elements {

		hit, tempRecord := element.Hit(r, tMin, closest)

		// if one is hit set it as the closest object and continue
		if hit {
			hitAnything = true
			closest = tempRecord.T
			record = tempRecord
		}
	}
	return hitAnything, record
}

func (l *HitableList) Add(h Hitable) {
	l.Elements = append(l.Elements, h)
}
