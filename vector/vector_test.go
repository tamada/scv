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

func TestJaccardIndex(t *testing.T) {
	testdata := []struct {
		giveString1    string
		giveString2    string
		wontSimilarity float64
	}{
		{"distance", "similarity", 0.3333333},
		{"android", "ipodtouch", 0.272727},
	}
	var threshold float64 = 1e-6
	for _, td := range testdata {
		vector1 := NewVectorFromString(td.giveString1)
		vector2 := NewVectorFromString(td.giveString2)
		algorithm, _ := NewAlgorithm("jaccard")
		result := algorithm.Compare(vector1, vector2)
		if math.Abs(result-td.wontSimilarity) > threshold {
			t.Errorf("jaccard(%s, %s) did not match, wont %f, got %f", td.giveString1, td.giveString2, td.wontSimilarity, result)
		}
	}
}

func TestSimpsonIndex(t *testing.T) {
	testdata := []struct {
		giveString1    string
		giveString2    string
		wontSimilarity float64
	}{
		{"distance", "similarity", 0.500000},
		{"android", "ipodtouch", 0.500000},
	}
	var threshold float64 = 1e-6
	for _, td := range testdata {
		vector1 := NewVectorFromString(td.giveString1)
		vector2 := NewVectorFromString(td.giveString2)
		algorithm, _ := NewAlgorithm("simpson")
		result := algorithm.Compare(vector1, vector2)
		if math.Abs(result-td.wontSimilarity) > threshold {
			t.Errorf("jaccard(%s, %s) did not match, wont %f, got %f", td.giveString1, td.giveString2, td.wontSimilarity, result)
		}
	}
}

func TestDiceIndex(t *testing.T) {
	testdata := []struct {
		giveString1    string
		giveString2    string
		wontSimilarity float64
	}{
		{"distance", "similarity", 0.500000},
		{"android", "ipodtouch", 0.428571},
	}
	var threshold float64 = 1e-6
	for _, td := range testdata {
		vector1 := NewVectorFromString(td.giveString1)
		vector2 := NewVectorFromString(td.giveString2)
		algorithm, _ := NewAlgorithm("dice")
		result := algorithm.Compare(vector1, vector2)
		if math.Abs(result-td.wontSimilarity) > threshold {
			t.Errorf("jaccard(%s, %s) did not match, wont %f, got %f", td.giveString1, td.giveString2, td.wontSimilarity, result)
		}
	}
}

func TestCosineSimilarity(t *testing.T) {
	testdata := []struct {
		giveString1    string
		giveString2    string
		wontSimilarity float64
	}{
		{"distance", "similarity", 0.530330},
		{"android", "ipodtouch", 0.502519},
	}
	var threshold float64 = 1e-6
	for _, td := range testdata {
		vector1 := NewVectorFromString(td.giveString1)
		vector2 := NewVectorFromString(td.giveString2)
		algorithm, _ := NewAlgorithm("cosine")
		result := algorithm.Compare(vector1, vector2)
		if math.Abs(result-td.wontSimilarity) > threshold {
			t.Errorf("jaccard(%s, %s) did not match, wont %f, got %f", td.giveString1, td.giveString2, td.wontSimilarity, result)
		}
	}
}
