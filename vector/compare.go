package vector

import (
	"fmt"
	"math"
	"strings"
)

type Algorithm interface {
	Compare(v1, v2 *Vector) float64
}

type simpsonComparator struct {
}

func (sc *simpsonComparator) Compare(v1, v2 *Vector) float64 {
	vp := NewVectorPair(v1, v2)
	intersect := vp.Intersect()
	return intersect.Length() / math.Min(v1.Length(), v2.Length())
}

type diceComparator struct {
}

func (dc *diceComparator) Compare(v1, v2 *Vector) float64 {
	vp := NewVectorPair(v1, v2)
	intersect := vp.Intersect()
	return 2.0 * intersect.Length() / (v1.Length() + v2.Length())
}

type jaccardComparator struct {
}

func (jc *jaccardComparator) Compare(v1, v2 *Vector) float64 {
	vp := NewVectorPair(v1, v2)
	intersect := vp.Intersect()
	union := vp.Union()
	return intersect.Length() / union.Length()
}

type cosineComparator struct {
}

func (sc *cosineComparator) Compare(v1, v2 *Vector) float64 {
	vp := NewVectorPair(v1, v2)
	return vp.InnerProduct()
}

func NewAlgorithm(comparatorType string) (Algorithm, error) {
	switch strings.ToLower(comparatorType) {
	case "simpson":
		return &simpsonComparator{}, nil
	case "dice":
		return &diceComparator{}, nil
	case "jaccard":
		return &jaccardComparator{}, nil
	case "cosine":
		return &cosineComparator{}, nil
	}
	return nil, fmt.Errorf("%s: unknown algorithm", comparatorType)
}
