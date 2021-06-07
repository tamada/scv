package vector

type Vector struct {
	Source Source
	values map[string]int
}

type VectorPair struct {
	Vector1 *Vector
	Vector2 *Vector
}

func (vp *VectorPair) Compare(algorithm Algorithm) float32 {
	return algorithm.Compare(vp.Vector1, vp.Vector2)
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
