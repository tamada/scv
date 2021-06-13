package main

import (
	"fmt"
	"strings"

	"github.com/tamada/scv/vector"
)

func constructVectors(opts *options) ([]*vector.Vector, error) {
	inputTypes := buildInputTypes(opts)
	vectors := []*vector.Vector{}
	for i, arg := range opts.args {
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
	case "file":
		return vector.NewVectorFromFile(source)
	}
	return nil, fmt.Errorf("%s: unknown input type", inputType)
}

func buildInputTypes(opts *options) []string {
	types := strings.Split(opts.inputType, ",")
	if len(types) == 1 {
		results := []string{}
		for range opts.args {
			results = append(results, types[0])
		}
		types = results
	}
	return types
}
