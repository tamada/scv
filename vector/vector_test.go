package vector

import (
	"math"
	"testing"
)

func TestConstructVector(t *testing.T) {
	testdata := []struct {
		giveString string
		wontVector map[string]int
	}{
		{"test", map[string]int{"t": 2, "e": 1, "s": 1}},
		{"abracadabra", map[string]int{"a": 5, "b": 2, "r": 2, "c": 1, "d": 1}},
	}
	for _, td := range testdata {
		vector := NewVectorFromString(td.giveString)
		for key, value := range td.wontVector {
			gotValue := vector.Get(key)
			if gotValue != value {
				t.Errorf("%s: vector did not match, wont %v, got %v", td.giveString, td.wontVector, vector.values)
				continue
			}
		}
	}
}

type TestData struct {
	giveString1    string
	giveString2    string
	wontSimilarity float64
}

func execTest(t *testing.T, data []TestData, algorithmName string) {
	var threshold float64 = 1e-6
	for _, datum := range data {
		v1 := NewVectorFromString(datum.giveString1)
		v2 := NewVectorFromString(datum.giveString2)
		algorithm, _ := NewAlgorithm(algorithmName)
		gotSimilarity, _ := algorithm.Compare(v1, v2)
		if math.Abs(gotSimilarity-datum.wontSimilarity) > threshold {
			t.Errorf("%s(%s, %s) did not match, wont %f, got %f", algorithmName, datum.giveString1, datum.giveString2, datum.wontSimilarity, gotSimilarity)
		}
	}
}

func TestJaccardIndex(t *testing.T) {
	execTest(t, []TestData{
		{"distance", "similarity", 0.3333333},
		{"android", "ipodtouch", 0.272727},
	}, "jaccard")
}

func TestSimpsonIndex(t *testing.T) {
	execTest(t, []TestData{
		{"distance", "similarity", 0.500000},
		{"android", "ipodtouch", 0.500000},
	}, "simpson")
}

func TestDiceIndex(t *testing.T) {
	execTest(t, []TestData{
		{"distance", "similarity", 0.500000},
		{"android", "ipodtouch", 0.428571},
	}, "dice")
}

func TestCosineSimilarity(t *testing.T) {
	execTest(t, []TestData{
		{"distance", "similarity", 0.530330},
		{"android", "ipodtouch", 0.502519},
	}, "cosine")
}

func TestPearsonCorrelation(t *testing.T) {
	execTest(t, []TestData{
		{"distance", "similarity", -0.147441956},
		{"android", "ipodtouch", -0.178885438},
	}, "pearson")
}

func TestEuclideanDistance(t *testing.T) {
	execTest(t, []TestData{
		{"distance", "similarity", 3.464101615},
		{"android", "ipodtouch", 3.16227766},
	}, "euclidean")
}

func TestLevenshteinDistance(t *testing.T) {
	execTest(t, []TestData{
		{"distance", "similarity", 8.0},
		{"android", "ipodtouch", 7.0},
	}, "levenshtein")
}
