package vector

import (
	"fmt"
	"math"
)

type VectorPair struct {
	Vector1 *Vector
	Vector2 *Vector
	values  map[string]int
}

func NewVectorPair(v1, v2 *Vector) *VectorPair {
	vp := &VectorPair{Vector1: v1, Vector2: v2, values: map[string]int{}}
	putAll(vp.values, v1.values)
	putAll(vp.values, v2.values)
	return vp
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

func (vp *VectorPair) Type() string {
	return "vector_pair"
}

func (vp *VectorPair) Value() string {
	return fmt.Sprintf("%s (%v, %v)", vp.Type(), vp.Vector1.Value(), vp.Vector2.Value())
}

func (vp *VectorPair) Compare(algorithm Algorithm) float64 {
	return algorithm.Compare(vp.Vector1, vp.Vector2)
}

func (vp *VectorPair) Length() float64 {
	return float64(len(vp.values))
}

func (vp *VectorPair) Union() *Vector {
	return &Vector{Source: vp, values: vp.values}
}

func (vp *VectorPair) Intersect() *Vector {
	result := map[string]int{}
	for key, value := range vp.Vector1.values {
		v, ok := vp.Vector2.values[key]
		if ok {
			result[key] = v + value
		}
	}
	return &Vector{Source: vp, values: result}
}

func (vp *VectorPair) InnerProduct() float64 {
	sum := 0
	for key, _ := range vp.values {
		v1 := vp.Vector1.values[key]
		v2 := vp.Vector2.values[key]
		sum = sum + (v1 * v2)
	}
	return float64(sum) / (vp.Vector1.Norm() * vp.Vector2.Norm())
}

type Vector struct {
	Source       Source
	values       map[string]int
	averageValue float64
}

func (vector *Vector) Length() float64 {
	return float64(len(vector.values))
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
