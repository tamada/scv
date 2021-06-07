package vector

import "testing"

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
