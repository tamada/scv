package main

import (
	"fmt"
	"strings"

	"github.com/tamada/scv/vector"
)

func constructVectors(opts *options) ([]*vector.Vector, error) {
	inputTypes := buildInputTypes(opts)
	vectors := []*vector.Vector{}
	for i, arg := range opts.input.args {
		vector, err := constructVector(arg, inputTypes[i])
		if err != nil {
			return nil, err
		}
		vectors = append(vectors, vector)
	}
	return vectors, nil
}

func constructVector(source, inputType string) (*vector.Vector, error) {
	switch inputType {
	case "string":
		return vector.NewVectorFromString(source), nil
	case "json":
		return vector.NewVectorFromJsonFile(source)
	case "term_file":
		return vector.NewTermVectorFromFile(source)
	case "byte_file":
		return vector.NewByteVectorFromFile(source)
	}
	return nil, fmt.Errorf("%s: unknown input type", inputType)
}

func buildInputTypes(opts *options) []string {
	types := strings.Split(opts.input.inputType, ",")
	if len(types) == 1 {
		results := []string{}
		for range opts.input.args {
			results = append(results, types[0])
		}
		types = results
	}
	return types
}
