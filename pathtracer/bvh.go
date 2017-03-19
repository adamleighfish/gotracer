package pathtracer

import (
        "math/rand"
        "sort"
        "fmt"
)

type Node struct {
        Box             AABB
        Left, Right     Hitable
}

func BuildNode(l *HitableList, n int, t0, t1 float64) *Node {
        // choose axis to sort around
        axis := int(rand.Float64() * 3)
        if axis ==  0 {
                sort.Sort(ByX(l.Elements))
        } else if axis == 1 {
                sort.Sort(ByY(l.Elements))
        } else {
                sort.Sort(ByZ(l.Elements))
        }

        // separate the hitables into the two sides
        var thisNode Node
        if n == 1 {
                thisNode.Left = l.Elements[0]
                thisNode.Right = l.Elements[0]
        } else if n == 2 {
                thisNode.Left = l.Elements[0]
                thisNode.Right = l.Elements[1]
        } else {
                l1 := *l
                l2 := *l
                l1.Elements = l.Elements[:n/2]
                l2.Elements = l.Elements[n/2:]
                thisNode.Left = BuildNode(&l1, n/2, t0, t1)
                thisNode.Right = BuildNode(&l2, n - n/2, t0, t1)
        }

        leftBool, leftBox := thisNode.Left.BoundingBox(t0, t1)
        rightBool, rightBox := thisNode.Right.BoundingBox(t0, t1)

        if (!leftBool || !rightBool) {
                fmt.Println("Test")
        }

        thisNode.Box = *surroundingBox(&leftBox, &rightBox)

        return &thisNode
}

func (n *Node) BoundingBox(t0, t1 float64) (bool, AABB) {
        box := n.Box
        return true, box
}

func (n *Node) Hit(r *Ray, tMin, tMax float64) (bool, HitRecord) {
        success := n.Box.Hit(r, tMin, tMax)

        if success {
                leftBool, leftRec := n.Left.Hit(r, tMin, tMax)
                rightBool, rightRec := n.Right.Hit(r, tMin, tMax)

                if leftBool && rightBool {
                        if (leftRec.T < rightRec.T) {
                                return true, leftRec
                        }
                        return true, rightRec
                } else if leftBool {
                        return true, leftRec
                } else if rightBool {
                        return true, rightRec
                }
        }
        return false, HitRecord{}
}
