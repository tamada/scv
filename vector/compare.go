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
	intersect := v1.Intersect(v2)
	return intersect.Length() / math.Min(v1.Length(), v2.Length())
}

type diceComparator struct {
}

func (dc *diceComparator) Compare(v1, v2 *Vector) float64 {
	intersect := v1.Intersect(v2)
	return 2.0 * intersect.Length() / (v1.Length() + v2.Length())
}

type jaccardComparator struct {
}

func (jc *jaccardComparator) Compare(v1, v2 *Vector) float64 {
	intersect := v1.Intersect(v2)
	union := v1.Union(v2)
	return intersect.Length() / union.Length()
}

type cosineComparator struct {
}

func (sc *cosineComparator) Compare(v1, v2 *Vector) float64 {
	return v1.InnerProduct(v2)
}

type pearsonCorrelation struct {
}

func calcDeviation(v, union *Vector) float64 {
	sum := float64(0)
	average := v.average(union.Length())
	for key, _ := range union.values {
		value := float64(v.values[key]) - average
		sum = sum + (value * value)
	}
	return math.Sqrt(sum)
}

func calcCovariance(v1, v2, union *Vector) float64 {
	covariance := float64(0)
	xAverage := v1.average(union.Length())
	yAverage := v2.average(union.Length())
	for key, _ := range union.values {
		x := float64(v1.values[key])
		y := float64(v2.values[key])
		covariance = covariance + ((x - xAverage) * (y - yAverage))
	}
	return covariance
}

func (pc *pearsonCorrelation) Compare(v1, v2 *Vector) float64 {
	union := v1.Union(v2)
	covariance := calcCovariance(v1, v2, union)
	deviation1 := calcDeviation(v1, union)
	deviation2 := calcDeviation(v2, union)
	return covariance / (deviation1 * deviation2)
}

type euclideanDistance struct {
}

func (ed *euclideanDistance) Compare(v1, v2 *Vector) float64 {
	union := v1.Union(v2)
	sum := 0
	for key := range union.values {
		value1 := v1.values[key]
		value2 := v2.values[key]
		sum = sum + ((value1 - value2) * (value1 - value2))
	}
	return math.Sqrt(float64(sum))
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
	case "pearson":
		return &pearsonCorrelation{}, nil
	case "euclidean":
		return &euclideanDistance{}, nil
	}
	return nil, fmt.Errorf("%s: unknown algorithm", comparatorType)
}
