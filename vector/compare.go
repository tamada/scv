package vector

import (
	"fmt"
	"strings"
)

type Algorithm interface {
	Compare(v1, v2 *Vector) float64
}

type simpsonComparator struct {
}

func (sc *simpsonComparator) Compare(v1, v2 *Vector) float64 {
	return 1.0
}

type diceComparator struct {
}

func (dc *diceComparator) Compare(v1, v2 *Vector) float64 {
	return 1.0
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
	return 1.0
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
