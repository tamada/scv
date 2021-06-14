package vector

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strings"
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

func (vp *VectorPair) Compare(algorithm Algorithm) (float64, error) {
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
	if s.Type() == "string" {
		return s.Value()
	}
	return fmt.Sprintf("%s (%v)", s.Type(), s.Value())
}

func newVector(source Source) *Vector {
	return &Vector{Source: source, values: map[string]int{}}
}

func NewVectorFromString(baseString string) *Vector {
	vector := newVector(NewSource("string", baseString))
	for _, c := range baseString {
		vector.Put(string(c))
	}
	return vector
}

func NewVectorFromJsonFile(baseString string) (*Vector, error) {
	file, err := os.Open(baseString)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	defer file.Close()
	raw, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("readall: %w", err)
	}
	return newVectorFromJsonData(raw, baseString)
}

func newVectorFromJsonData(raw []byte, baseString string) (*Vector, error) {
	readData := map[string]interface{}{}
	if err := json.Unmarshal(raw, &readData); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}
	result := map[string]int{}
	for key := range readData {
		result[key] = int(readData[key].(float64))
	}
	return &Vector{Source: NewSource("json", baseString), values: result}, nil
}

func NewTermVectorFromFile(baseString string) (*Vector, error) {
	file, err := os.Open(baseString)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return NewTermVectorFromReader(file, NewSource("term_file", baseString))
}

func NewByteVectorFromFile(baseString string) (*Vector, error) {
	reader, err := os.Open(baseString)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return NewByteVectorFromReader(reader, NewSource("byte_file", baseString))
}

func removeSpecialCharacters(line string) string {
	specialCharacters := []string{".", ",", ":", ";", "!", "[", "]", "(", ")", "<", ">", "@", "/", "{", "}", "?"}
	for _, sc := range specialCharacters {
		line = strings.ReplaceAll(line, sc, " ")
	}
	return line
}

func putTerms(line string, values map[string]int) {
	line = removeSpecialCharacters(line)
	terms := strings.Split(line, " ")
	for _, term := range terms {
		t := strings.ToLower(strings.TrimSpace(term))
		if t != "" {
			value := values[t]
			values[t] = value + 1
		}
	}
}

func NewTermVectorFromReader(reader io.Reader, source Source) (*Vector, error) {
	bufReader := bufio.NewReader(reader)
	values := map[string]int{}
	for {
		line, err := bufReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		putTerms(line, values)
	}
	return &Vector{Source: source, values: values}, nil
}

func putData(values map[string]int, data []byte, length int) {
	for i := 0; i < length; i++ {
		key := string(data[i])
		value := values[key]
		values[key] = (value + 1)
	}
}

func NewByteVectorFromReader(reader io.Reader, source Source) (*Vector, error) {
	values := map[string]int{}
	for {
		data := make([]byte, 1024)
		n, err := reader.Read(data)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		putData(values, data, n)
	}
	return &Vector{Source: source, values: values}, nil
}
