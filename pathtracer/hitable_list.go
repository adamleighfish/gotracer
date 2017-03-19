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

func (l *HitableList) BoundingBox(t0, t1 float64) (bool, AABB) {
        var tempBox AABB

        if len(l.Elements) < 1 {
                return false, tempBox
        }
        first, tempBox := l.Elements[0].BoundingBox(t0, t1)

        if !first {
                return false, tempBox
        }
        box := tempBox

        for _, element := range l.Elements {
                success, tempBox := element.BoundingBox(t0, t1)
                if success {
                        box = *surroundingBox(&box, &tempBox)
                } else {
                        return false, tempBox
                }
        }

        return true, box
}

func (l *HitableList) Add(h Hitable) {
	l.Elements = append(l.Elements, h)
}

type ByX []Hitable
type ByY []Hitable
type ByZ []Hitable

func (a ByX) Len() int { return len(a) }
func (a ByY) Len() int { return len(a) }
func (a ByZ) Len() int { return len(a) }

func (a ByX) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByY) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByZ) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a ByX) Less(i, j int) bool {
        _, tempI := a[i].BoundingBox(0.0, 1.0)
        _, tempJ := a[j].BoundingBox(0.0, 1.0)
        return tempI.Min.X < tempJ.Min.X
}

func (a ByY) Less(i, j int) bool {
        _, tempI := a[i].BoundingBox(0.0, 1.0)
        _, tempJ := a[j].BoundingBox(0.0, 1.0)
        return tempI.Min.Y < tempJ.Min.Y
}
func (a ByZ) Less(i, j int) bool {
        _, tempI := a[i].BoundingBox(0.0, 1.0)
        _, tempJ := a[j].BoundingBox(0.0, 1.0)
        return tempI.Min.Z < tempJ.Min.Z
}
