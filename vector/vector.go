package vector

import (
	"fmt"
	"math"
)

type VectorPair struct {
	Vector1 *Vector
	Vector2 *Vector
}

func NewVectorPair(v1, v2 *Vector) *VectorPair {
	return &VectorPair{Vector1: v1, Vector2: v2}
}

func putAll(to, from map[string]int) {
	for key, value := range from {
		valueOfPair, ok := to[key]
		if ok {
			valueOfPair = valueOfPair + value
		}
		to[key] = valueOfPair
	}
}

func (vp *VectorPair) Compare(algorithm Algorithm) float64 {
	return algorithm.Compare(vp.Vector1, vp.Vector2)
}

type Vector struct {
	Source       Source
	values       map[string]int
	averageValue float64
}

func (vector *Vector) Length() float64 {
	return float64(len(vector.values))
}

func (vector *Vector) Union(other *Vector) *Vector {
	result := map[string]int{}
	putAll(result, vector.values)
	putAll(result, other.values)
	return &Vector{Source: NewSource("union", kind("union", vector, other)), values: result}
}

func kind(prefix string, v1, v2 *Vector) string {
	return fmt.Sprintf("%s(%s, %s)", prefix, v1.Source.Value(), v2.Source.Value())
}

func (vector *Vector) Intersect(other *Vector) *Vector {
	result := map[string]int{}
	for key, value := range vector.values {
		v, ok := other.values[key]
		if ok {
			result[key] = v + value
		}
	}
	return &Vector{Source: NewSource("intersect", kind("intersect", vector, other)), values: result}
}

func (vector *Vector) InnerProduct(other *Vector) float64 {
	sum := 0
	union := vector.Union(other)
	for key := range union.values {
		v1 := vector.values[key]
		v2 := other.values[key]
		sum = sum + (v1 * v2)
	}
	return float64(sum) / (vector.Norm() * other.Norm())
}

func (vector *Vector) Norm() float64 {
	sum := 0
	for _, value := range vector.values {
		sum = sum + value*value
	}
	return math.Sqrt(float64(sum))
}

func (vector *Vector) Put(key string) (postFrequency int) {
	value, ok := vector.values[key]
	if !ok {
		value = 0
	}
	postFrequency = value + 1
	vector.values[key] = postFrequency
	return postFrequency
}

func (vector *Vector) Get(key string) int {
	value, ok := vector.values[key]
	if !ok {
		return 0
	}
	return value
}

func (vector *Vector) average(n float64) float64 {
	if vector.averageValue == 0.0 {
		sum := 0
		for _, value := range vector.values {
			sum = sum + value
		}
		vector.averageValue = float64(sum) / n
	}
	return vector.averageValue
}

func (vector *Vector) Value() string {
	s := vector.Source
	return fmt.Sprintf("%s (%v)", s.Type(), s.Value())
}

func newVector(source Source) *Vector {
	return &Vector{Source: source, values: map[string]int{}}
}

func NewVectorFromString(baseString string) *Vector {
	vector := newVector(newStringSource(baseString))
	for _, c := range baseString {
		vector.Put(string(c))
	}
	return vector
}
