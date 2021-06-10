package vector

import (
	"fmt"
	"math"
	"strings"
)

type Algorithm interface {
	Compare(v1, v2 *Vector) (float64, error)
}

type simpsonComparator struct {
}

func (sc *simpsonComparator) Compare(v1, v2 *Vector) (float64, error) {
	intersect := v1.Intersect(v2)
	return intersect.Length() / math.Min(v1.Length(), v2.Length()), nil
}

type diceComparator struct {
}

func (dc *diceComparator) Compare(v1, v2 *Vector) (float64, error) {
	intersect := v1.Intersect(v2)
	return 2.0 * intersect.Length() / (v1.Length() + v2.Length()), nil
}

type jaccardComparator struct {
}

func (jc *jaccardComparator) Compare(v1, v2 *Vector) (float64, error) {
	intersect := v1.Intersect(v2)
	union := v1.Union(v2)
	return intersect.Length() / union.Length(), nil
}

type cosineComparator struct {
}

func (sc *cosineComparator) Compare(v1, v2 *Vector) (float64, error) {
	return v1.InnerProduct(v2), nil
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

func (pc *pearsonCorrelation) Compare(v1, v2 *Vector) (float64, error) {
	union := v1.Union(v2)
	covariance := calcCovariance(v1, v2, union)
	deviation1 := calcDeviation(v1, union)
	deviation2 := calcDeviation(v2, union)
	return covariance / (deviation1 * deviation2), nil
}

type euclideanDistance struct {
}

func (ed *euclideanDistance) Compare(v1, v2 *Vector) (float64, error) {
	union := v1.Union(v2)
	sum := 0
	for key := range union.values {
		value1 := v1.values[key]
		value2 := v2.values[key]
		sum = sum + ((value1 - value2) * (value1 - value2))
	}
	return math.Sqrt(float64(sum)), nil
}

type levenshteinDistance struct {
}

func (ld *levenshteinDistance) Compare(v1, v2 *Vector) (float64, error) {
	if v1.Source.Type() != "string" || v2.Source.Type() != "string" {
		return 0, fmt.Errorf("levenshtein distance: type of two vector must be string")
	}
	s1, s2 := v1.Source.Value(), v2.Source.Value()
	table := constructTable(s1, s2)
	calcLevenshtein(table, s1, s2)
	return float64(table[len(s1)][len(s2)]), nil
}

func calcLevenshtein(table [][]int, s1, s2 string) {
	for i := 1; i < len(table); i++ {
		for j := 1; j < len(table[i]); j++ {
			cost := 1
			if s1[i-1] == s2[j-1] {
				cost = 0
			}
			d1 := table[i-1][j] + 1
			d2 := table[i][j-1] + 1
			d3 := table[i-1][j-1] + cost
			table[i][j] = min(d1, d2, d3)
		}
	}
}

func constructTable(s1, s2 string) [][]int {
	table := [][]int{}
	for j := 0; j <= len(s1); j++ {
		values := []int{}
		for i := 0; i <= len(s2); i++ {
			values = append(values, 0)
		}
		table = append(table, values)
	}
	return initTable(table)
}

func initTable(table [][]int) [][]int {
	for i := 0; i < len(table); i++ {
		table[i][0] = i
	}
	for j := 0; j < len(table[0]); j++ {
		table[0][j] = j
	}
	return table
}

func min(values ...int) int {
	min := values[0]
	for _, value := range values {
		if min > value {
			min = value
		}
	}
	return min
}

/*
func printTable(table [][]int) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			fmt.Printf("%2d ", table[i][j])
		}
		fmt.Println()
	}
}
*/

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
	case "levenshtein":
		return &levenshteinDistance{}, nil
	}
	return nil, fmt.Errorf("%s: unknown algorithm", comparatorType)
}
