package vector

import "fmt"

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

func (vp *VectorPair) Intersect() *Vector {
	return &Vector{Source: vp, values: vp.values}
}

func (vp *VectorPair) Union() *Vector {
	result := map[string]int{}
	for key, value := range vp.Vector1.values {
		v, ok := vp.Vector2.values[key]
		if ok {
			result[key] = v + value
		}
	}
	return &Vector{Source: vp, values: result}
}

type Vector struct {
	Source Source
	values map[string]int
}

func (vector *Vector) Length() float64 {
	return float64(len(vector.values))
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
